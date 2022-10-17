package storage

import (
	"context"
	"log"

	"github.com/T-V-N/gourlshortener/internal/config"
	"github.com/jackc/pgx/v5"
)

type DBStorage struct {
	conn *pgx.Conn
	cfg  config.Config
}

func InitDBStorage(cfg *config.Config) (*DBStorage, error) {
	conn, err := pgx.Connect(context.Background(), cfg.DatabaseDSN)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err.Error())
		return nil, err
	}

	_, err = conn.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS 
	URLS 
	(uid varchar, hash varchar, original_url varchar)
	`)

	if err != nil {
		log.Printf("Unable to create db: %v\n", err.Error())
		return nil, err
	}

	defer conn.Close(context.Background())

	return &DBStorage{conn, *cfg}, nil
}

func (db *DBStorage) SaveURL(ctx context.Context, url, uid string) (string, error) {
	return "aa", nil
}

func (db *DBStorage) GetURL(ctx context.Context, hash string) (string, error) {
	return "aa", nil
}

func (db *DBStorage) GetUrlsByUID(ctx context.Context, uid string) ([]URL, error) {
	return []URL{URL{"aaa", "aaa", "a"}}, nil
}

func (db *DBStorage) IsAlive(ctx context.Context) (bool, error) {
	err := db.conn.Ping(ctx)
	if err != nil {
		return false, err
	}

	return true, nil
}
