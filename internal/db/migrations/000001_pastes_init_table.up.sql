CREATE TABLE IF NOT EXISTS pastes(
    id           BIGSERIAL PRIMARY KEY,
    file_id      VARCHAR UNIQUE NOT NULL,
    file_content BYTEA,
    created_at   TIMESTAMPTZ DEFAULT now()
);
