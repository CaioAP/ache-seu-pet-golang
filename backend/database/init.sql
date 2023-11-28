DROP SCHEMA IF EXISTS public;

CREATE SCHEMA IF NOT EXISTS public;

DROP TABLE IF EXISTS users;

CREATE TABLE users (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "email" VARCHAR NOT NULL,
  "phone" VARCHAR(11),
  "birthday" VARCHAR(10),
  "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  "password" VARCHAR NOT NULL
);

CREATE TABLE addresses (
  "user_id" INT REFERENCES users ON DELETE CASCADE,
  "cep" VARCHAR(8) NOT NULL,
  "uf" VARCHAR(2) NOT NULL,
  "city" VARCHAR NOT NULL,
  "neighborhood" VARCHAR NOT NULL,
  "street" VARCHAR NOT NULL,
  "complement" VARCHAR,
  "number" INT,
  PRIMARY KEY (user_id)
);

INSERT INTO users (id, name, email, phone, birthday, password)
VALUES (1, 'Teste', 'teste@teste.com', NULL, NULL, '$2y$14$WbqsVi9WjxtDWkxsuj7NcOzz3J6SqvEuhvbiB1bxvToCavdqY4QMm');