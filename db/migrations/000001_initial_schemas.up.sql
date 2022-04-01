CREATE TABLE IF NOT EXISTS accounts (
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(255)    NOT NULL,
    cpf             VARCHAR(11)     NOT NULL    UNIQUE,
    secret_hash     VARCHAR(255)    NOT NULL,
    secret_salt     VARCHAR(255)    NOT NULL,
    balance         NUMERIC(10,2)   NOT NULL    DEFAULT 100.00,
    created_at      TIMESTAMP       NOT NULL    DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS transfers (
    id                  SERIAL          PRIMARY KEY,
    origin_id           INTEGER         NOT NULL,
    destination_id      INTEGER         NOT NULL,
    amount              NUMERIC(10,2)   NOT NULL,
    created_at          TIMESTAMP       NOT NULL    DEFAULT CURRENT_TIMESTAMP
);

