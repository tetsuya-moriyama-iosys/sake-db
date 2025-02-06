#!/bin/sh

echo "Waiting for MongoDB to start at ${MONGODB_HOST}:${MONGODB_PORT}..."

# nc コマンドで MongoDB が起動するまで待機
while ! nc -z ${MONGODB_HOST} ${MONGODB_PORT}; do
  sleep 1
done

echo "MongoDB is up - starting backend"

# コンテナのメインプロセスを実行（Dockerfile の `CMD`）
exec "$@"