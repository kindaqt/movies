-- Create UUID method
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create DB
CREATE TABLE IF NOT EXISTS movies(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(50) NOT NULL,
    watched BOOLEAN DEFAULT FALSE
);

-- Seed Data
INSERT INTO movies(title) VALUES ('The Shawshank Redemption');
INSERT INTO movies(title) VALUES ('The Godfather');
INSERT INTO movies(title) VALUES ('The Godfather: Part II');
INSERT INTO movies(title) VALUES ('The Dark Knight');
INSERT INTO movies(title) VALUES ('12 Angry Men');
