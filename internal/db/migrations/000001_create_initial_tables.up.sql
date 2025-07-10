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
    playlists (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        createdAt TIMESTAMP NOT NULL DEFAULT now (),
        updated_at TIMESTAMP NOT NULL DEFAULT now (),
        name TEXT NOT NULL,
        description TEXT,
        genre TEXT,
        pictureUrl TEXT,
        likes INT4,
        dislikes INT4,
        saves INT4,
        authorId UUID REFERENCES users (id),
        -- no dude with the same two playlists
        UNIQUE (authorName, name)
    );

CREATE UNIQUE INDEX playlist_idx ON playlists (authorName, name);

CREATE TABLE
    songs (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        createdAt TIMESTAMP NOT NULL DEFAULT now (),
        updated_at TIMESTAMP NOT NULL DEFAULT now (),
        name TEXT NOT NULL,
        pictureUrl TEXT,
        songUrl TEXT,
        genre TEXT,
        likes INT4,
        dislikes INT4,
        artist TEXT,
        album TEXT,
        description TEXT,
        authorName TEXT,
        options JSONB
    );

CREATE UNIQUE INDEX song_idx ON songs (authorName, name);

CREATE TABLE
    song_playlist_rel (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        song_id UUID REFERENCES songs (id),
        playlist_id UUID REFERENCES playlists (id)
    );

CREATE UNIQUE INDEX song_playlist_rel_idx ON song_playlist_rel (song_id, playlist_id);