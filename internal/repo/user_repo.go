package repo

import (
	"context"
	"errors"

	"github.com/xpzouying/go-clean-code-template/internal/domain"

	"gorm.io/gorm"
)

type userPO struct {
	gorm.Model

	Name   string `gorm:"column:name"`
	Avatar string `gorm:"column:avatar"`
}

func (userPO) TableName() string {

	return "user"
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) domain.UserRepo {
	_ = db.AutoMigrate(&userPO{})

	return &userRepo{db}
}

func (rp *userRepo) CreateUser(ctx context.Context, name, avatar string) (*domain.User, error) {

	r := userPO{
		Name:   name,
		Avatar: avatar,
	}

	if err := rp.db.WithContext(ctx).Create(&r).Error; err != nil {
		return nil, err
	}

	do := rp.convertToUser(r)
	return &do, nil
}

// GetUser get user by uid.
func (rp *userRepo) GetUser(ctx context.Context, uid int) (*domain.User, bool, error) {

	var r userPO
	if err := rp.db.WithContext(ctx).Where("id = ?", uid).First(&r).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}

		return nil, false, err
	}

	do := rp.convertToUser(r)
	return &do, true, nil
}

func (rp *userRepo) convertToUser(r userPO) domain.User {
	return domain.User{
		Uid:    int(r.ID),
		Name:   r.Name,
		Avatar: r.Avatar,
	}
}
