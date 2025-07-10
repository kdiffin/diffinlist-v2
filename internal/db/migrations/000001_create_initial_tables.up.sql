CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- dnot edit ths anymore get some migrations
CREATE TABLE
    users (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        username TEXT UNIQUE NOT NULL,
        password_hash TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT now (),
        path_to_pfp TEXT NOT NULL DEFAULT '/assets/images/goku.jpg',
        updated_at TIMESTAMP NOT NULL DEFAULT now ()
    );

CREATE TABLE
    sessions (
        id TEXT PRIMARY KEY DEFAULT gen_random_uuid (),
        user_id UUID REFERENCES users (id) ON DELETE CASCADE,
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
        authorId UUID REFERENCES users (id) ON DELETE CASCADE,
        -- no dude with the same two playlists
        UNIQUE (authorId, name)
    );

CREATE UNIQUE INDEX playlist_idx ON playlists (authorId, name);

ALTER TABLE users
ADD COLUMN default_playlist_id UUID NOT NULL REFERENCES playlists (id) ON DELETE CASCADE;

CREATE TABLE
    songs (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        createdAt TIMESTAMP NOT NULL DEFAULT now (),
        updated_at TIMESTAMP NOT NULL DEFAULT now (),
        name TEXT NOT NULL,
        pictureUrl TEXT,
        songUrl TEXT,
        genre TEXT,
        artist TEXT,
        album TEXT,
        description TEXT,
        likes INT4,
        dislikes INT4,
        options JSONB,
        authorId UUID REFERENCES users (id) ON DELETE CASCADE
    );

CREATE UNIQUE INDEX song_idx ON songs (authorId, name);

CREATE TABLE
    song_playlist_rel (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        song_id UUID REFERENCES songs (id) ON DELETE CASCADE,
        playlist_id UUID REFERENCES playlists (id) ON DELETE CASCADE
    );

CREATE UNIQUE INDEX song_playlist_rel_idx ON song_playlist_rel (song_id, playlist_id);