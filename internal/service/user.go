package service

import (
	"context"

	"github.com/xpzouying/go-clean-code-template/api"
	"github.com/xpzouying/go-clean-code-template/internal/usecase"

	"go.uber.org/zap"
)

// UserService is the user service.
type UserService struct {
	userUC *usecase.UserUsecase

	logger *zap.SugaredLogger
}

// NewUserService creates a new user service.
func NewUserService(userUC *usecase.UserUsecase) *UserService {
	return &UserService{
		userUC: userUC,
		logger: zap.S().Named("service"),
	}
}

// CreateUser create a new user.
func (us *UserService) CreateUser(ctx context.Context, req *api.CreateUserReq) (*api.CreateUserReply, error) {

	user, err := us.userUC.CreateUser(ctx, req.Username, req.AvatarURL)
	if err != nil {
		return nil, err
	}

	return &api.CreateUserReply{
		UID: user.Uid,
	}, nil
}