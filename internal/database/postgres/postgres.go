package postgres

import (
	"context"
	"fmt"
	"github.com/skyris/auth-server/internal/env"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

const (
	defaultConnTimeout = time.Second * 5
)

type Settings struct {
	Host        string
	Port        string
	Database    string
	User        string
	Password    string
	SSLMode     string
	TimeZone    string
	ConnTimeout time.Duration
}

func NewSettings() Settings {
	s := Settings{
		Host:        env.DB_HOST,
		Port:        env.DB_PORT,
		Database:    env.DB_NAME,
		User:        env.DB_USER,
		Password:    env.DB_PASSWORD,
		SSLMode:     env.DB_SSLMODE,
		TimeZone:    env.DB_TIMEZONE,
		ConnTimeout: defaultConnTimeout,
	}

	return s
}

// ToDSN method: Data Source Name
func (s Settings) ToDSN() string {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		s.Host,
		s.User,
		s.Password,
		s.Database,
		s.Port,
		s.SSLMode,
		s.TimeZone,
	)

	return dsn
}

type Store struct {
	Pool *pgxpool.Pool
}

func New(settings Settings) (*Store, error) {
	config, err := pgxpool.ParseConfig(settings.ToDSN())
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), settings.ConnTimeout)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return &Store{Pool: pool}, nil
}
