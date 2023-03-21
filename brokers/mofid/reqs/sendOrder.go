package mofid

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/MrMohebi/sar-khati-bot/brokers/mofid"
	mofidTypes "github.com/MrMohebi/sar-khati-bot/brokers/mofid/types"
	"github.com/MrMohebi/sar-khati-bot/common"
	"io"
	"net/http"
)

func SendOrder(authToken string) (mofidTypes.SendOrderRes, error) {
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

	req, err := http.NewRequest("POST", mofid.SendOrderURL, bytes.NewBuffer(dataJSON))
	common.IsErr(err, false)

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,fa-IR;q=0.8,fa;q=0.7")
	req.Header.Set("Authorization", "BasicAuthentication "+authToken)
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", "https://mofidonline.com")
	req.Header.Set("Referer", "https://mofidonline.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="111", "Not(A:Brand";v="8", "Chromium";v="111"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
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
