package app

import (
	"net/url"
	"strings"

	"github.com/T-V-N/gourlshortener/internal/storage"
)

type App struct {
	db *storage.Storage
}

func InitApp(st *storage.Storage) *App {
	return &App{st}
}

func (app *App) SaveURL(URL string) (string, error){
	u, err := url.ParseRequestURI(URL)
	if err != nil {
		return URL, err
	}
	hash, err := app.db.SaveUrl(strings.ToLower(u.String()))
	if err != nil {
		return u.String(), err
	}
	return hash, nil
}

func (app *App) GetUrl(id string) (string, error){
	url, err := app.db.GetUrl(id)
	if err != nil {
		return id, err
	}
	return url, nil
}
