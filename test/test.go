package test

import (
	"os"
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

	bootstrap.InitDatabase(database.GetDatabaseURL())
}

func setRootDir() {
	_, currentFile, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(currentFile), "../")

	err := os.Chdir(root)
	if err != nil {
		log.Fatal("Failed to set root directory: ", err)
	}
}
