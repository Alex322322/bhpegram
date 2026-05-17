package repository

import (
	"context"
	"fmt"

	"github.com/Alex322322/bhpegram/internal/domain"
	"github.com/jmoiron/sqlx"
)

var _ UserRepository = (*PostgresUserRepository)(nil)

type PostgresUserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(ctx context.Context, req *domain.CreateUserRequest) (*domain.User, error) {
	query := `
	INSERT INTO users (username, phone_number, email, password_hash) 
	VALUES ($1, $2, $3, $4)
	RETURNING *`

	var user domain.User
    err := r.db.QueryRowxContext(ctx, query, req.Username, req.PhoneNumber, req.Email, req.Password).StructScan(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user data to database: %w", err)
	}
	return &user, nil
}

func (r *PostgresUserRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user from database: %w", err)
	}
	return nil
}

func (r *PostgresUserRepository) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := `
	UPDATE users
	SET username = $2, phone_number = $3, email = $4, password_hash = $5, updated_at = NOW()
	WHERE id = $1
	RETURNING *`

	var updatedUser domain.User
	err := r.db.QueryRowxContext(ctx, query, user.ID, user.Username, user.PhoneNumber, user.Email, user.PasswordHash).StructScan(&updatedUser)
	if err != nil {
		return nil, fmt.Errorf("failed to update user from database: %w", err)
	}
	return &updatedUser, nil
}

func (r *PostgresUserRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	query := `SELECT * FROM users WHERE id = $1`

	var user domain.User
	err := r.db.QueryRowxContext(ctx, query, id).StructScan(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	return &user, nil
}

func (r *PostgresUserRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	query := `SELECT * FROM users WHERE username = $1`

	var user domain.User
	err := r.db.QueryRowxContext(ctx, query, username).StructScan(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}
	return &user, nil
}

func (r *PostgresUserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT * FROM users WHERE email = $1`

	var user domain.User
	err := r.db.QueryRowxContext(ctx, query, email).StructScan(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	return &user, nil
}
