FROM node:18.17-alpine AS build1
WORKDIR /web
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
ARG NPM_REGISTRY=https://registry.npmjs.org/

COPY ./web .
RUN apk add --no-cache libc6-compat && \
    npm config set registry ${NPM_REGISTRY} && \
    npm install -g pnpm && \
    pnpm install --force && \
    pnpm run build

FROM golang:1.22.2-alpine AS build2
WORKDIR /gin-blog
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

COPY . .
COPY --from=build1 /web/dist /gin-blog/web/dist
ENV GOPROXY=https://goproxy.io,direct

RUN go mod download
RUN go build -tags netgo -ldflags '-w -s -extldflags "-static"' -o gin-blog

FROM alpine:latest
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache ca-certificates tzdata && \
    update-ca-certificates
COPY --from=build2 /gin-blog/gin-blog /app/gin-blog 
COPY /config /config
RUN chmod +x /app/gin-blog
ENV APP_ENV=production

EXPOSE 5678
ENTRYPOINT ["/app/gin-blog"]

