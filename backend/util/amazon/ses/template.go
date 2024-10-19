package ses

import (
	"bytes"
	"os"
	"text/template"
)

type passwordReset struct {
	Token string
}

// generatePwRstStr はパスワードリセットメール本文のテンプレートです
func generatePwRstStr() string {
	return "URLは " + os.Getenv("FRONT_URI") + "/auth/password/reset/{{ .Token }} です"
}

func pwRstTemp(cfg *passwordReset) (string, error) {
	tmpl, err := template.New("psw-rst-template").Parse(generatePwRstStr())
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
