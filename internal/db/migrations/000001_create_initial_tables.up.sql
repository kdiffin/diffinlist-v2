CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE
    users (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        username TEXT UNIQUE NOT NULL,
        password_hash TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT now (),
        updated_at TIMESTAMP NOT NULL DEFAULT now ()
    );

CREATE TABLE
    sessions (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        user_id UUID REFERENCES users (id),
        created_at TIMESTAMP NOT NULL DEFAULT now ()
    );

CREATE TABLE
    playlist (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        createdAt TIMESTAMP NOT NULL DEFAULT now (),
        updated_at TIMESTAMP NOT NULL DEFAULT now (),
        name TEXT NOT NULL,
        genre TEXT,
        pictureUrl TEXT,
        authorName TEXT,
        -- no dude with the same two playlists
        UNIQUE (authorName, name)
    );

CREATE UNIQUE INDEX playlist_idx ON songs (authorName, name);

CREATE TABLE
    song (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        createdAt TIMESTAMP NOT NULL DEFAULT now (),
        updated_at TIMESTAMP NOT NULL DEFAULT now (),
        name TEXT NOT NULL,
        pictureUrl TEXT,
        songUrl TEXT,
        genre TEXT,
        rating INT4,
        artist TEXT,
        album TEXT,
        description TEXT,
        authorName TEXT,
        playlistId UUID REFERENCES playlist (id),
        options JSONB
    );

CREATE UNIQUE INDEX song_idx ON songs (authorName, name, playlistId);