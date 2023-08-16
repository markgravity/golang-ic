package bootstrap

import (
	"github.com/markgravity/golang-ic/database"
)

func Init() {
	LoadConfig()

	InitDatabase(database.GetDatabaseURL())
}
