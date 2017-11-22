package qq_OAuth

import (
	"errors"
	"strings"
	"encoding/json"
)

const (
	QAuth_Api_AuthCode = "https://graph.qq.com/oauth2.0/authorize"
	QAuth_Api_AccessToken = "https://graph.qq.com/oauth2.0/token"
	QAuth_Api_OpenId = "https://graph.qq.com/oauth2.0/me"
)

type OAuth struct {
	AppId string
	AppSecret string
	Callback string
	Scope string
}

func NewOAuth(appId string, appSecret string, callback string, scope string) *OAuth {
	return &OAuth{
		AppId: appId,
		AppSecret: appSecret,
		Callback: callback,
		Scope: scope,
	}
}

// login action auth code
// params: state rand string
// return login url
func (oAuth *OAuth) AuthCode(state string) string {
	value := map[string]string{
		"response_type": "code",
		"client_id": oAuth.AppId,
		"redirect_uri": oAuth.Callback,
		"state": state,
		"scope": "",
	}

	loginUrl := NewUtils().QueryBuilder(QAuth_Api_AuthCode, value)
	return loginUrl
}

// get access token
func (oAuth *OAuth) GetAccessToken(authCode string) (string, error) {

	value := map[string]string {
		"grant_type": "authorization_code",
		"client_id": oAuth.AppId,
		"redirect_uri": oAuth.Callback,
		"client_secret": oAuth.AppSecret,
		"code": authCode,
	}
	response, httpCode, err := NewUtils().HttpGet(QAuth_Api_AccessToken, value, nil)
	if err != nil {
		return "", err
	}
	if httpCode != 200 {
		return "", errors.New("http code error")
	}

	params := NewUtils().ParseString(response)
	accessToken, ok := params["access_token"]
	if !ok {
		if msg, ok := params["msg"]; ok {
			return "", errors.New(msg)
		}else {
			return "", nil
		}
	}
	return accessToken, nil
}

// get open id
func (oAuth *OAuth) GetOpenId(accessToken string) (string, error) {
	value := map[string]string {
		"access_token": accessToken,
	}
	response, httpCode, err := NewUtils().HttpGet(QAuth_Api_OpenId, value, nil)
	if err != nil {
		return "", err
	}
	if httpCode != 200 {
		return "", errors.New("http code error")
	}

	if strings.Contains(response, "callback") {
		response = strings.TrimLeft(response, "callback(")
		response = strings.TrimRight(response, ")")
	}

	var resData map[string]string
	err = json.Unmarshal([]byte(response), &resData)
	if err != nil {
		return "", nil
	}
	if openid, ok := resData["openid"]; ok {
		return openid, nil
	}else {
		return "", nil
	}
}