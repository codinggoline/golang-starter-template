CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL
);

INSERT INTO roles (name) VALUES ('ROLE_GUEST');
INSERT INTO roles (name) VALUES ('ROLE_USER');
INSERT INTO roles (name) VALUES ('ROLE_ADMIN');