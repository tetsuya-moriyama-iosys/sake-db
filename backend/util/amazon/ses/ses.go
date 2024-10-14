package ses

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"os"
)

const (
	pwResetTitle = "パスワードリセット"
)

type emailContent struct {
	Subject string
	To      string
	//Bcc *[]string
	Text string
}

func SendPasswordReset(ctx context.Context, email string, token string) error {
	//メールテンプレートを作る
	msg, err := pwRstTemp(&passwordReset{
		Token: token,
	})
	//メールを送信する
	err = sendMail(ctx, &emailContent{
		Subject: pwResetTitle,
		To:      email,
		Text:    msg,
	})
	return err
}

func sendMail(ctx context.Context, content *emailContent) error {
	// 1. AWSの設定を読み込む
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(os.Getenv("AWS_REGION")), config.WithCredentialsProvider(
		credentials.NewStaticCredentialsProvider(os.Getenv("AWS_SES_ACCESS_KEY"), os.Getenv("AWS_SES_ACCESS_SECRET"), ""),
	))
	if err != nil {
		return err
	}

	// 2. SESクライアントを作成
	client := sesv2.NewFromConfig(cfg)
	from := os.Getenv("AWS_SES_FROM")

	input := &sesv2.SendEmailInput{
		Destination: &types.Destination{
			ToAddresses: []string{content.To},
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Body: &types.Body{
					Text: &types.Content{
						Data: &content.Text,
					},
				},
				Subject: &types.Content{
					Data: &content.Subject,
				},
			},
		},
		FromEmailAddress: &from,
	}

	_, err = client.SendEmail(ctx, input)

	return err
}
