package bootstrap

import (
	"os"

	"github.com/markgravity/golang-ic/helpers/log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("app")
	viper.SetConfigType("toml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}
}

func LoadENV() {
	env := gin.Mode()

	// Skip release mode because there is NO ENV file
	if env == gin.ReleaseMode {
		return
	}

	if env == gin.TestMode {
		// Skip loading ENV file if CI is running
		if os.Getenv("CI") == "true" {
			return
		}

		err := godotenv.Load(".env." + env)
		if err != nil {
			log.Fatal("Failed to load .env.", env, err)
		}
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env", err)
	}
}
