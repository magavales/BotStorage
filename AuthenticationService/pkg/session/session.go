package session

import (
	"net/http"
	"time"
)

type Cookie struct {
	Cookie *http.Cookie
}

func (c *Cookie) SetCookie() {
	c.Cookie = http.Cookie{
		Name:       "",
		Value:      "",
		Path:       "",
		Domain:     "",
		Expires:    time.Time{},
		RawExpires: "",
		MaxAge:     0,
		Secure:     false,
		HttpOnly:   false,
		SameSite:   0,
		Raw:        "",
		Unparsed:   nil,
	}
}
