package fabricators

import (
	"github.com/markgravity/golang-ic/database"
	"github.com/markgravity/golang-ic/lib/models"

	"github.com/onsi/ginkgo"
)

func FabricateKeyword(keyword string, user *models.User) *models.Keyword {
	db := database.GetDB()
	keywordObject := &models.Keyword{
		Text: keyword,
		User: user,
	}

	err := keywordObject.Save(db)
	if err != nil {
		ginkgo.Fail("Fabricate keyword error:" + err.Error())
	}

	return keywordObject
}
