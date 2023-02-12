// Package auth allows server to handle requests using auth middleware which automatically checks users by request cookie and handles all related stuff ;)
package auth

import (
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/T-V-N/gourlshortener/internal/config"
)

// UIDKey ensures a user UID will be safe in the user context and won't be re-written by other layers
type UIDKey struct{}

// GenerateRandom generates random byte arr of given size
func GenerateRandom(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func isValidCookie(c *http.Cookie, key string) (bool, error) {
	value, err := hex.DecodeString(c.Value)
	if err != nil {
		return false, err
	}
	signature := value[:32]

	h := hmac.New(sha256.New, []byte(key))
	h.Write(value[32:])
	dst := h.Sum(nil)

	return hmac.Equal(dst, signature), nil
}

func generateCookie(key string) (c *http.Cookie, err error) {
	uid, err := GenerateRandom(4)
	if err != nil {
		return nil, err
	}

	h := hmac.New(sha256.New, []byte(key))
	h.Write(uid)
	signature := h.Sum(nil)

	cookieVal := append(signature, uid...)

	return &http.Cookie{
		Name:  "auth_token",
		Value: hex.EncodeToString(cookieVal),
		Path:  "/",
	}, nil

}

// InitAuth creates a MW that parses an incoming request's cookie and tries to extract UID stored in the extracted data.
// In case there is no cookie available or it is available but invalid, the auth mw generates a new UID and cookie and sets it to the
// Context.
func InitAuth(cfg *config.Config) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, cookieErr := r.Cookie("auth_token")
			if cookieErr != nil {
				newCookie, err := generateCookie(cfg.SecretKey)
				if err != nil {
					http.Error(w, "Something went wrong", http.StatusBadRequest)
					return
				}

				uid := newCookie.Value[32:]
				ctx := context.WithValue(r.Context(), UIDKey{}, uid)

				http.SetCookie(w, newCookie)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			isValid, err := isValidCookie(cookie, cfg.SecretKey)
			if err != nil {
				http.Error(w, "Can't validate the token", http.StatusInternalServerError)
				return
			}

			var ctx context.Context
			if isValid {
				hexValue := cookie.Value
				ctx = context.WithValue(r.Context(), UIDKey{}, hexValue[32:])
			} else {
				newCookie, err := generateCookie(cfg.SecretKey)
				if err != nil {
					http.Error(w, "Something went wrong", http.StatusBadRequest)
					return
				}

				uid := newCookie.Value[32:]
				ctx = context.WithValue(r.Context(), UIDKey{}, uid)

				http.SetCookie(w, newCookie)
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
