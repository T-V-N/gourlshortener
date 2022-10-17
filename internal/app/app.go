package app

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"net/url"

	"github.com/T-V-N/gourlshortener/internal/config"
	"github.com/T-V-N/gourlshortener/internal/storage"
)

type App struct {
	db     storage.Storage
	Config *config.Config
}

func InitApp(st storage.Storage, cfg *config.Config) *App {
	return &App{st, cfg}
}

func (app *App) SaveURL(rawURL, UID string, ctx context.Context) (string, error) {
	u, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return rawURL, err
	}

	hash := md5.Sum([]byte(u.String()))
	stringHash := hex.EncodeToString(hash[:4])
	err = app.db.SaveURL(ctx, u.String(), UID, stringHash)

	if err != nil {
		return u.String(), err
	}

	return app.Config.BaseURL + "/" + stringHash, nil
}

func (app *App) GetURL(id string, ctx context.Context) (string, error) {
	u, err := app.db.GetURL(ctx, id)

	if err != nil {
		return id, err
	}

	return u, nil
}

func (app *App) GetURLByUID(uid string, ctx context.Context) ([]storage.URL, error) {
	u, err := app.db.GetUrlsByUID(ctx, uid)

	for i, el := range u {
		u[i].ShortURL = app.Config.BaseURL + "/" + el.ShortURL
	}

	if err != nil {
		return []storage.URL{}, err
	}

	return u, nil
}

func (app *App) PingStorage(ctx context.Context) error {
	_, err := app.db.IsAlive(ctx)
	if err != nil {
		return err
	}

	return nil
}
