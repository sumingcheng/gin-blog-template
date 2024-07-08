FROM node:18.17-alpine AS build1
WORKDIR /web
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
COPY ./web .
RUN npm config set registry https://registry.npmmirror.com/ && \
    apk add --no-cache libc6-compat && \
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
COPY --from=build2 /gin-blog/gin-blog /gin-blog
COPY /config /config
RUN chmod +x /gin-blog
ENV CONFIG_PATH=/config

EXPOSE 5678
ENTRYPOINT ["/gin-blog"]
