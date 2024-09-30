package ses

import (
	"bytes"
	"text/template"
)

type passwordReset struct {
	Token string
}

const (
	pwRstRawText = "トークンは {{ .Token }} です"
)

func pwRstTemp(cfg *passwordReset) (string, error) {
	tmpl, err := template.New("psw-rst-template").Parse(pwRstRawText)
	if err != nil {
		return "", err
	}

	// バッファにテンプレートの結果を出力
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, *cfg)
	if err != nil {
		return "", err
	}

	// バッファの内容を文字列として返す
	return buf.String(), nil
}
