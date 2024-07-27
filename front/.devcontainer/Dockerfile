# front/.devcontainer/Dockerfile
FROM node:latest AS builder

WORKDIR /app

COPY package*.json ./
RUN npm cache clean --force
RUN npm install
RUN npm install @rollup/rollup-linux-x64-gnu --save-optional

COPY . .

EXPOSE 3000

RUN npm run build
