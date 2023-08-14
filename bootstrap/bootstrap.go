package bootstrap

import (
	"github.com/nimblehq/mark-ic/database"
)

func Init() {
	LoadConfig()

	InitDatabase(database.GetDatabaseURL())
}
