--
-- PostgreSQL database dump
--
-- Dumped from database version 17.4
-- Dumped by pg_dump version 17.4
--
-- PostgreSQL database schema
--
-- Set client encoding
SET
    client_encoding = 'UTF8';

SET
    standard_conforming_strings = on;

SET
    client_min_messages = warning;

-- Create posts table
CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    author VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    slug VARCHAR(255) UNIQUE NOT NULL
);

-- Create index on slug for faster lookups
CREATE INDEX idx_posts_slug ON posts (slug);

--
-- PostgreSQL database dump complete
--
