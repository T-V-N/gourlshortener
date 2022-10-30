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

type App struct {
	DB         storage.Storage
	Config     *config.Config
	deleteChan chan storage.DeletionEntry
}

func InitApp(st storage.Storage, cfg *config.Config) *App {
	delChan := make(chan storage.DeletionEntry)
	app := &App{st, cfg, delChan}

	go app.deletionConsumer(delChan)

	return app
}

func (app *App) deletionConsumer(ch chan storage.DeletionEntry) {
	buff := []storage.DeletionEntry{}
	ticker := time.NewTicker(30 * time.Second)
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
	u, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return rawURL, err
	}

	hash := md5.Sum([]byte(u.String()))
	stringHash := hex.EncodeToString(hash[:4])

	err = app.DB.SaveURL(ctx, u.String(), UID, stringHash)

	if err != nil {
		return stringHash, err
	}

	return app.Config.BaseURL + "/" + stringHash, nil
}

func (app *App) GetURL(id string, ctx context.Context) (storage.URL, error) {
	u, err := app.DB.GetURL(ctx, id)

	if err != nil {
		return storage.URL{}, err
	}

	return u, nil
}

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

func (app *App) PingStorage(ctx context.Context) error {
	_, err := app.DB.IsAlive(ctx)
	if err != nil {
		return err
	}

	return nil
}

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

func (app *App) DeleteListURL(ctx context.Context, rawHashes []string, uid string) error {
	go func() {
		for _, rawHash := range rawHashes {
			app.deleteChan <- storage.DeletionEntry{Hash: rawHash, UID: uid}
		}
	}()

	return nil
}
