CREATE TABLE users
(
    id          SERIAL PRIMARY KEY,
    telegram_id BIGINT
);

CREATE TABLE card_types
(
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE cards
(
    id        SERIAL PRIMARY KEY,
    name      TEXT,
    image_url TEXT,
    price     INT,
    weight    INT,
    type_id   INT REFERENCES card_types (id)
);

CREATE TABLE user_cards
(
    user_id  INT REFERENCES users (id),
    card_id  INT REFERENCES cards (id),
    quantity INT,
    PRIMARY KEY (user_id, card_id)
);