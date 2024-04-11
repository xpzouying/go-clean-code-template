package repo

import (
	"context"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/xpzouying/go-clean-code-template/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestCreateUser(t *testing.T) {

	rp := NewUserRepo(newTestDB())

	var (
		ctx = context.Background()
	)

	t.Run("GetNotExistsUser_ShouldNotExists", func(t *testing.T) {
		uidNotEixists := 1024
		user, exists, err := rp.GetUser(ctx, uidNotEixists)
		assert.NoError(t, err)

		assert.False(t, exists)
		assert.Nil(t, user)
	})

	var user *domain.User

	t.Run("CreateUser_ShouldSuccess", func(t *testing.T) {

		var err error
		user, err = rp.CreateUser(ctx, "name", "avatar")
		assert.NoError(t, err)

		assert.Equal(t, "name", user.Name)
		assert.Equal(t, "avatar", user.Avatar)
	})

	t.Run("GetExistsUser_ShouldExists", func(t *testing.T) {

		got, exists, err := rp.GetUser(ctx, user.Uid)
		assert.NoError(t, err)

		assert.True(t, exists)

		assert.Equal(t, user, got)
	})

}

func newTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	return db
}
