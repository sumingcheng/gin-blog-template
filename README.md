<h1 align="center">Gin-Blog-Template</h1>
<p align="center">
  <a href="https://raw.githubusercontent.com/sumingcheng/gin-blog/main/LICENSE"><img src="https://img.shields.io/github/license/sumingcheng/gin-blog?color=brightgreen" alt="license"></a><a href="https://hub.docker.com/repository/docker/smcroot/gin-blog"><img src="https://img.shields.io/docker/pulls/smcroot/gin-blog?color=brightgreen" alt="docker pull"></a><a href="https://goreportcard.com/report/github.com/sumingcheng/gin-blog"><img src="https://goreportcard.com/badge/github.com/sumingcheng/gin-blog" alt="GoReportCard"></a>
</p>



## 项目功能
-  **双 Token 登录**：实现了基于令牌的鉴权机制。
-  **日志管理**：引入了 `logrus` 进行日志文件的自动切割和轮换。
-  **配置管理**：使用 `viper` 配置格式及环境变量的集成。
-  **错误处理**： `translate` 实现错误信息的翻译。
-  **文档生成**：`Swagger` API 文档。
-  **数据库操作**： `GORM` 操作 `MySQL` 数据库。
-  **性能监控**：暴露 `Metric` 指标，使用 `Prometheus + Grafana` 监控。
-  **前端技术栈**：`vite + react + chakra-ui` 。

## 部署

`git clone https://github.com/sumingcheng/gin-blog.git`进入项目目录

### 手动构建镜像
进入项目目录

```
make build
```

### Docker-compose 启动

```
docker-compose up -d
```

```
启动成功访问 ——> 部署地址：5678
```

**注意：启动后立刻请求，可能会有`500`的错误，请等待`MySQL`完全启动后再试**

### 监控配置

项目启动后，可以直接导入 `grafana` 仪表盘 `deploy/grafana/gin-blog.json` 

![image](https://github.com/user-attachments/assets/a3b15eea-dcf7-4ced-88da-4126d29e6190)


## Token 流程
![Snipaste_2024-07-17_14-20-42](https://github.com/user-attachments/assets/8cea318f-2302-4f19-b5a1-301714d1a00e)

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=sumingcheng/gin-blog-template&type=Timeline)](https://star-history.com/#sumingcheng/gin-blog-template&Timeline)

