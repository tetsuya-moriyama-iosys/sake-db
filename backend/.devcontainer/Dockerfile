FROM golang:1.23

# 必要なツールをインストール（netcat を含む）
RUN apt-get update \
    && apt-get install -y netcat-openbsd \
    && go install github.com/go-delve/delve/cmd/dlv@latest \
    && go install github.com/air-verse/air@latest

# 作業ディレクトリを設定
WORKDIR /app

# Go Modules の依存関係をダウンロード
COPY go.mod ./
COPY go.sum ./
RUN go mod tidy
RUN go mod download

# ソースコードをコンテナにコピー
COPY . .

# MongoDB の起動を待つエントリポイントスクリプトをコピー
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# アプリケーションをビルド
RUN go build -o /backend

# ポート 8080 を公開
EXPOSE 8080

# エントリポイントスクリプトを実行
ENTRYPOINT ["/entrypoint.sh"]
CMD ["/backend"]