package app_test

import (
	"context"
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/T-V-N/gourlshortener/internal/app"
	"github.com/T-V-N/gourlshortener/internal/config"
	"github.com/T-V-N/gourlshortener/internal/storage"
)

// Generate a random url function
func GenURL() string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	ret := make([]byte, 10)

	for i := 0; i < 10; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return ""
		}
		ret[i] = letters[num.Int64()]
	}

	return "https://www." + string(ret) + ".com"
}

func BenchmarkSaveUrl(b *testing.B) {
	cfg := &config.Config{}
	st := storage.InitStorage(map[string]storage.URL{}, cfg)
	a := app.NewApp(context.Background(), st, cfg)
	a.Init()

	for i := 0; i < b.N; i++ {
		a.SaveURL(context.Background(), GenURL(), "test")
	}
}
