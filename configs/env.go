package configs

import (
	"github.com/MrMohebi/sar-khati-bot/common"
	"github.com/joho/godotenv"
	"os"
)

var IsInitOnce = false

func EnvSetup() {
	if !IsInitOnce {
		err := godotenv.Load()
		common.IsErr(err, false, "Error loading .env file")
		IsInitOnce = true
	} else {
		println("initialized envs once")
	}
}

func EnvMongoURI() string {
	return os.Getenv("MONGO_FINAL_URI")
}

func EvnMongoDB() string {
	return os.Getenv("MONGO_DB")
}
