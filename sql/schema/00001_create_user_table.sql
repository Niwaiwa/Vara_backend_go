-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    PRIMARY KEY (id),
    id UUID NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    nickname VARCHAR(255),
    avatar VARCHAR(255),
    header VARCHAR(255),
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    locale VARCHAR(10) NOT NULL DEFAULT 'en-US',
    last_login TIMESTAMP,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    is_staff BOOLEAN NOT NULL DEFAULT FALSE,
    is_superuser BOOLEAN NOT NULL DEFAULT FALSE
);

-- Setup System user.
INSERT INTO users (id, username, password, email, is_active, is_staff, is_superuser)
VALUES ('00000000-0000-0000-0000-000000000000', 'admin', '$2a$10$ANBEhteXGgVXubIDZpwHqubmG0JNlR32XAEot7i5jlix9HERxuy/q', 'admin@localhost', TRUE, TRUE, TRUE)
ON CONFLICT(id) DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
