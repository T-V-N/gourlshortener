package storage

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"

	"github.com/T-V-N/gourlshortener/internal/config"
)

type URL struct {
	UID  string `json:"-"`
	Hash string `json:"short_url"`
	URL  string `json:"original_url"`
}

type Storage struct {
	db  map[string]URL
	cfg config.Config
}

func InitStorage(data map[string]URL, cfg *config.Config) *Storage {
	if cfg.FileStoragePath == "" {
		if data == nil {
			return &Storage{make(map[string]URL), *cfg}
		}

		return &Storage{data, *cfg}
	}

	file, err := os.OpenFile(cfg.FileStoragePath, os.O_RDONLY, 0o777)
	if err != nil {
		return &Storage{data, *cfg}
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		url := URL{}
		if err := json.NewDecoder(bytes.NewBuffer(scanner.Bytes())).Decode(&url); err != nil {
			break
		}

		data[url.Hash] = url

	}

	defer file.Close()

	return &Storage{data, *cfg}
}

func (st *Storage) SaveURL(url, UID string) (string, error) {
	hash := md5.Sum([]byte(url))
	shortHash := st.cfg.BaseURL + "/" + hex.EncodeToString(hash[:4])

	st.db[shortHash] = URL{UID, shortHash, url}

	if st.cfg.FileStoragePath != "" {
		file, err := os.OpenFile(st.cfg.FileStoragePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o777)
		if err != nil {
			return "", err
		}

		data, err := json.Marshal(&URL{UID, shortHash, url})
		if err != nil {
			return "", err
		}

		data = append(data, '\n')

		_, err = file.Write(data)

		defer file.Close()

		return shortHash, err
	}

	return shortHash, nil
}

func (st *Storage) GetURL(hash string) (string, error) {
	url, exists := st.db[hash]
	if !exists {
		return hash, errors.New("an URL with this hash doesn't exist")
	}

	return url.URL, nil
}

func (st *Storage) GetUrlsByUID(uid string) ([]URL, error) {
	result := []URL{}
	for _, url := range st.db {
		if url.UID == uid {
			result = append(result, url)
		}
	}

	return result, nil
}
