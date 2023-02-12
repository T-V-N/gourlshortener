package storage_test

import (
	"bufio"
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/T-V-N/gourlshortener/internal/config"
	"github.com/T-V-N/gourlshortener/internal/middleware/auth"
	"github.com/T-V-N/gourlshortener/internal/storage"
	"github.com/caarlos0/env/v6"

	"github.com/stretchr/testify/assert"
)

func InitFileTestConfig(withFile bool) (*config.Config, error) {
	var cfg *config.Config
	if withFile {
		cfg = &config.Config{FileStoragePath: "./file"}
	} else {
		cfg = &config.Config{}
	}

	err := env.Parse(cfg)
	cfg.DatabaseDSN = ""

	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	return cfg, nil
}

func Test_FileStorage(t *testing.T) {
	t.Run("check in memory save and get", func(t *testing.T) {
		cfg, _ := InitFileTestConfig(false)
		st := storage.InitStorage(map[string]storage.URL{}, cfg)

		url := "https://yandexx.ru"
		uid := "1"
		hash := "xxxxxx"
		brokenHash := "xyzyxyzy"

		st.SaveURL(context.Background(), url, uid, hash)

		goodUrl, err := st.GetURL(context.Background(), hash)

		if err != nil {
			panic("Error getting info from memory")
		}

		_, err = st.GetURL(context.Background(), brokenHash)

		assert.Equal(t, goodUrl.URL, url)
		assert.NotNil(t, err)
	})

	t.Run("check file db save and get", func(t *testing.T) {
		cfg, _ := InitFileTestConfig(true)
		st := storage.InitStorage(map[string]storage.URL{}, cfg)

		url := "https://yandexx.ru"
		uid := "1"
		hash1 := "xxxxxx"
		hash2 := "xxxxxxz"
		hash3 := "xxxxxxc"

		st.SaveURL(context.Background(), url, uid, hash1)
		st.SaveURL(context.Background(), url, uid, hash2)
		st.SaveURL(context.Background(), url, uid, hash3)

		file, err := os.OpenFile(cfg.FileStoragePath, os.O_RDONLY, 0o777)
		if err != nil {
			panic("no file available")
		}

		scanner := bufio.NewScanner(file)

		urls := []storage.URL{}

		for scanner.Scan() {
			url := storage.URL{}
			if err := json.NewDecoder(bytes.NewBuffer(scanner.Bytes())).Decode(&url); err != nil {
				break
			}
			urls = append(urls, url)
		}

		defer file.Close()
		defer os.Remove(cfg.FileStoragePath)

		assert.Equal(t, 3, len(urls))
		assert.Equal(t, urls[0].ShortURL, hash1)
		assert.Equal(t, urls[1].ShortURL, hash2)
		assert.Equal(t, urls[2].ShortURL, hash3)

		err = st.KillConn()
		assert.Nil(t, err, "file storage must always be killable")
		ok, err := st.IsAlive(context.Background())
		assert.True(t, ok, "file storage must always be alive")
	})

	t.Run("deletion test", func(t *testing.T) {
		cfg, _ := InitFileTestConfig(false)
		st := storage.InitStorage(map[string]storage.URL{
			"xxx": {UID: "1", ShortURL: "xxx", URL: "https://ya.ru", IsDeleted: false},
			"yyy": {UID: "1", ShortURL: "yyy", URL: "https://ya2.ru", IsDeleted: false},
			"zzz": {UID: "1", ShortURL: "zzz", URL: "https://ya3.ru", IsDeleted: false},
		}, cfg)

		urls, _ := st.GetUrlsByUID(context.Background(), "1")
		assert.Equal(t, 3, len(urls))

		st.DeleteURLs(context.Background(), []storage.DeletionEntry{
			{"1", "xxx"},
			{"1", "yyy"},
		})

		urls, _ = st.GetUrlsByUID(context.Background(), "1")

		assert.Equal(t, 1, len(urls))
	})

	t.Run("batch save to file test", func(t *testing.T) {
		cfg, _ := InitFileTestConfig(true)
		defer os.Remove(cfg.FileStoragePath)

		st := storage.InitStorage(map[string]storage.URL{}, cfg)

		urls := []storage.URL{
			{UID: "1", ShortURL: "xxx", URL: "https://ya.ru", IsDeleted: false},
			{UID: "1", ShortURL: "yyy", URL: "https://ya2.ru", IsDeleted: false},
			{UID: "1", ShortURL: "zzz", URL: "https://ya3.ru", IsDeleted: false}}

		st.BatchSaveURL(context.Background(), urls)

		urlsInStorage, _ := st.GetUrlsByUID(context.Background(), "1")

		assert.Equal(t, len(urls), len(urlsInStorage))
	})
}

