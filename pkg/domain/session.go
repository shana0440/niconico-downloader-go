package domain

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type Session struct {
	URL     *url.URL
	Cookies []*http.Cookie
}

func MakeSession(resp *http.Response) Session {
	if resp == nil {
		return Session{}
	}

	return Session{
		URL:     resp.Request.URL,
		Cookies: resp.Cookies(),
	}
}

func (s *Session) IsLoginFailure() bool {
	for _, cookie := range s.Cookies {
		if cookie.Name == "user_session" {
			return false
		}
	}
	return true
}

func (s *Session) WithClient(client *http.Client) *http.Client {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalln(err)
	}
	jar.SetCookies(s.URL, s.Cookies)
	client.Jar = jar
	return client
}
