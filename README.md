<h1 align="center">Gin-Blog-Template</h1>
<p align="center">
  <a href="https://raw.githubusercontent.com/sumingcheng/gin-blog/main/LICENSE">
    <img src="https://img.shields.io/github/license/sumingcheng/gin-blog?color=brightgreen" alt="license">
  </a>
  <a href="https://hub.docker.com/repository/docker/smcroot/gin-blog">
    <img src="https://img.shields.io/docker/pulls/smcroot/gin-blog?color=brightgreen" alt="docker pull">
  </a>
  <a href="https://goreportcard.com/report/github.com/sumingcheng/gin-blog">
    <img src="https://goreportcard.com/badge/github.com/sumingcheng/gin-blog" alt="GoReportCard">
  </a>
</p>

## 项目功能
-  **双 Token 登录**：实现了基于令牌的鉴权机制。
-  **日志管理**：引入了 `logrus` 进行日志文件的自动切割和轮换。
-  **配置管理**：使用 `viper` 支持多种配置格式及环境变量的集成。
-  **错误处理**：通过 `translate` 实现错误信息的翻译。
-  **文档生成**：自动生成 `Swagger` API 文档。
-  **数据库操作**：采用 `GORM` 操作 `MySQL` 数据库。
-  **性能监控**：暴露 `Metric` 指标，使用 `Prometheus + Grafana` 监控。
-  **前端技术栈**：前端采用 `vite + react + chakra-ui` 。

## 部署
### 手动部署
进入项目目录

```
make build
make run
```

日志数据将会保存在宿主机的 `/home/logs` 目录，也可以自行修改。

### 基于 Docker 进行部署

`git clone https://github.com/sumingcheng/gin-blog.git`进入项目目录

```
docker-compose up -d
```

```
启动成功访问 ——> 部署地址：5678
```

### 建表语句
**执行SQL**

```sql
create database blog;

create table if not exists user
(
    id       int auto_increment comment '用户id,主键,自增',
    name     varchar(20) not null comment '用户名',
    password char(32)    not null comment '密码md5',
    primary key (id),
    unique key idx_username (name)
) default charset = utf8mb4 comment '用户登录表';

insert into user (name, password)
values ('admin', 'e10adc3949ba59abbe56e057f20f883e'),
       ('user', 'e10adc3949ba59abbe56e057f20f883e');

create table if not exists blog
(
    id          int auto_increment comment '博客id,主键,自增',
    user_id     int          not null comment '作者id',
    title       varchar(100) not null comment '标题',
    article     text         not null comment '文章内容',
    create_time timestamp default current_timestamp comment '创建时间',
    update_time timestamp default current_timestamp on update current_timestamp comment '最后修改时间',
    primary key (id),
    key idx_user_id (user_id)
) default charset = utf8mb4 comment '博客表';

insert into blog (user_id, title, article)
values (1, '博客标题1', '博客内容1'),
       (1, '博客标题2', '博客内容2'),
       (2, '博客标题3', '博客内容3'),
       (2, '博客标题4', '博客内容4'),
       (2, '博客标题5', '博客内容5');
```
### 监控配置

项目启动后，可以直接导入 `grafana` 仪表盘 `deploy/grafana/gin-blog.json` 

![image](https://github.com/user-attachments/assets/a3b15eea-dcf7-4ced-88da-4126d29e6190)


## Token 流程
![Snipaste_2024-07-17_14-20-42](https://github.com/user-attachments/assets/8cea318f-2302-4f19-b5a1-301714d1a00e)

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=sumingcheng/gin-blog-template&type=Timeline)](https://star-history.com/#sumingcheng/gin-blog-template&Timeline)

