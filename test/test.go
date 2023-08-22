package test

import (
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/markgravity/golang-ic/bootstrap"
	"github.com/markgravity/golang-ic/database"
	"github.com/markgravity/golang-ic/helpers/log"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func SetupTestEnvironment() {
	gin.SetMode(gin.TestMode)

	setRootDir()

	bootstrap.LoadConfig()
	bootstrap.LoadEnv()

	bootstrap.InitDatabase(database.GetDatabaseURL())

	// Clean all data
	database.GetDB().Exec("TRUNCATE TABLE users")
}

func RootDir() string {
	_, currentFile, _, _ := runtime.Caller(0)
	currentFilePath := path.Join(filepath.Dir(currentFile))
	return filepath.Dir(currentFilePath)
}

func setRootDir() {
	_, currentFile, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(currentFile), "../")

	err := os.Chdir(root)
	if err != nil {
		log.Fatal("Failed to set root directory: ", err)
	}
}
