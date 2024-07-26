create database blog;

create user 'blog' identified by '123456';

grant all on blog.* to 'blog' @'localhost' identified by 'blog';

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