package main

import (
	"github.com/MrMohebi/sar-khati-bot/brokers/mofid/reqs"
	"github.com/MrMohebi/sar-khati-bot/common"
	"github.com/MrMohebi/sar-khati-bot/configs"
)

// nodemon --exec go run main.go --signal SIGTERM

func main() {
	configs.Setup()
	println("Mofid test")

	//mofidAuthToken := configs.IniGet("broker_mofid", "authToken")

	res, err := mofid.SearchSymbol()
	common.IsErr(err, false)

	if err == nil {
		println(res[0].Isin)
	}

}
