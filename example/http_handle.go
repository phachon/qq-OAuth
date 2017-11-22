package main

import (
	"net/http"
	"github.com/phachon/qq-OAuth"
	"log"
)

type HttpHandle struct {

}

var state = qq_OAuth.NewUtils().RandString(8)

// login action
func Login(res http.ResponseWriter, req *http.Request) {
	appId := ""
	appSecret := ""
	callback := ""
	scope := ""

	oAuth := qq_OAuth.NewOAuth(appId, appSecret, callback, scope)
	loginUrl := oAuth.GetAuthorURL(state)

	http.Redirect(res, req, loginUrl, 302)
}

// callback action
func Callback(res http.ResponseWriter, req *http.Request)  {
	reqState := req.Form.Get("state")
	if reqState != state {
		res.Write([]byte("error"))
	}

	authCode := req.Form.Get("code")

	appId := ""
	appSecret := ""
	callback := ""
	scope := ""

	oAuth := qq_OAuth.NewOAuth(appId, appSecret, callback, scope)
	oAuth.Access(authCode)

	oAuth.GetUserInfo()
}

func main()  {

	http.HandleFunc("/qq/login", Login)
	http.HandleFunc("/qq/callback", Callback)

	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Println(err.Error())
	}
}