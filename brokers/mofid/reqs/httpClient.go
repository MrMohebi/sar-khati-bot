package mofid

import (
	"bytes"
	"fmt"
	"github.com/MrMohebi/sar-khati-bot/common"
	"github.com/MrMohebi/sar-khati-bot/configs"
	"net/http"
	"net/url"
)

type QueryParams struct {
	Key   string
	Value string
}

func PostHttp(postUrl string, dataJSON []byte) *http.Request {

	req, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(dataJSON))
	common.IsErr(err, false)

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,fa-IR;q=0.8,fa;q=0.7")
	req.Header.Set("Authorization", "BasicAuthentication "+configs.IniGet("broker_mofid", "authToken"))
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

	return req
}

func GetHttp(getUrl string, params []QueryParams) *http.Request {

	if len(params) > 0 {
		for _, param := range params {
			if len(params) == 1 {
				getUrl += "?" + param.Key + "=" + param.Value
			} else if len(params) > 1 {
				getUrl += "&" + param.Key + "=" + param.Value
			}
		}
		u, _ := url.ParseRequestURI(getUrl)
		u.RawQuery = u.Query().Encode()
		urlStr := fmt.Sprintf("%v", u)
		getUrl = urlStr
		println(getUrl)
	}

	req, err := http.NewRequest("GET", getUrl, nil)
	common.IsErr(err, false)

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,fa-IR;q=0.8,fa;q=0.7")
	req.Header.Set("Authorization", "BasicAuthentication "+configs.IniGet("broker_mofid", "authToken"))
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

	return req

}
