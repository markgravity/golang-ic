package bootstrap

import (
	"github.com/markgravity/golang-ic/database"
	"github.com/markgravity/golang-ic/helpers/log"
	"github.com/markgravity/golang-ic/services/oauth"
)

func Init() {
	LoadConfig()
	LoadENV()

	InitDatabase(database.GetDatabaseURL())

	err := oauth.SetUpOAuthServer()
	if err != nil {
		log.Error("Error when setting up OAuth server: %v", err.Error())
	}
}
