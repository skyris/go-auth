package database

import (
	"fmt"
	"github.com/skyris/auth-server/internal/env"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pressly/goose"
)

type Options struct {
	DefaultLimit  uint64
	DefaultOffset uint64
}

type Repository struct {
	db      *pgxpool.Pool
	options Options
}

func New(db *pgxpool.Pool, options Options) (*Repository, error) {
	if err := migrations(db); err != nil {
		return nil, err
	}
	r := &Repository{
		db: db,
	}

	r.SetOptions(options)
	return r, nil
}

func (r *Repository) SetOptions(options Options) {
	if r.options.DefaultLimit == 0 {
		r.options.DefaultLimit = 10
	}
	if r.options != options {
		r.options = options
	}
}

func migrations(pool *pgxpool.Pool) (err error) {
	db, err := goose.OpenDBWithDriver("postgres", pool.Config().ConnConfig.ConnString())
	if err != nil {
		return err
	}

	defer func() {
		err = db.Close()
	}()

	goose.SetTableName("auth_version")
	if err := goose.Run("up", db, env.MIGRATIONS_DIR); err != nil {
		return fmt.Errorf("goose %s error: %w", "up", err)
	}

	return nil
}

func Ls(name string) {
	entries, err := os.ReadDir(name)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		log.Println(e.Name())
	}
}
