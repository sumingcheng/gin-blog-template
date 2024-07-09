FROM node:18.17-alpine AS build1
WORKDIR /web
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
COPY ./web .
RUN apk add --no-cache libc6-compat && \
    npm config set registry https://mirrors.huaweicloud.com/repository/npm/ && \
    npm install -g pnpm && \
    pnpm install && \
    pnpm run build

FROM golang:1.22.2 AS build2
WORKDIR /gin-blog
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

COPY . .
COPY --from=build1 /web/dist /gin-blog/web/dist
ENV GOPROXY=https://goproxy.io,direct

RUN go mod download
RUN go build -tags netgo -ldflags '-w -s -extldflags "-static"' -o gin-blog

FROM alpine AS build3
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache ca-certificates tzdata && \
    update-ca-certificates
COPY --from=build2 /gin-blog/gin-blog /app/gin-blog 
COPY /config /config
RUN chmod +x /app/gin-blog
ENV APP_ENV=production

EXPOSE 5678
ENTRYPOINT ["/app/gin-blog"]

