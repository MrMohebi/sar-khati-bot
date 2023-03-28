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

func SendOrder() (mofidTypes.SendOrderRes, error) {
	data := mofidTypes.OrderReq{
		IsSymbolCautionAgreement:  false,
		CautionAgreementSelected:  false,
		IsSymbolSepahAgreement:    false,
		SepahAgreementSelected:    false,
		OrderCount:                8585,
		OrderPrice:                198798,
		FinancialProviderID:       1,
		MinimumQuantity:           0,
		MaxShow:                   0,
		OrderID:                   0,
		Isin:                      "IRTKLOTF0001",
		OrderSide:                 65,
		OrderValidity:             74,
		OrderValiditydate:         nil,
		ShortSellIsEnabled:        false,
		ShortSellIncentivePercent: 0,
	}

	dataJSON, err := json.Marshal(data)
	common.IsErr(err, false)

	client := &http.Client{}

	req := PostHttp(mofid.SendOrderURL, dataJSON)

	resp, err := client.Do(req)
	common.IsErr(err, false)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		common.IsErr(err, false)
	}(resp.Body)

	bodyText, err := io.ReadAll(resp.Body)
	common.IsErr(err, false)

	var res mofidTypes.SendOrderRes
	var errorResponse error

	if resp.StatusCode == http.StatusOK {
		err = json.Unmarshal(bodyText, &res)
		common.IsErr(err, false)
	} else if resp.StatusCode == http.StatusUnauthorized {
		errorResponse = errors.New("unauthorized")
	}

	return res, errorResponse
}
