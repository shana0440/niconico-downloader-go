package auth

import (
	"log"
	"net/http"
	"net/url"

	"github.com/shana0440/niconico-downloader-go/pkg/domain"
)

const (
	LOGIN_URL = "https://account.nicovideo.jp/api/v1/login?site=niconico"
)

func Login(account, password string) domain.Session {
	if account == "" && password == "" {
		log.Println("Account and password is empty, execute as guest mode")
		return domain.Session{}
	}

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.PostForm(LOGIN_URL, url.Values{
		"mail_tel": {account},
		"password": {password},
	})

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	session := domain.MakeSession(resp)
	if session.IsLoginFailure() {
		log.Fatalln("Failed to login, please check the account and password is correct.")
	}

	return session
}
