package storage

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"

	"github.com/T-V-N/gourlshortener/internal/config"
)

type URL struct {
	Hash string `json:"hash"`
	URL  string `json:"url"`
}

type Storage struct {
	db              map[string]string
	FileStoragePath string
}

func InitStorage(data map[string]string, cfg *config.Config) *Storage {
	if cfg.FileStoragePath == "" {
		if data == nil {
			return &Storage{make(map[string]string), cfg.FileStoragePath}
		}

		return &Storage{data, cfg.FileStoragePath}
	}

	file, err := os.OpenFile(cfg.FileStoragePath, os.O_RDONLY, 0o777)
	if err != nil {
		return &Storage{data, cfg.FileStoragePath}
	}

	scanner := bufio.NewScanner(file)

	for {
		if !scanner.Scan() {
			break
		}

		lineBytes := scanner.Bytes()
		url := URL{}
		err = json.Unmarshal(lineBytes, &url)

		if err != nil {
			break
		}

		data[url.Hash] = url.URL
	}

	defer file.Close()

	return &Storage{data, cfg.FileStoragePath}
}

func (st *Storage) SaveURL(url string) (string, error) {
	hash := md5.Sum([]byte(url))
	shortHash := hex.EncodeToString(hash[:4])
	st.db[shortHash] = url

	if st.FileStoragePath != "" {
		file, err := os.OpenFile(st.FileStoragePath+"/db", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o777)
		if err != nil {
			return "", err
		}

		data, err := json.Marshal(&URL{shortHash, url})
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

	return url, nil
}
