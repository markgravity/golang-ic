package bootstrap

import (
	"github.com/markgravity/golang-ic/database"
	"github.com/markgravity/golang-ic/helpers/log"
	"github.com/markgravity/golang-ic/lib/services/oauth"
)

func Init() {
	LoadConfig()
	LoadENV()

	InitDatabase(database.GetDatabaseURL())

	err := oauth.SetUpOAuthServer()
	if err != nil {
		log.Errorf("Error when setting up OAuth server: %v", err)
	}
}
