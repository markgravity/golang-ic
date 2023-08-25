package forms

import (
	"github.com/markgravity/golang-ic/database"
	"github.com/markgravity/golang-ic/helpers"
	"github.com/markgravity/golang-ic/helpers/log"
	"github.com/markgravity/golang-ic/lib/models"
	
	"strings"
)

type SignUpForm struct {
	Email                string `form:"email" binding:"required,email"`
	Password             string `form:"password" binding:"required,min=6,confirmed"`
	PasswordConfirmation string `form:"password_confirmation" binding:"required"`
}

func (f *SignUpForm) Save() (*models.User, error) {
	hashedPassword, err := helpers.HashPassword(f.Password)
	if err != nil {
		log.Error("Encryption error:", err)
		return nil, err
	}

	user := &models.User{
		Email:             strings.ToLower(f.Email),
		EncryptedPassword: hashedPassword,
	}

	db := database.GetDB()
	err = user.Create(db)
	if err != nil {
		log.Error("Create error:", err)
		return nil, err
	}

	return user, err
}
