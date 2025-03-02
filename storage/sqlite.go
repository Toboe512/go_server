package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"go_server/storage/dao"
	"os"
	"strings"
)

type Storage struct {
	db *sql.DB
}

const defaultPerm = 0774

var ErrNoLoadData = errors.New("no load data")

var DB *Storage

func New(path string) (*Storage, error) {

	splitPath := strings.Split(path, "/")
	basePath := strings.Join(splitPath[:len(splitPath)-1], "/")
	if err := os.MkdirAll(basePath, defaultPerm); err != nil {
		return nil, err
	}
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		return nil, fmt.Errorf("can't open database sqlite: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("can't open database sqlite: %w", err)
	}

	return &Storage{db: db}, nil
}

func (s Storage) GetById(ctx context.Context, id string) (*dao.Data, error) {
	q := `SELECT data FROM usr_data WHERE id = ?`

	var data string

	err := s.db.QueryRowContext(ctx, q, id).Scan(&data)
	if err == sql.ErrNoRows {
		return nil, ErrNoLoadData
	}
	if err != nil {
		return nil, fmt.Errorf("can't get data by id: %w", err)
	}

	if err != nil {
		return nil, err
	}

	return &dao.Data{
		ID:   id,
		Data: data,
	}, nil
}

func (s Storage) Save(ctx context.Context, data *dao.Data) error {

	q := `INSERT INTO usr_data (id, data) VALUES (?, ?)`

	if _, err := s.db.ExecContext(ctx, q, data.ID, data.Data); err != nil {
		return fmt.Errorf("can't save data in sqlite: %w", err)
	}

	return nil

}

func (s Storage) Delete(ctx context.Context, id string) error {
	q := `DELETE FROM usr_data WHERE id = ?`

	if _, err := s.db.ExecContext(ctx, q, id); err != nil {
		return fmt.Errorf("can't remove data in sqlite: %w", err)
	}
	return nil
}

func (s Storage) Init(ctx context.Context) error {
	//TODO определить столбцы таблицы
	q := `CREATE TABLE IF NOT EXISTS usr_data (id TEXT UNIQUE, data TEXT)`

	_, err := s.db.ExecContext(ctx, q)

	if err != nil {
		return fmt.Errorf("can't create table in sqlite: %w", err)
	}

	return nil
}
