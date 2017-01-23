package auth

import (
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
)

// NewCookieStore returns a new instance of CookieStore
func NewCookieStore(cookieName string, hashKey, blockKey []byte) *CookieStore {
	return &CookieStore{
		sc:   securecookie.New(hashKey, blockKey),
		name: cookieName,
	}
}

// CookieStore wraps a secure cookie and provides admin specific functionality
type CookieStore struct {
	sc   *securecookie.SecureCookie
	name string
}

// SetSession encodes and sets the session on the response
func (c *CookieStore) SetSession(w http.ResponseWriter, session interface{}) error {
	if session == nil {
		http.SetCookie(w, &http.Cookie{
			Name:    c.name,
			Path:    "/",
			Expires: time.Now(),
			MaxAge:  -1,
		})

		return nil
	}

	enc, err := c.sc.Encode(c.name, session)

	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:  c.name,
		Value: enc,
		Path:  "/",
	})

	return nil
}

// GetSession reads the session cookie, decodes and returns the value
func (c *CookieStore) GetSession(r *http.Request, dest interface{}) error {
	ck, err := r.Cookie(c.name)

	if err != nil {
		return err
	}

	return c.sc.Decode(c.name, ck.Value, dest)
}
