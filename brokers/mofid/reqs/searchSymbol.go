package mofid

import (
	"github.com/MrMohebi/sar-khati-bot/brokers/mofid"
	mofidTypes "github.com/MrMohebi/sar-khati-bot/brokers/mofid/types"
	"github.com/MrMohebi/sar-khati-bot/common"
	"net/http"
)

func searchSymbol(authToken string) (mofidTypes.SendOrderRes, error) {
	client := &http.Client{}

	req := GetHttp(mofid.SearchSymbolURL, "طلا")

	resp, err := client.Do(req)
	common.IsErr(err, false)

	println(resp)

	return resp.Status, err
}
