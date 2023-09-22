package database

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/skyris/auth-server/internal/database/dto"
	"github.com/skyris/auth-server/internal/database/transaction"
)

func (r *Repository) Register(username, email, password string) (*dto.User, error) {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)
RETURNING id, username, email, password, active, created_at, modified_at;`
	var u dto.User
	ctx := context.Background()
	tx, err := r.db.Begin(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer func(ctx context.Context, tx pgx.Tx) {
		err = transaction.Finish(ctx, tx, err)
	}(ctx, tx)

	row := tx.QueryRow(ctx, query, username, email, password)
	err = row.Scan(&u.Id, &u.Username, &u.Email, &u.Password, &u.Active, &u.CreatedAt, &u.ModifiedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &u, nil
}

func (r *Repository) GetUserByEmail(email string) (*dto.User, error) {
	query := `SELECT id, username, email, password
FROM users
WHERE email=$1;`
	var u dto.User
	ctx := context.Background()
	tx, err := r.db.Begin(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer func(ctx context.Context, tx pgx.Tx) {
		err = transaction.Finish(ctx, tx, err)
	}(ctx, tx)

	row := tx.QueryRow(ctx, query, email)
	err = row.Scan(&u.Id, &u.Username, &u.Email, &u.Password)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &u, nil
}

func (r *Repository) UpdateUser(id uuid.UUID) (*dto.User, error) {

	return new(dto.User), nil
}

func (r *Repository) SoftDeleteUser(id uuid.UUID) error {
	query := `UPDATE users SET active=FALSE
WHERE id=$1;`
	ctx := context.Background()
	tx, err := r.db.Begin(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	defer func(ctx context.Context, tx pgx.Tx) {
		err = transaction.Finish(ctx, tx, err)
	}(ctx, tx)

	_, err = tx.Exec(ctx, query, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
func (r *Repository) HardDeleteUser(id uuid.UUID) error {
	query := `DELETE FROM users
WHERE id=$1;`
	ctx := context.Background()
	tx, err := r.db.Begin(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	defer func(ctx context.Context, tx pgx.Tx) {
		err = transaction.Finish(ctx, tx, err)
	}(ctx, tx)
	_, err = tx.Exec(ctx, query, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
