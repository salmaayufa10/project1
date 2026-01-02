package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"library/internal/model"

	"github.com/lib/pq"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(s *model.Lib_user) error {
	const query = `
		INSERT INTO lib_user (email, name, password)
		VALUES ($1, $2, $3)
	`
	_, err := r.db.Exec(
		query,
		s.Email,
		s.Name,
		s.Password,
	)

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return fmt.Errorf("email already used")
		}
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (*model.Lib_user, error) {
	const query = `
		SELECT email, name, password
		FROM users
		WHERE email = $1;
	`
	var s model.Lib_user
	if err := r.db.QueryRow(query, email).Scan(
		&s.Email,
		&s.Name,
		&s.Password,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &s, nil
}
