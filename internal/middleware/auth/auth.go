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

type UIDKey struct{}

func generateRandom(size int) ([]byte, error) {
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
	uid, err := generateRandom(4)
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

func InitAuth(cfg *config.Config) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("auth_token")
			if err != nil {
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

			valid, err := isValidCookie(cookie, cfg.SecretKey)
			if err != nil {
				http.Error(w, "Can't validate the token", http.StatusBadRequest)
				return
			}

			var ctx context.Context
			if valid {
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
