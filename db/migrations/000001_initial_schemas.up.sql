CREATE TABLE IF NOT EXISTS accounts (
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(255)    NOT NULL,
    cpf             VARCHAR(11)     NOT NULL    UNIQUE,
    secret_hash     VARCHAR(255)    NOT NULL,
    secret_salt     VARCHAR(255)    NOT NULL,
    balance         BIGINT          NOT NULL    DEFAULT 10000,
    created_at      TIMESTAMP       NOT NULL    DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP       NOT NULL    DEFAULT CURRENT_TIMESTAMP,
    deleted_at      TIMESTAMP       NULL
);

CREATE TABLE IF NOT EXISTS transfers (
    id                  SERIAL          PRIMARY KEY,
    origin_id           INTEGER         NOT NULL,
    destination_id      INTEGER         NOT NULL,
    amount              BIGINT          NOT NULL,
    created_at          TIMESTAMP       NOT NULL    DEFAULT CURRENT_TIMESTAMP
);

