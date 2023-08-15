package bootstrap

import (
	"github.com/markgravity/golang-ic/database"
)

func InitDatabase(databaseURL string) {
	database.InitDatabase(databaseURL)
}
