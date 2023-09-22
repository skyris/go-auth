-- +goose Up
-- +goose StatementBegin

-- SELECT 'CREATE DATABASE "auth"'
-- WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'auth')\gexec
-- \c auth

-- CREATE SCHEMA IF NOT EXISTS auth;
-- COMMENT ON SCHEMA auth IS 'Authentication/authorization database schema';

-- SET search_path = auth, public;

CREATE TABLE IF NOT EXISTS users (
   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
   username TEXT NOT NULL CHECK (username <> ''),
   email TEXT UNIQUE CHECK (email <> ''),
   password TEXT NOT NULL CHECK (password <> ''),
   created_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
   modified_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
   active BOOLEAN default TRUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;

-- DROP SCHEMA IF EXISTS auth;

-- DROP DATABASE IF EXISTS "auth";
-- +goose StatementEnd
