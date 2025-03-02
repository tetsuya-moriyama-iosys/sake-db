#!/bin/sh

cd /app
npm install  # パッケージをインストール
exec npm run dev -- --host
