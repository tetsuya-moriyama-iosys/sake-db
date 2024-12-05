package x

import (
	"golang.org/x/oauth2"
	"os"
)

// Twitter の OAuth2 エンドポイントを手動で定義
var twitterEndpoint = oauth2.Endpoint{
	AuthURL:  "https://twitter.com/i/oauth2/authorize",
	TokenURL: "https://api.twitter.com/2/oauth2/token",
}

var oauthConfig = &oauth2.Config{
	ClientID:     os.Getenv("TWITTER_OAUTH_CLIENT"),
	ClientSecret: os.Getenv("TWITTER_OAUTH_SECRET"),
	RedirectURL:  os.Getenv("BACK_URI") + "/auth/callback",
	Endpoint:     twitterEndpoint, // 上で定義したエンドポイントを使用
	Scopes:       []string{"tweet.read", "users.read"},
}
