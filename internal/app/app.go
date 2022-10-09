package app

import (
	"net/url"
	"strings"

	"github.com/T-V-N/gourlshortener/internal/config"
	"github.com/T-V-N/gourlshortener/internal/storage"
)

type App struct {
	db     *storage.Storage
	Config *config.Config
}

func InitApp(st *storage.Storage, cfg *config.Config) *App {
	return &App{st, cfg}
}

func (app *App) SaveURL(rawURL, UID string) (string, error) {
	u, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return rawURL, err
	}

	hash, err := app.db.SaveURL(strings.ToLower(u.String()), UID)

	if err != nil {
		return u.String(), err
	}

	return hash, nil
}

func (app *App) GetURL(id string) (string, error) {
	u, err := app.db.GetURL(id)

	if err != nil {
		return id, err
	}

	return u, nil
}
func (app *App) GetURLByUID(uid string) ([]storage.URL, error) {
	u, err := app.db.GetUrlsByUID(uid)

	if err != nil {
		return []storage.URL{}, err
	}

	return u, nil
}
