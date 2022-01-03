package helpers

import (
	"net/http"
	"time"
)

func CreateCookie(token string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = "jon"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	return cookie
}
