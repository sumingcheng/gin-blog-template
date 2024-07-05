FROM node:18.17-alpine AS webBuild

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
COPY --from=webBuild /gin-blog/web/dist ./gin-blog/web/dist
ENV GOPROXY=https://goproxy.io,direct
RUN go mod download
RUN go build -ldflags "-s -w -extldflags '-static'" -o gin-blog

FROM alpine:3.16

WORKDIR /data

RUN apk update \
    && apk upgrade \
    && apk add --no-cache ca-certificates tzdata \
    && update-ca-certificates 2>/dev/null || true
ENV PORT=5678
COPY --from=goBuild /gin-blog/gin-blog /
EXPOSE 5678

ENTRYPOINT ["/gin-blog"]
