CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    firstname VARCHAR(100) NOT NULL,
    lastname VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(255) NOT NULL,
    date_of_birth DATE NOT NULL,
    username VARCHAR(50) NOT NULL,
    phone VARCHAR(15) NOT NULL,
    gender VARCHAR(10) NOT NULL,
    avatar VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (email, username, phone)
);

INSERT INTO users (firstname, lastname, email, password, date_of_birth, username, phone, gender, avatar) VALUES ('John', 'Doe', 'johndoe@user.com', 'password', '1990-01-01', 'johndoe', '1234567890', 'Male', 'https://cdn.pixabay.com/photo/2020/07/01/12/58/icon-5359553_960_720.png');
INSERT INTO users (firstname, lastname, email, password, date_of_birth, username, phone, gender, avatar) VALUES ('Jane', 'Doe', 'janedoe@user.com', 'password', '1995-01-01', 'janedoe', '1234367890', 'Female', 'https://cdn.pixabay.com/photo/2020/07/01/12/58/icon-5359554_960_720.png');
INSERT INTO users (firstname, lastname, email, password, date_of_birth, username, phone, gender, avatar) VALUES ('Marcus', 'Doe', 'marcusdoe@user.com', 'password', '1970-01-01', 'marcusdoe', '1234567877', 'Male', 'https://cdn.pixabay.com/photo/2020/07/01/12/58/icon-5359553_960_720.png');