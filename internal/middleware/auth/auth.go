package auth

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/T-V-N/gourlshortener/internal/config"
)

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
	}, nil

}

func InitAuth(cfg *config.Config) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("auth_token")
			if err != nil {
				newCookie, err := generateCookie(cfg.SecretKey)

				uid := newCookie.Value[32:]
				r.Header.Set("uid", uid)

				if err != nil {
					http.Error(w, "Something went wrong", http.StatusBadRequest)
					return
				}

				http.SetCookie(w, newCookie)
				next.ServeHTTP(w, r)
				return
			}

			valid, err := isValidCookie(cookie, cfg.SecretKey)
			if err != nil {
				http.Error(w, "Can't validate the token", http.StatusBadRequest)
				return
			}

			if valid {
				hexValue := cookie.Value
				r.Header.Set("uid", hexValue[32:])
			} else {
				newCookie, err := generateCookie(cfg.SecretKey)

				uid := newCookie.Value[32:]
				r.Header.Set("uid", uid)
				if err != nil {
					http.Error(w, "Something went wrong", http.StatusBadRequest)
					return
				}
				http.SetCookie(w, newCookie)
			}

			next.ServeHTTP(w, r)
		})
	}
}
