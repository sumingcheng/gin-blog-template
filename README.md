# Gin-Blog 模板

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

**执行SQL**

```
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

## 配置

### 环境变量