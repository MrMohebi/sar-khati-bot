package mofid

import (
	"encoding/json"
	"errors"
	"github.com/MrMohebi/sar-khati-bot/brokers/mofid"
	mofidTypes "github.com/MrMohebi/sar-khati-bot/brokers/mofid/types"
	"github.com/MrMohebi/sar-khati-bot/common"
	"io"
	"net/http"
)

func SearchSymbol() (mofidTypes.SearchSymbolRes, error) {
	client := &http.Client{}

	// To use this function change the key based on the query param you want to use
	//and the value based on the value you want to search for
	var params = []QueryParams{
		{Key: "term", Value: "طلا"},
	}

	req := GetHttp(mofid.SearchSymbolURL, params)

	resp, err := client.Do(req)
	common.IsErr(err, false)

	var res mofidTypes.SearchSymbolRes
	var errorResponse error

	bodyText, err := io.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusOK {
		err = json.Unmarshal(bodyText, &res)
		common.IsErr(err, false)
	} else if resp.StatusCode == http.StatusUnauthorized {
		errorResponse = errors.New("unauthorized")
	}

	println(string(bodyText))

	return res, errorResponse
}
