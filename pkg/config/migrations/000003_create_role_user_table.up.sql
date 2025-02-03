CREATE TABLE IF NOT EXISTS role_user (
    id SERIAL PRIMARY KEY NOT NULL,
    role_id INT NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (role_id) REFERENCES roles (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

INSERT INTO role_user (role_id, user_id) VALUES (1, 1);
INSERT INTO role_user (role_id, user_id) VALUES (2, 1);
INSERT INTO role_user (role_id, user_id) VALUES (2, 2);
INSERT INTO role_user (role_id, user_id) VALUES (3, 2);
INSERT INTO role_user (role_id, user_id) VALUES (1, 3);