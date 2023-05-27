-- Tabela "user_roles"
CREATE TABLE user_roles (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Inserir as regras existentes
INSERT INTO user_roles (id, name) VALUES ('1', 'admin'), ('2', 'instructor'), ('3', 'student');

-- Tabela "users"
CREATE TABLE users (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    role_id VARCHAR(36) NOT NULL,
    active BOOLEAN NOT NULL DEFAULT true,
    FOREIGN KEY (role_id) REFERENCES user_roles(id)
);

-- Tabela "gyms"
CREATE TABLE gyms (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    gym_name VARCHAR(255) NOT NULL,
    team_name VARCHAR(255) NOT NULL,
    active BOOLEAN NOT NULL DEFAULT true,
    FOREIGN KEY (user_id) REFERENCES users(id));

-- Tabela "students"
CREATE TABLE students (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    gym_id VARCHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    graduation VARCHAR(255) NOT NULL,
    active BOOLEAN NOT NULL DEFAULT true,
    training_time VARCHAR(255) NOT NULL,
    FOREIGN KEY (gym_id) REFERENCES gyms(id)
);
