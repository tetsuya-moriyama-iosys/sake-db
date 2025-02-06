# front/.devcontainer/Dockerfile
FROM node:latest AS builder

WORKDIR /app

COPY package*.json ./
RUN npm cache clean --force
RUN npm install
# 不要な場合、この行を削除
# RUN npm install @rollup/rollup-linux-x64-gnu --save-optional

COPY . .

EXPOSE 3000

# submithunnderの型エラーは無視すべきなので
#RUN npm run type-check || true
#RUN npm run build
ENTRYPOINT ["/app/entrypoint.sh"]