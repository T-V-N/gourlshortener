package auth

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"log"
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
		Name:   "auth_token",
		Value:  hex.EncodeToString(cookieVal),
		MaxAge: 300,
	}, nil

}

func InitAuth(cfg *config.Config) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("auth_token")
			if err != nil {
				log.Println("no cookie set, creating a new one")
			}

			if cookie != nil {
				valid, err := isValidCookie(cookie, cfg.SecretKey)
				if err != nil {
					http.Error(w, "Can't validate the token", http.StatusBadRequest)
				}

				if valid {
					hexValue, _ := hex.DecodeString(cookie.Value)
					r.Header.Set("uid", hex.EncodeToString(hexValue[32:]))
				} else {
					newCookie, err := generateCookie(cfg.SecretKey)
					if err != nil {
						http.Error(w, "Something went wrong", http.StatusBadRequest)
					}
					http.SetCookie(w, newCookie)
				}
			} else {
				newCookie, err := generateCookie(cfg.SecretKey)
				if err != nil {
					http.Error(w, "Something went wrong", http.StatusBadRequest)
				}
				http.SetCookie(w, newCookie)
			}

			next.ServeHTTP(w, r)
		})
	}
}
