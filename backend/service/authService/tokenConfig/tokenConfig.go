package tokenConfig

import (
	"os"
	"time"
)

type TokenConfig struct {
	AccessSecretKey  []byte
	RefreshSecretKey []byte
	AccessExpire     time.Duration
	RefreshExpire    time.Duration
	FrontDomain      string
}

func NewTokenConfig() *TokenConfig {
	return &TokenConfig{
		AccessSecretKey:  []byte(os.Getenv("JWT_SECRET_KEY")),
		RefreshSecretKey: []byte(os.Getenv("JWT_REFRESH_KEY")),
		AccessExpire:     15 * time.Minute,
		RefreshExpire:    7 * 24 * time.Hour,
		FrontDomain:      os.Getenv("FRONT_URI"),
	}
}
