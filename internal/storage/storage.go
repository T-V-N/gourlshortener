package storage

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
)

type Storage struct {
	db map[string]string
}

func NewStorage(data map[string]string) *Storage {
	if data == nil {
		return &Storage{make(map[string]string)}
	}

	return &Storage{data}
}

func (st *Storage) SaveURL(url string) (string, error) {
	hash := md5.Sum([]byte(url))
	shortHash := hex.EncodeToString(hash[:4])
	st.db[shortHash] = url

	return shortHash, nil
}

func (st *Storage) GetURL(hash string) (string, error) {
	url, exists := st.db[hash]
	if !exists {
		return hash, errors.New("an URL with this hash doesn't exist")
	}

	return url, nil
}
