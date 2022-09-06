package storage

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
)

type Storage struct {
	db map[string]string
}

func NewStorage() *Storage {
	return &Storage{make(map[string]string)}
}

func (st *Storage) SaveUrl(url string) (string) {
	hash := md5.Sum([]byte(url))
	shortHash := hex.EncodeToString(hash[:4])
	st.db[shortHash] = url
	return shortHash
}

func (st *Storage) GetUrl(hash string) (string, error){
	if url, exists := st.db[hash]; exists {
		return url, nil
	}
	return "", errors.New("An URL with this hash doesn't exist ;(")
}

