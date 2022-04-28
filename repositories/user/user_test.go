package user

import (
	"final-project/configs"
	U "final-project/entities/user"
	SeederUser "final-project/repositories/mocks/user"
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

func TestCreate(t *testing.T) {
	Migrator()
	repo := NewUserRepository(db)

	mockUser := SeederUser.UserSeeder()

	t.Run("positive", func(t *testing.T) {
		res, err := repo.Create(mockUser)
		assert.Nil(t, err)
		assert.Equal(t, mockUser.Name, res.Name)
	})

	t.Run("negative", func(t *testing.T) {
		_, err := repo.Create(mockUser)
		assert.NotNil(t, err)
	})
}

func TestGet(t *testing.T) {
	Migrator()
	repo := NewUserRepository(db)

	mockUser := SeederUser.UserSeeder()

	t.Run("negative", func(t *testing.T) {
		_, err := repo.Get(1)
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		repo.Create(mockUser)

		res, err := repo.Get(1)
		assert.Nil(t, err)
		assert.Equal(t, mockUser.Name, res.Name)
	})
}

func TestGetByID(t *testing.T) {
	Migrator()
	repo := NewUserRepository(db)

	mockUser := SeederUser.UserSeeder()

	t.Run("negative", func(t *testing.T) {
		_, err := repo.GetByID(1)
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		repo.Create(mockUser)

		res, err := repo.GetByID(1)
		assert.Nil(t, err)
		assert.Equal(t, mockUser.Name, res.Name)
	})
}

func TestGetAllUsers(t *testing.T) {
	Migrator()
	repo := NewUserRepository(db)

	mockUser := SeederUser.UserSeeder()

	t.Run("negative", func(t *testing.T) {
		_, err := repo.GetAllUsers()
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		repo.Create(mockUser)

		res, err := repo.GetAllUsers()
		assert.Nil(t, err)
		assert.Equal(t, mockUser.Name, res[0].Name)
	})
}

func TestUpdate(t *testing.T) {
	Migrator()
	repo := NewUserRepository(db)

	mockUser := SeederUser.UserSeeder()

	t.Run("positive", func(t *testing.T) {
		repo.Create(mockUser)

		mockUser2 := SeederUser.UserSeeder()
		mockUser2.ID = 1
		mockUser2.Name = "ucup_Updated"

		res, err := repo.Update(mockUser2)
		assert.Nil(t, err)
		assert.Equal(t, mockUser2.Name, res.Name)
	})

	t.Run("negative", func(t *testing.T) {
		mockUser2 := SeederUser.UserSeeder()

		_, err := repo.Update(mockUser2)
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	Migrator()
	repo := NewUserRepository(db)

	mockUser := SeederUser.UserSeeder()

	t.Run("negative", func(t *testing.T) {
		err := repo.Delete(1)
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		repo.Create(mockUser)

		err := repo.Delete(1)
		assert.Nil(t, err)
	})
}
