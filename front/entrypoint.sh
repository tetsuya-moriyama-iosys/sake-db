#!/bin/sh

cd /app
npm install  # パッケージをインストール
exec npm run dev  # `exec` で `npm run dev` を実行
