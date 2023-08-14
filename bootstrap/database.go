package bootstrap

import (
	"github.com/nimblehq/mark-ic/database"
)

func InitDatabase(databaseURL string) {
	database.InitDatabase(databaseURL)
}
