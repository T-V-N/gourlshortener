// App is responsible for the business logic of the service
package app

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/url"
	"time"

	"github.com/T-V-N/gourlshortener/internal/config"
	"github.com/T-V-N/gourlshortener/internal/storage"
)

// App struct contains all the necessary objects for app to perform storage CRUD and the business logic proccess
type App struct {
	DB         storage.Storage            // file, db or memory-storage
	Config     *config.Config             // set of configs
	deleteChan chan storage.DeletionEntry // channel used by an URL deletion goroutine
}

// InitApp creates and returns an application from st storage and cfg config.
// Also inits a deletion channel and runs a deletion goroutine
func InitApp(st storage.Storage, cfg *config.Config) *App {
	delChan := make(chan storage.DeletionEntry)
	app := &App{st, cfg, delChan}

	go app.deletionConsumer(delChan)

	return app
}

func (app *App) deletionConsumer(ch chan storage.DeletionEntry) {
	buff := []storage.DeletionEntry{}
	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case el := <-ch:
			if len(buff) == 5 {
				err := app.DB.DeleteURLs(context.Background(), buff)

				if err != nil {
					log.Println(err)
				}

				buff = buff[:0]
			}

			buff = append(buff, storage.DeletionEntry{Hash: el.Hash, UID: el.UID})
		case <-ticker.C:
			err := app.DB.DeleteURLs(context.Background(), buff)

			if err != nil {
				log.Println(err)
			}

			buff = buff[:0]
		}
	}
}

func (app *App) SaveURL(rawURL, UID string, ctx context.Context) (string, error) {
	_, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return rawURL, err
	}

	hash := md5.Sum([]byte(rawURL))
	stringHash := hex.EncodeToString(hash[:4])

	err = app.DB.SaveURL(ctx, rawURL, UID, stringHash)

	if err != nil {
		return stringHash, err
	}

	return app.Config.BaseURL + "/" + stringHash, nil
}

// GetURL searches for and URL having id and if found returns it
func (app *App) GetURL(ctx context.Context, id string) (storage.URL, error) {
	u, err := app.DB.GetURL(ctx, id)

	if err != nil {
		return storage.URL{}, err
	}

	return u, nil
}

// GetURLByUID tries to search all URLs bound to a user with UID and returns a list of URLs
func (app *App) GetURLByUID(uid string, ctx context.Context) ([]storage.URL, error) {
	u, err := app.DB.GetUrlsByUID(ctx, uid)

	for i, el := range u {
		u[i].ShortURL = app.Config.BaseURL + "/" + el.ShortURL
	}

	if err != nil {
		return []storage.URL{}, err
	}

	return u, nil
}

// PingStorage just checks whether the app storage connection is alive or not
func (app *App) PingStorage(ctx context.Context) error {
	_, err := app.DB.IsAlive(ctx)
	if err != nil {
		return err
	}

	return nil
}

// BatchSaveURL takes a list of URLs and saves them binding to a user with UID
func (app *App) BatchSaveURL(ctx context.Context, obj []storage.BatchURL, uid string) ([]storage.BatchURL, error) {
	urls := []storage.URL{}
	responseURLs := []storage.BatchURL{}

	for _, rawURL := range obj {
		u, err := url.ParseRequestURI(rawURL.OriginalURL)
		if err != nil {
			continue
		}

		hash := rawURL.CorrelationID

		urls = append(urls, storage.URL{UID: uid, ShortURL: hash, URL: u.String()})
		responseURLs = append(responseURLs, storage.BatchURL{OriginalURL: "", CorrelationID: hash, ShortURL: app.Config.BaseURL + "/" + hash})
	}

	err := app.DB.BatchSaveURL(ctx, urls)
	if err != nil {
		return nil, err
	}

	return responseURLs, nil
}

// DeleteListURL stages rawHashes list containing hashes of urls for deletion.
func (app *App) DeleteListURL(ctx context.Context, rawHashes []string, uid string) error {
	go func() {
		for _, rawHash := range rawHashes {
			app.deleteChan <- storage.DeletionEntry{Hash: rawHash, UID: uid}
		}
	}()

	return nil
}
