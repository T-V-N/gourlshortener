package storage

import (
	"bufio"
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
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

func (st *FileStorage) SaveURL(ctx context.Context, url, UID string) (string, error) {
	hash := md5.Sum([]byte(url))
	ShortURL := st.cfg.BaseURL + "/" + hex.EncodeToString(hash[:4])

	st.db[hex.EncodeToString(hash[:4])] = URL{UID, ShortURL, url}

	if st.cfg.FileStoragePath != "" {
		file, err := os.OpenFile(st.cfg.FileStoragePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o777)
		if err != nil {
			return "", err
		}

		data, err := json.Marshal(&URL{UID, ShortURL, url})
		if err != nil {
			return "", err
		}

		data = append(data, '\n')

		_, err = file.Write(data)

		defer file.Close()

		return ShortURL, err
	}

	return ShortURL, nil
}

func (st *FileStorage) GetURL(ctx context.Context, hash string) (string, error) {
	url, exists := st.db[hash]
	if !exists {
		return hash, errors.New("an URL with this hash doesn't exist")
	}

	return url.URL, nil
}

func (st *FileStorage) GetUrlsByUID(ctx context.Context, uid string) ([]URL, error) {
	result := []URL{}
	for _, url := range st.db {
		if url.UID == uid {
			result = append(result, url)
		}
	}

	return result, nil
}

func (st *FileStorage) IsAlive(ctx context.Context) (bool, error) {
	//file storage always alive
	return true, nil
}
