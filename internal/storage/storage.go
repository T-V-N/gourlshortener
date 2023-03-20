package storage

import (
	"context"

	"github.com/T-V-N/gourlshortener/internal/config"
)

// DeletionEntry is a struct used for url deletion
type DeletionEntry struct {
	UID  string // user id
	Hash string // url hash
}

// URL struct describes URL obj and its json format
type URL struct {
	UID       string `json:"-"`            // user uid, ommited in JSON responses
	ShortURL  string `json:"short_url"`    // url hash
	URL       string `json:"original_url"` // full url
	IsDeleted bool   // flag if a url was deleted
}

// BatchURL used for URL lists
type BatchURL struct {
	OriginalURL   string `json:"original_url,omitempty"` // full url
	CorrelationID string `json:"correlation_id"`         // url hash
	ShortURL      string `json:"short_url"`              // link to a server which redirects to the original url
}

// Storage is the main interface used by app for storing URLs
type Storage interface {
	SaveURL(ctx context.Context, url, uid, hash string) error    // Saves an URL to a storage
	GetURL(ctx context.Context, hash string) (URL, error)        // Returns an URL from a storage
	GetUrlsByUID(ctx context.Context, uid string) ([]URL, error) // Returns all URLs belonging to a user with uid
	IsAlive(ctx context.Context) (bool, error)                   // Checks if storage is alive
	BatchSaveURL(ctx context.Context, urls []URL) error          // Saves a list of urls to a storage
	KillConn() error                                             // Gracefully stops a storage connection
	DeleteURLs(context.Context, []DeletionEntry) error           // Deletes URLs from storage
	GetStats(context.Context) (users, urls int, err error)       // Get amount of users, urls in DB
}

// InitStorage creates a storage based on file saving strategy (file or db) and returns it
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
