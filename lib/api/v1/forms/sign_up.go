package forms

import (
	"github.com/markgravity/golang-ic/database"
	"github.com/markgravity/golang-ic/helpers"
	"github.com/markgravity/golang-ic/helpers/log"
	"github.com/markgravity/golang-ic/lib/models"
	"github.com/markgravity/golang-ic/lib/validators"
)

type SignUpForm struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=6"`
}

func (f *SignUpForm) Validate() error {
	return validators.Validate(f)
}

func (f *SignUpForm) Save() (*models.User, error) {
	err := f.Validate()
	if err != nil {
		log.Error("Validation error:", err)
		return nil, err
	}

	hashedPassword, err := helpers.HashPassword(f.Password)
	if err != nil {
		log.Error("Encryption error:", err)
		return nil, err
	}

	user := &models.User{
		Email:             f.Email,
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
