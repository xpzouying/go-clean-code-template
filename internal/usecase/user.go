// Package usecase: This file contains the implementation of the application service for the user domain.
package usecase

import (
	"context"

	"github.com/xpzouying/go-clean-code-template/internal/domain"

	"go.uber.org/zap"
)

type UserUsecase struct {
	userService domain.UserDomainService

	logger *zap.SugaredLogger
}

// NewUserUsecase creates a new application service: UserService.
func NewUserUsecase(userService domain.UserDomainService) *UserUsecase {
	return &UserUsecase{
		userService: userService,
		logger:      zap.S().Named("usecase"),
	}
}

func (s *UserUsecase) CreateUser(ctx context.Context, name, avatar string) (*domain.User, error) {

	user, err := s.userService.CreateUser(ctx, name, avatar)
	if err != nil {
		s.logger.Errorf("Failed to create user: %v", err)
		return nil, err
	}

	s.logger.Infof("User created: uid=%v", user.Uid)

	return user, nil
}
