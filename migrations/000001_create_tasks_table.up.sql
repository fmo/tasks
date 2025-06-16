DROP TABLE IF EXISTS tasks;

CREATE TABLE tasks(
    id SERIAL PRIMARY KEY,
    name VARCHAR(250) NOT NULL,
    owner VARCHAR(50) NOT NULL,
    priority int NOT NULL
);