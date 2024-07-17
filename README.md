# Gin-Blog 模板

<p align="center">
  <a href="https://raw.githubusercontent.com/sumingcheng/gin-blog/main/LICENSE">
    <img src="https://img.shields.io/github/license/sumingcheng/gin-blog?color=brightgreen" alt="license">
  </a>
  <a href="https://github.com/sumingcheng/gin-blog/releases/latest">
    <img src="https://img.shields.io/github/v/release/sumingcheng/gin-blog?color=brightgreen&include_prereleases" alt="release">
  </a>
  <a href="https://hub.docker.com/repository/docker/justsong/gin-blog">
    <img src="https://img.shields.io/docker/pulls/justsong/gin-blog?color=brightgreen" alt="docker pull">
  </a>
  <a href="https://github.com/sumingcheng/gin-blog/releases/latest">
    <img src="https://img.shields.io/github/downloads/sumingcheng/gin-blog/total?color=brightgreen&include_prereleases" alt="release">
  </a>
  <a href="https://goreportcard.com/report/github.com/sumingcheng/gin-blog">
    <img src="https://goreportcard.com/badge/github.com/sumingcheng/gin-blog" alt="GoReportCard">
  </a>
</p>
## 功能

+ [x] 双`Token`登录，基于令牌的鉴权

+ [x] 使用`logrus`日志文件的自动切割和轮换

+ [x] 使用`viper`多种配置格式和环境变量

+ [x] 使用`translate`翻译错误

+ [x] 自动生成`Swagger`文档

+ [x] 使用`GORM`操作`MySQL`

+ [x] 暴露`Metric`可以直接使用`prometheus`进行监控

+ [x] 前端使用：`vite + react + chakra-ui`

  

## 部署
### 基于 Docker 进行部署
进入项目目录

```
make build
make run
```

数据将会保存在宿主机的 `/home/ubuntu/data/gin-blog` 目录。

### 手动部署
1. 从 [GitHub Releases](https://github.com/sumingcheng/gin-blog/releases/latest) 下载可执行文件或者从源码编译：
   ```shell
   git clone https://github.com/sumingcheng/gin-blog.git
   cd gin-blog/web
   npm install
   npm run build
   cd ..
   go mod download
   go build -ldflags "-s -w" -o gin-blog
   ````
2. 运行：
   ```shell
   chmod u+x gin-blog
   ./gin-blog --port 3000 --log-dir ./logs
   ```
3. 访问 [http://localhost:3000/](http://localhost:3000/) 并登录。初始账号用户名为 `root`，密码为 `123456`。

更加详细的部署教程[参见此处](https://iamazing.cn/page/how-to-deploy-a-website)。

## 配置
系统本身开箱即用。

你可以通过设置环境变量或者命令行参数进行配置。

等到系统启动后，使用 `root` 用户登录系统并做进一步的配置。

### 环境变量
1. `REDIS_CONN_STRING`：设置之后将使用 Redis 作为请求频率限制的存储，而非使用内存存储。
   + 例子：`REDIS_CONN_STRING=redis://default:redispw@localhost:49153`
2. `SESSION_SECRET`：设置之后将使用固定的会话密钥，这样系统重新启动后已登录用户的 cookie 将依旧有效。
   + 例子：`SESSION_SECRET=random_string`
3. `SQL_DSN`：设置之后将使用指定数据库而非 SQLite。
   + 例子：`SQL_DSN=root:123456@tcp(localhost:3306)/gin-blog`

### 命令行参数
1. `--port <port_number>`: 指定服务器监听的端口号，默认为 `3000`。
   + 例子：`--port 3000`
2. `--log-dir <log_dir>`: 指定日志文件夹，如果没有设置，日志将不会被保存。
   + 例子：`--log-dir ./logs`
3. `--version`: 打印系统版本号并退出。