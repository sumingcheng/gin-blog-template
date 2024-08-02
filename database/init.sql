CREATE TABLE IF NOT EXISTS user
(
    id        INT AUTO_INCREMENT COMMENT '自增',
    name      VARCHAR(20) NOT NULL COMMENT '用户名',
    password  CHAR(32)    NOT NULL COMMENT '密码MD5',
    update_at BIGINT DEFAULT NULL COMMENT '最后修改时间',
    delete_at BIGINT DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (id),
    UNIQUE KEY idx_username (name)
    ) DEFAULT CHARSET = utf8mb4 COMMENT ='用户登录表';


insert into user (name, password)
values ('admin', 'e10adc3949ba59abbe56e057f20f883e');

CREATE TABLE IF NOT EXISTS blog
(
    id        INT AUTO_INCREMENT COMMENT '自增',
    user_id   INT          NOT NULL COMMENT '作者ID',
    title     VARCHAR(100) NOT NULL COMMENT '标题',
    article   TEXT         NOT NULL COMMENT '文章内容',
    update_at BIGINT DEFAULT NULL COMMENT '最后修改时间',
    delete_at BIGINT DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (id),
    KEY idx_user_id (user_id)
    ) DEFAULT CHARSET = utf8mb4 COMMENT ='博客表';


insert into blog (user_id, title, article)
values (1, 'The Magic of the Forest at Dawn',
        'As the first light of dawn stretches across the forest, the world slowly awakens to a symphony of bird calls and rustling leaves. This serene moment captures the essence of nature''s quiet majesty. In the early morning light, the trees stand tall, their leaves glistening with dew, creating a shimmering tapestry of greens and golds. The forest floor, a mosaic of ferns and fallen leaves, invites wanderers to lose themselves in its depths. The air, fresh and crisp, carries the earthy scent of moss and wood. Here, amidst the ancient groves, one can truly feel the pulse of the earth and its ceaseless, tranquil breath. This magical time offers a profound peace and a rare solitude that rejuvenates the spirit and clears the mind, reminding us of the simple beauty that nature generously offers to those who seek it.'),
       (1, 'The Stars Above: Gazing into Infinity',
        'On a clear night, far from the glaring lights of the city, the sky reveals its true self—a vast canvas sprinkled with stars, planets, and distant galaxies. As you gaze upwards, the universe seems both immeasurably vast and surprisingly intimate. Each star is a sun, possibly orbited by its own planets, holding secrets of distant worlds. The constellations, patterns imprinted in human lore, tell ancient stories and guide the explorers'' path. The Milky Way stretches across the sky, a swath of milky brightness that speaks to our galaxy''s depth and complexity. To observe the night sky is to look back in time, for the light from these stars has traveled unimaginable distances to reach us. It is a humbling experience that challenges the soul and expands the mind, offering a silent yet overwhelming proof of the vastness and beauty of our universe.These articles offer a glimpse into the beauty of the natural world, whether it''s the quiet majesty of a forest at dawn or the awe-inspiring expanse of the night sky. Each piece invites readers to reflect on their own place in the universe and the wonders that surround us daily.');