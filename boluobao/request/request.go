package request

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type HttpUtils struct {
	url         string
	method      string
	Cookie      string
	cookie      []*http.Cookie
	response    *http.Request
	query_data  *url.Values
	result_body []byte
}

func NewHttpUtils(host string, path string, method string) *HttpUtils {
	return &HttpUtils{method: method, query_data: &url.Values{}, url: host + path}
}

func (is *HttpUtils) NewRequests() *HttpUtils {
	var err error
	if is.method == "GET" {
		is.response, err = http.NewRequest(is.method, is.url+"?"+is.query_data.Encode(), nil)
	} else {
		is.response, err = http.NewRequest(is.method, is.url, is.GetEncodeParams())
	}
	if err != nil {
		panic(err)
	}
	is.result_body = nil
	is.response.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	is.response.Header.Set("sf-minip-info", "minip_novel/1.0.70(android;11)/wxmp")
	is.response.Header.Set("Cookie", is.Cookie)

	if response, ok := http.DefaultClient.Do(is.response); ok == nil {
		is.cookie = response.Cookies()
		is.result_body, _ = io.ReadAll(response.Body)
	} else {
		fmt.Println("Error: ", ok)
	}
	return is
}

func (is *HttpUtils) Unmarshal(s any) *HttpUtils {
	err := json.Unmarshal(is.result_body, s)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return is
}