func Test_DB(t *testing.T) {
	cfg, _ := config.Init()

	t.Run("connect db and ensure its ready and gracefully stopable", func(t *testing.T) {
		st, err := storage.InitDBStorage(cfg)

		assert.Nil(t, err, "connection must be made")

		alive, err := st.IsAlive(context.Background())

		assert.Equal(t, true, alive)
		assert.Nil(t, err, "connection must be alive")

		err = st.KillConn()
		assert.Nil(t, err, "connection must be killed")
	})

	t.Run("Create, get, delete batch test", func(t *testing.T) {
		st, err := storage.InitDBStorage(cfg)

		randBytes, _ := auth.GenerateRandom(4)
		randInd := hex.EncodeToString(randBytes)
		defer st.KillConn()

		assert.Nil(t, err, "connection must be made")

		inUrls := []storage.URL{
			{UID: randInd, ShortURL: randInd + "1", URL: "https://ya.ru"},
			{UID: randInd, ShortURL: randInd + "2", URL: "https://ya2.ru"},
			{UID: randInd, ShortURL: randInd + "3", URL: "https://ya3.ru"},
		}

		expectedUrls := []storage.URL{
			{ShortURL: randInd + "1", URL: "https://ya.ru"},
			{ShortURL: randInd + "2", URL: "https://ya2.ru"},
			{ShortURL: randInd + "3", URL: "https://ya3.ru"},
		}

		deletedUrls := []storage.DeletionEntry{
			{UID: randInd, Hash: randInd + "1"},
			{UID: randInd, Hash: randInd + "2"},
		}

		err = st.BatchSaveURL(context.Background(), inUrls)
		assert.Nil(t, err, "links must be saved properly")

		outUrls, err := st.GetUrlsByUID(context.Background(), randInd)
		assert.Nil(t, err, "links must be returned properly")
		assert.EqualValues(t, expectedUrls, outUrls)

		err = st.DeleteURLs(context.Background(), deletedUrls)
		assert.Nil(t, err, "links must be deleted properly")

		outUrlsWithoutDeleted, err := st.GetUrlsByUID(context.Background(), randInd)
		assert.Nil(t, err, "links must be returned properly")
		assert.Equal(t, []storage.URL{expectedUrls[2]}, outUrlsWithoutDeleted)
	})

	t.Run("Create, get test", func(t *testing.T) {
		st, err := storage.InitDBStorage(cfg)

		randBytes, _ := auth.GenerateRandom(4)
		randInd := hex.EncodeToString(randBytes)
		defer st.KillConn()

		assert.Nil(t, err, "connection must be made")

		inUrl := storage.URL{UID: randInd, ShortURL: randInd, URL: "https://ya.ru"}

		err = st.SaveURL(context.Background(), inUrl.URL, inUrl.UID, inUrl.ShortURL)
		assert.Nil(t, err, "link must be saved properly")

		outUrls, err := st.GetURL(context.Background(), randInd)
		assert.Nil(t, err, "link must be returned properly")
		assert.EqualValues(t, inUrl, outUrls)
	})

}
