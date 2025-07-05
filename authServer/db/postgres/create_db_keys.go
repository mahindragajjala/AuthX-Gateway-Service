package postgres

/*
sudo -u postgres psql - Use the postgres user and database

CREATE DATABASE db_auth; - create the database

\c db_auth

CREATE TABLE users (
    id UUID PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_verified BOOLEAN DEFAULT FALSE,
    status TEXT DEFAULT 'inactive',
    last_login TIMESTAMP,
    login_count INTEGER DEFAULT 0,
    role TEXT DEFAULT 'user'
);

*/
