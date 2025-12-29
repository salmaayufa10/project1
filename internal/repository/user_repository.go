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
		RETURNING id;
	`
	err := r.db.QueryRow(
		query,
		s.Email,
		s.Name,
		s.Password,
	).Scan(&s.Id)
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
		FROM lib_user
		WHERE email = $1;
	`
	var s model.Lib_user
	if err := r.db.QueryRow(query, email).Scan(
		&s.Id,
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

func (r *UserRepository) GetUserByID(id int64) (*model.Lib_user, error) {
	const query = `
		SELECT id, email, name, password
		FROM lib_user
		WHERE id = $1;
	`
	var s model.Lib_user
	if err := r.db.QueryRow(query, id).Scan(
		&s.Id,
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

func (r *UserRepository) ListUser() ([]model.Lib_user, error) {
	rows, err := r.db.Query(`
		SELECT id, email, name
		FROM lib_user
		ORDER BY id;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.Lib_user
	for rows.Next() {
		var c model.Lib_user
		if err := rows.Scan(&c.Id, &c.Email, &c.Name); err != nil {
			return nil, err
		}
		users = append(users, c)
	}
	return users, nil
}
