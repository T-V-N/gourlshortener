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

func (db *DBStorage) SaveURL(ctx context.Context, url, uid, hash string) error {
	conn, err := pgx.Connect(ctx, db.cfg.DatabaseDSN)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err.Error())
		return err
	}

	defer conn.Close(ctx)

	sqlStatement := `
	INSERT INTO urls (uid, hash, original_url)
	VALUES ($1, $2, $3)`

	_, err = conn.Exec(ctx, sqlStatement, uid, hash, url)

	if err != nil {
		return err
	}

	return nil
}

func (db *DBStorage) GetURL(ctx context.Context, hash string) (string, error) {
	conn, err := pgx.Connect(ctx, db.cfg.DatabaseDSN)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err.Error())
		return "", err
	}

	defer conn.Close(ctx)

	row := conn.QueryRow(ctx, "Select original_url from urls where hash = $1", hash)

	var originalURL string
	err = row.Scan(&originalURL)

	if err != nil {
		return "", err
	}

	return originalURL, nil
}

func (db *DBStorage) GetUrlsByUID(ctx context.Context, uid string) ([]URL, error) {
	urls := make([]URL, 0)
	conn, err := pgx.Connect(ctx, db.cfg.DatabaseDSN)

	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err.Error())
		return urls, err
	}

	rows, err := conn.Query(ctx, "SELECT hash, original_url from urls where uid = $1", uid)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// пробегаем по всем записям
	for rows.Next() {
		var u URL
		err = rows.Scan(&u.ShortURL, &u.URL)

		if err != nil {
			return nil, err
		}

		urls = append(urls, u)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return urls, nil
}

func (db *DBStorage) IsAlive(ctx context.Context) (bool, error) {
	conn, err := pgx.Connect(ctx, db.cfg.DatabaseDSN)

	if err != nil {
		return false, err
	}

	defer conn.Close(ctx)

	return true, nil
}
