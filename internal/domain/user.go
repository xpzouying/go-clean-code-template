package domain

import (
	"context"
)

type User struct {
	Uid    int
	Name   string
	Avatar string
}

type UserRepo interface {
	CreateUser(ctx context.Context, name, avatar string) (*User, error)
	GetUser(ctx context.Context, uid int) (*User, bool, error)
}

type UserDomainService interface {
	// CreateUser create a new user.
	CreateUser(ctx context.Context, name, avatar string) (*User, error)

	// GetUser get user by uid. If the user does not exist, the second return value is false.
	GetUser(ctx context.Context, uid int) (*User, bool, error)
}

type userDomainService struct {
	repo UserRepo
}

func NewUserDomainService(repo UserRepo) UserDomainService {
	return &userDomainService{repo: repo}
}

func (s *userDomainService) CreateUser(ctx context.Context, name, avatar string) (*User, error) {
	return s.repo.CreateUser(ctx, name, avatar)
}

func (s *userDomainService) GetUser(ctx context.Context, uid int) (*User, bool, error) {
	user, exists, err := s.repo.GetUser(ctx, uid)
	if err != nil {
		return nil, false, err
	}
	if !exists {
		return nil, false, nil
	}

	return user, true, nil
}
