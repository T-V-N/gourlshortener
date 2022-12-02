package storage

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"os"

	"github.com/T-V-N/gourlshortener/internal/config"
)

type FileStorage struct {
	db  map[string]URL
	cfg config.Config
}

func InitFileStorage(data map[string]URL, cfg *config.Config) *FileStorage {
	if cfg.FileStoragePath == "" {
		if data == nil {
			return &FileStorage{make(map[string]URL), *cfg}
		}

		return &FileStorage{data, *cfg}
	}

	file, err := os.OpenFile(cfg.FileStoragePath, os.O_RDONLY, 0o777)
	if err != nil {
		return &FileStorage{data, *cfg}
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		url := URL{}
		if err := json.NewDecoder(bytes.NewBuffer(scanner.Bytes())).Decode(&url); err != nil {
			break
		}

		data[url.ShortURL] = url
	}

	defer file.Close()

	return &FileStorage{data, *cfg}
}

func (st *FileStorage) SaveURL(ctx context.Context, url, uid, hash string) error {
	st.db[hash] = URL{uid, hash, url, false}

	if st.cfg.FileStoragePath != "" {
		file, err := os.OpenFile(st.cfg.FileStoragePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o777)
		if err != nil {
			return err
		}

		data, err := json.Marshal(&URL{uid, hash, url, false})
		if err != nil {
			return err
		}

		data = append(data, '\n')

		_, err = file.Write(data)

		defer file.Close()

		return err
	}

	return nil
}

func (st *FileStorage) GetURL(ctx context.Context, hash string) (URL, error) {
	url, exists := st.db[hash]
	if !exists {
		return url, errors.New("an URL with this hash doesn't exist")
	}

	return url, nil
}

func (st *FileStorage) GetUrlsByUID(ctx context.Context, uid string) ([]URL, error) {
	result := []URL{}

	for _, url := range st.db {
		if url.UID == uid && !url.IsDeleted {
			result = append(result, url)
		}
	}

	return result, nil
}

func (st *FileStorage) IsAlive(context.Context) (bool, error) {
	// file storage always alive
	return true, nil
}

func (st *FileStorage) BatchSaveURL(ctx context.Context, urls []URL) error {
	for _, url := range urls {
		st.db[url.ShortURL] = URL{url.UID, url.ShortURL, url.URL, false}
	}

	if st.cfg.FileStoragePath != "" {
		file, err := os.OpenFile(st.cfg.FileStoragePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o777)
		if err != nil {
			return err
		}

		err = json.NewEncoder(file).Encode(urls)
		if err != nil {
			return err
		}

		_, err = file.Write([]byte{'\n'})

		defer file.Close()

		return err
	}

	return nil
}

func (st *FileStorage) KillConn() error {
	return nil
}

func (st *FileStorage) DeleteURLs(ctx context.Context, entries []DeletionEntry) error {
	for _, entry := range entries {
		url, exists := st.db[entry.Hash]
		if exists && st.db[entry.Hash].UID == entry.UID {
			url.IsDeleted = true
			st.db[entry.Hash] = url
		}
	}

	return nil
}
