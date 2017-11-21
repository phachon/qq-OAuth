package qq_OAuth

import (
	"strings"
	"net/url"
	"net/http"
	"io/ioutil"
	"math/rand"
	"time"
)

type Utils struct {

}

func NewUtils() *Utils {
	return &Utils{}
}

// http get request
func (utils *Utils) HttpGet(queryUrl string, queryValues map[string]string, headerValues map[string]string) (body string, code int, err error) {
	if !strings.Contains(queryUrl, "?") {
		queryUrl += "?"
	}

	queryString := ""
	for queryKey, queryValue := range queryValues {
		queryString = queryString + "&" + queryKey + "=" + url.QueryEscape(queryValue)
	}
	queryString = strings.Replace(queryString, "&", "", 1)
	queryUrl += queryString

	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		return
	}
	if (headerValues != nil) && (len(headerValues) > 0) {
		for key, value := range headerValues {
			req.Header.Set(key, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	code = resp.StatusCode
	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return string(bodyByte), code, nil
}

// http post request
func (utils *Utils) HttpPost(queryUrl string, queryValues map[string]string, headerValues map[string]string) (body string, code int, err error) {
	if !strings.Contains(queryUrl, "?") {
		queryUrl += "?"
	}
	queryString := ""
	for queryKey, queryValue := range queryValues {
		queryString = queryString + "&" + queryKey + "=" + url.QueryEscape(queryValue)
	}
	queryString = strings.Replace(queryString, "&", "", 1)
	queryUrl += queryString

	req, err := http.NewRequest("POST", queryUrl, strings.NewReader(queryString))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if (headerValues != nil) && (len(headerValues) > 0) {
		for key, value := range headerValues {
			req.Header.Set(key, value)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	code = resp.StatusCode
	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return string(bodyByte), code, nil
}

// get rand string
func (utils *Utils) RandString(strLen int) string {
	codes := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	codeLen := len(codes)
	data := make([]byte, strLen)
	rand.Seed(time.Now().UnixNano() + rand.Int63() + rand.Int63() + rand.Int63() + rand.Int63())
	for i := 0; i < strLen; i++ {
		idx := rand.Intn(codeLen)
		data[i] = byte(codes[idx])
	}
	return string(data)
}

