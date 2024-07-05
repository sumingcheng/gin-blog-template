FROM node:18.17-alpine AS feBuild

WORKDIR /gin-blog/web

COPY ./web .
RUN npm config set registry https://registry.npmmirror.com/ && \
    apk add --no-cache libc6-compat && \
    npm install -g pnpm && \
    pnpm install && \
    pnpm run build

FROM golang:1.22.2 AS goBuild

WORKDIR /gin-blog

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux

COPY . .
COPY --from=feBuild /gin-blog/web/dist ./gin-blog/web/dist
ENV GOPROXY=https://goproxy.io,direct
RUN go mod download && \
    go build -ldflags "-s -w -extldflags '-static'" -o gin-blog

FROM alpine:3.16
# 替换为阿里云镜像源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

WORKDIR /data

RUN apk add --no-cache ca-certificates tzdata && \
    update-ca-certificates

ENV PORT=5678
COPY --from=goBuild /gin-blog/gin-blog /
EXPOSE 5678

ENTRYPOINT ["/gin-blog"]
