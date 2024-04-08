DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS address_data;
DROP TABLE IF EXISTS geo_data;

CREATE TABLE users(
    id       SERIAL PRIMARY KEY,
    login    VARCHAR(255),
    password VARCHAR(255)
);

CREATE TABLE address_data(
    id      SERIAL PRIMARY KEY,
    address VARCHAR(255),
    data    VARCHAR(255)
);

CREATE TABLE geo_data(
    id SERIAL PRIMARY KEY,
    geo VARCHAR(255),
    data VARCHAR(255)
);