<h1 align="center">Gin Blog Template</h1>

<p align="center">
基于 Go + React 的生产级博客模板。<br>
三层架构，双 Token 鉴权，一键部署。
</p>

<p align="center">
  <a href="https://raw.githubusercontent.com/sumingcheng/gin-blog/main/LICENSE"><img src="https://img.shields.io/github/license/sumingcheng/gin-blog?color=353535" alt="license"></a>
  <a href="https://hub.docker.com/repository/docker/smcroot/gin-blog"><img src="https://img.shields.io/docker/pulls/smcroot/gin-blog?color=353535" alt="docker pull"></a>
  <a href="https://goreportcard.com/report/github.com/sumingcheng/gin-blog"><img src="https://goreportcard.com/badge/github.com/sumingcheng/gin-blog" alt="Go Report Card"></a>
</p>

<p align="center">
  <a href="./README.md">English</a>
</p>

---

## 技术栈

| 层级 | 技术 |
|---|---|
| 后端 | Go 1.25, Gin, GORM, PostgreSQL 16, Redis |
| 前端 | React 18, Vite, Chakra UI |
| 鉴权 | JWT (golang-jwt/v5) + Refresh Token, bcrypt |
| 可观测性 | Prometheus, Grafana, Logrus |
| 部署 | Docker, Docker Compose, 多阶段构建 |

## 项目结构

```
.
├── config/          # YAML 配置（数据库、Redis、JWT 密钥）
├── database/        # GORM 模型、数据库/Redis 连接、初始化 SQL
├── handler/         # HTTP 处理器（薄层，委托给 service）
├── middleware/       # 鉴权、CORS、限流、请求日志、指标采集
├── model/           # 统一响应结构、错误码、分页
├── router/          # 路由定义
├── service/         # 业务逻辑
├── util/            # JWT、bcrypt、配置加载、日志、校验翻译
├── web/             # React 前端（Vite）
├── deploy/          # Prometheus 与 Grafana 配置
├── Dockerfile       # 多阶段构建
├── docker-compose.yaml
└── main.go
```

## 接口

### 公开接口

| 方法 | 路径 | 说明 |
|---|---|---|
| POST | `/api/register` | 注册 |
| POST | `/api/login` | 登录（返回 JWT，设置 refresh token cookie） |
| POST | `/api/logout` | 登出 |
| GET | `/api/token` | 刷新 auth token |
| GET | `/api/blog/list` | 博客列表（分页、搜索） |
| GET | `/api/blog/:bid` | 博客详情 |

### 需要登录

| 方法 | 路径 | 说明 |
|---|---|---|
| GET | `/api/user/profile` | 获取个人信息 |
| POST | `/api/user/password` | 修改密码 |
| POST | `/api/blog/create` | 创建博客 |
| POST | `/api/blog/update` | 更新博客 |
| DELETE | `/api/blog/:bid` | 删除博客（软删除） |

### 其他

| 路径 | 说明 |
|---|---|
| `/health` | 健康检查（DB + Redis） |
| `/metrics` | Prometheus 指标 |
| `/swagger/*` | Swagger 文档 |

## 快速开始

### 前置条件

- Docker & Docker Compose

### 启动

```bash
git clone https://github.com/sumingcheng/gin-blog-template.git
cd gin-blog-template

# 构建镜像
make build

# 启动所有服务
docker compose up -d
```

访问 `http://localhost:5678`。

默认账号：`admin` / `123456`

### 本地开发

```bash
# 只启动 PostgreSQL + Redis
docker compose -f docker-compose-dev.yaml up -d postgres redis

# 修改 config/postgres.yaml 和 config/redis.yaml 的 host 为 localhost

# 启动后端
go run main.go

# 另开终端，启动前端
cd web && npm install && VITE_APP_ENV=development npm run dev
```

前端开发服务器运行在 `http://localhost:5173`，API 请求代理到 `:5678`。

## 鉴权流程

```
客户端                    服务端                     Redis
  |                         |                         |
  |-- POST /login --------->|                         |
  |                         |-- 校验密码（bcrypt）      |
  |                         |-- 生成 JWT               |
  |                         |-- 生成 refresh token     |
  |                         |-- SET refresh:auth ---->|
  |<-- JWT + cookie --------|                         |
  |                         |                         |
  |-- GET /api/* ---------->|                         |
  |   (auth_token header)   |-- 校验 JWT              |
  |<-- 响应 ----------------|                         |
  |                         |                         |
  |-- GET /token ---------->|                         |
  |   (refresh cookie)      |-- GET refresh token --->|
  |<-- 新 auth_token -------|<-- auth token ----------|
```

## 监控

启动后，将 `deploy/grafana/gin-blog.json` 导入 Grafana 仪表盘。

- Prometheus：`http://localhost:59090`
- Grafana：`http://localhost:53000`（admin / admin123456）

## 许可证

[MIT](./LICENSE)
