package x

import (
	"backend/middlewares/customError"
	"crypto/sha256"
	"encoding/base64"
	"golang.org/x/oauth2"
	"net/url"
	"os"
)

// TwitterのOAuth2.0エンドポイントを手動で定義
var twitterEndpoint = oauth2.Endpoint{
	AuthURL:  "https://twitter.com/i/oauth2/authorize",
	TokenURL: "https://api.twitter.com/2/oauth2/token",
}

type TwitterUser struct {
	ID    string `bson:"id"`
	Name  string `bson:"name"`
	Image string `bson:"profile_image_url"`
}

// TwitterOAuthConfig oauth2.Configをtwitter用にラップしたもの
type TwitterOAuthConfig struct {
	oauth2.Config
}

func NewOAuthConfig() *TwitterOAuthConfig {
	return &TwitterOAuthConfig{
		Config: oauth2.Config{
			ClientID:     os.Getenv("TWITTER_OAUTH_CLIENT"),
			ClientSecret: os.Getenv("TWITTER_OAUTH_SECRET"),
			RedirectURL:  os.Getenv("BACK_URI") + "/x/callback",
			Scopes:       []string{"users.read", "tweet.read"}, // tweet.readは必須らしい(Forbiddenになってしまう)
			Endpoint:     twitterEndpoint,
		},
	}
}

func generateCodeChallenge(verifier string) string {
	// code_challengeを生成
	hash := sha256.Sum256([]byte(verifier))
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(hash[:])
}

func generateCodeVerifier() string {
	// ランダムなcode_verifierを生成
	//codeVerifier := helper.RandomStr(43) // 推奨長さ: 43〜128文字
	return "test_code_verifier"
}

// GenerateAuthCodeURL twitter用に最適化されたAuthCodeURLを生成
func (c *TwitterOAuthConfig) GenerateAuthCodeURL(state string) (*string, *customError.Error) {
	// URLをパース
	parsedURL, err := url.Parse(c.AuthCodeURL(state))
	if err != nil {
		// エラー処理（パース失敗時）
		return nil, errParseURL(err)
	}

	// 既存のクエリパラメータを取得
	query := parsedURL.Query()

	// 新しいクエリパラメータを追加
	code := generateCodeVerifier()
	challengeCode := generateCodeChallenge(code)
	query.Set("code_challenge", challengeCode)
	query.Set("code_challenge_method", "S256")

	// クエリパラメータを更新
	parsedURL.RawQuery = query.Encode()

	// 完成したURLを返す
	generatedUrl := parsedURL.String()
	return &generatedUrl, nil
}
