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

	return &DBStorage{conn, *cfg}, nil
}

func (st *DBStorage) SaveURL(url, UID string) (string, error) {
	return "hey", nil
}

func (st *DBStorage) GetURL(hash string) (string, error) {
	return "hey", nil
}

func (st *DBStorage) GetUrlsByUID(uid string) ([]URL, error) {
	return []URL{URL{"Hey", "HEY", "HEHEHEY"}}, nil
}

func (st *DBStorage) IsAlive() (bool, error) {
	err := st.conn.Ping(context.Background())
	if err != nil {
		return false, err
	}
	return true, nil
}
