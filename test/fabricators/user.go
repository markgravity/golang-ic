package fabricators

import (
	"github.com/markgravity/golang-ic/database"
	"github.com/markgravity/golang-ic/helpers"
	"github.com/markgravity/golang-ic/lib/models"

	"github.com/onsi/ginkgo"
)

func FabricateUser(email string, password string) *models.User {
	db := database.GetDB()

	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		ginkgo.Fail("Hashing password error:" + err.Error())
	}
	user := &models.User{Email: email}
	user.EncryptedPassword = hashedPassword

	err = user.Create(db)
	if err != nil {
		ginkgo.Fail("Create error:" + err.Error())
	}

	return user
}

func FabricateTester() *models.User {
	return FabricateUser("test@gmail.com", "123456")
}
