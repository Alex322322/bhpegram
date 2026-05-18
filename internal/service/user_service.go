package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Alex322322/bhpegram/internal/domain"
	"github.com/Alex322322/bhpegram/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var _ UserService = (*userService)(nil)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *userService {
	return &userService{repo: repo}
}

func (s *userService) Create(ctx context.Context, user *domain.CreateUserRequest) (*domain.User, error) {
	u, err := s.repo.GetByEmail(ctx, user.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("failed to check email: %w", err)
	}
	if u != nil {
		return nil, fmt.Errorf("failed to create user, user with this email already exists")
	}

	u, err = s.repo.GetByUsername(ctx, user.Username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("failed to check username: %w", err)
	}
	if u != nil {
		return nil, fmt.Errorf("failed to create user, user with this username already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed generate password hash: %w", err)
	}

	user.Password = string(hash)
	u, err = s.repo.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return u, nil
}

func (s *userService) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	u, err := s.repo.GetByID(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}
	return u, nil
}

func (s *userService) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	u, err := s.repo.GetByUsername(ctx, username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}
	return u, nil
}

func (s *userService) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	u, err := s.repo.GetByEmail(ctx, email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	return u, nil
}

func (s *userService) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	u, err := s.repo.Update(ctx, user)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}
	return u, nil
}

func (s *userService) Delete(ctx context.Context, id int64) error {
	err := s.repo.Delete(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
