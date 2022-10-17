package storage

import (
	"fmt"

	"github.com/T-V-N/gourlshortener/internal/config"
)

type URL struct {
	UID      string `json:"-"`
	ShortURL string `json:"short_url"`
	URL      string `json:"original_url"`
}

type Storage interface {
	SaveURL(url, uid string) (string, error)
	GetURL(hash string) (string, error)
	GetUrlsByUID(uid string) ([]URL, error)
	IsAlive() (bool, error)
}

func InitStorage(data map[string]URL, cfg *config.Config) Storage {
	if cfg.DatabaseDSN == "" {
		fmt.Println("File storage")
		return InitFileStorage(data, cfg)
	}
	return InitFileStorage(data, cfg)
}
