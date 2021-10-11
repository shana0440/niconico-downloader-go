package domain

import "net/http"

type Session struct {
	Cookies map[string]*http.Cookie
}

func MakeSession(resp *http.Response) Session {
	if resp == nil {
		return Session{}
	}

	cookies := make(map[string]*http.Cookie)
	for _, cookie := range resp.Cookies() {
		cookies[cookie.Name] = cookie
	}
	return Session{
		Cookies: cookies,
	}
}

func (s *Session) IsLoginFailure() bool {
	_, ok := s.Cookies["user_session"]
	return !ok
}
