package domain

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	UserTable = "users"
)

type User struct {
	ID       uuid.UUID `gorm:"primaryKey;type:char(36)" json:"id"`
	Name     string    `gorm:"size:255;index" json:"name"`
	Email    string    `gorm:"unique;size:255;index" json:"email"`
	Password string    `json:"-"`
	IsActive bool      `gorm:"index" json:"is_active"`
	Role     string    `gorm:"size:16;index" json:"role"`
}

type RegisterUser struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

func (ru *RegisterUser) ToUser() (user User, err error) {
	if ru.Password != ru.ConfirmPassword {
		return User{}, errors.New("password not match")
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(ru.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return User{}, err
	}

	user.ID, err = uuid.NewUUID()
	if err != nil {
		return User{}, err
	}
	user.Name = ru.Name
	user.Email = ru.Email
	user.Password = string(encryptedPassword)
	user.IsActive = false
	user.Role = "admin"
	return
}

type UserRepository interface {
	Create(c context.Context, user User) error
	Retrieve(c context.Context, filter Filter) (users []User, meta MetaResponse, err error)
	Update(c context.Context, id uuid.UUID, data User) (user User, err error)
	Delete(c context.Context, id uuid.UUID) error
}

type UserUsecase interface {
	Create(c context.Context, user User) error
	Retrieve(c context.Context, filter Filter) (users []User, meta MetaResponse, err error)
	Update(c context.Context, id uuid.UUID, data User) (user User, err error)
	Delete(c context.Context, id uuid.UUID) error
}