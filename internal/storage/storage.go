package storage

import (
	"context"

	"github.com/T-V-N/gourlshortener/internal/config"
)

type DeletionEntry struct {
	UID  string
	Hash string
}

type URL struct {
	UID       string `json:"-"`
	ShortURL  string `json:"short_url"`
	URL       string `json:"original_url"`
	IsDeleted bool
}

type BatchURL struct {
	OriginalURL   string `json:"original_url,omitempty"`
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}

type Storage interface {
	SaveURL(ctx context.Context, url, uid, hash string) error
	GetURL(ctx context.Context, hash string) (URL, error)
	GetUrlsByUID(ctx context.Context, uid string) ([]URL, error)
	IsAlive(ctx context.Context) (bool, error)
	BatchSaveURL(ctx context.Context, urls []URL) error
	KillConn() error
	DeleteURLs(context.Context, []DeletionEntry) error
}

func InitStorage(data map[string]URL, cfg *config.Config) Storage {
	if cfg.DatabaseDSN != "" {
		storage, err := InitDBStorage(cfg)
		if err != nil {
			return InitFileStorage(data, cfg)
		}

		return storage
	}

	return InitFileStorage(data, cfg)
}
