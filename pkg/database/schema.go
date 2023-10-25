package database

var schema = `
CREATE TABLE if not exists password(
    id SERIAL PRIMARY KEY,
    appName text unique,
    password text
);
`
