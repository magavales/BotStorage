package session

import (
	"github.com/google/uuid"
	"net/http"
	"time"
)

type Cookie struct {
	Cookie *http.Cookie
}

func (c *Cookie) SetCookie(now time.Time) {
	c.Cookie = &http.Cookie{
		Name:     "Session",
		Value:    uuid.New().String(),
		Path:     "/",
		Domain:   "localhost",
		Expires:  now.Add(time.Hour * 4),
		MaxAge:   0,
		Secure:   true,
		HttpOnly: true,
	}
}

func (c *Cookie) Compare(cookie http.Cookie) bool {
	if c.Cookie.Name == cookie.Name && c.Cookie.Value == cookie.Value && c.Cookie.Expires == cookie.Expires && c.Cookie.Path == cookie.Path {
		return true
	} else {
		return false
	}
}
