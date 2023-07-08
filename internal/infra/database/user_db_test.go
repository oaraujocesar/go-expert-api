package database

import (
	"testing"

	"github.com/oaraujocesar/go-expert-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db.AutoMigrate(&entity.User{})

	user, err := entity.NewUser("John", "john@doe.com", "123456")
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating an user entity", err)
	}

	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error

	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotEmpty(t, userFound.Password)
	assert.NotNil(t, userFound.Password)
}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db.AutoMigrate(&entity.User{})

	t.Run("It should get user by email", func(t *testing.T) {
		user, err := entity.NewUser("John", "john@doe.com", "123456")
		assert.Nil(t, err)

		userDB := NewUser(db)
		userDB.Create(user)

		userFound, err := userDB.FindByEmail(user.Email)

		assert.Nil(t, err)
		assert.Equal(t, user.ID, userFound.ID)
		assert.Equal(t, user.Name, userFound.Name)
		assert.Equal(t, user.Email, userFound.Email)
		assert.NotEmpty(t, userFound.Password)
		assert.NotNil(t, userFound.Password)
	})

	t.Run("It should return error when user not found", func(t *testing.T) {
		userDB := NewUser(db)
		user, err := userDB.FindByEmail("hey@notfound.com")

		assert.Nil(t, user)
		assert.NotNil(t, err)
	})
}
