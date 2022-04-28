package auth

import (
	"final-project/configs"
	"final-project/deliveries/helpers/hash"
	U "final-project/entities/user"
	SeederUser "final-project/repositories/mocks/user"
	"final-project/repositories/user"
	"final-project/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	config = configs.GetConfig(true)
	db     = utils.InitDB(config)
)

func Migrator() {
	db.Migrator().DropTable(&U.Users{})

	db.AutoMigrate(&U.Users{})
}

func TestLogin(t *testing.T) {
	Migrator()
	userRepo := user.NewUserRepository(db)
	repo := NewAuthRepository(db)
	mockUser := SeederUser.UserSeeder()

	t.Run("negative", func(t *testing.T) {
		_, err := repo.Login(mockUser.Email, mockUser.Password)
		assert.NotNil(t, err)
	})

	t.Run("success to login", func(t *testing.T) {
		tempPassword := mockUser.Password
		mockUser.Password, _ = hash.HashPassword(mockUser.Password)
		userRepo.Create(mockUser)
		_, err := repo.Login(mockUser.Email, tempPassword)
		assert.Nil(t, err)
	})
}
