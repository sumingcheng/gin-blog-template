CREATE TABLE IF NOT EXISTS "user"
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(20)  NOT NULL,
    password   VARCHAR(72)  NOT NULL,
    created_at BIGINT       NOT NULL DEFAULT (EXTRACT(EPOCH FROM NOW())::BIGINT),
    update_at  BIGINT       DEFAULT NULL,
    delete_at  BIGINT       DEFAULT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_username ON "user" (name);

INSERT INTO "user" (name, password)
VALUES ('admin', '$2a$10$CVZr2jIPRIbNPCsXCAMWo.I5d74pIImtJ7dirvl17PV9n2TRUaHZW')
ON CONFLICT (name) DO NOTHING;

CREATE TABLE IF NOT EXISTS blog
(
    id         SERIAL PRIMARY KEY,
    user_id    INT          NOT NULL,
    title      VARCHAR(100) NOT NULL,
    article    TEXT         NOT NULL,
    created_at BIGINT       NOT NULL DEFAULT (EXTRACT(EPOCH FROM NOW())::BIGINT),
    update_at  BIGINT       DEFAULT NULL,
    delete_at  BIGINT       DEFAULT NULL
);

CREATE INDEX IF NOT EXISTS idx_user_id ON blog (user_id);

INSERT INTO blog (user_id, title, article)
VALUES (1, 'The Magic of the Forest at Dawn',
        'As the first light of dawn stretches across the forest, the world slowly awakens to a symphony of bird calls and rustling leaves. This serene moment captures the essence of nature''s quiet majesty. In the early morning light, the trees stand tall, their leaves glistening with dew, creating a shimmering tapestry of greens and golds. The forest floor, a mosaic of ferns and fallen leaves, invites wanderers to lose themselves in its depths. The air, fresh and crisp, carries the earthy scent of moss and wood. Here, amidst the ancient groves, one can truly feel the pulse of the earth and its ceaseless, tranquil breath. This magical time offers a profound peace and a rare solitude that rejuvenates the spirit and clears the mind, reminding us of the simple beauty that nature generously offers to those who seek it.'),
       (1, 'The Stars Above: Gazing into Infinity',
        'On a clear night, far from the glaring lights of the city, the sky reveals its true self—a vast canvas sprinkled with stars, planets, and distant galaxies. As you gaze upwards, the universe seems both immeasurably vast and surprisingly intimate. Each star is a sun, possibly orbited by its own planets, holding secrets of distant worlds. The constellations, patterns imprinted in human lore, tell ancient stories and guide the explorers'' path. The Milky Way stretches across the sky, a swath of milky brightness that speaks to our galaxy''s depth and complexity. To observe the night sky is to look back in time, for the light from these stars has traveled unimaginable distances to reach us. It is a humbling experience that challenges the soul and expands the mind, offering a silent yet overwhelming proof of the vastness and beauty of our universe.');
