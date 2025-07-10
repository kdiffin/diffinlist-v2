DROP INDEX IF EXISTS song_playlist_rel_idx;

DROP TABLE IF EXISTS song_playlist_rel;

DROP INDEX IF EXISTS song_idx;

DROP TABLE IF EXISTS songs;

ALTER TABLE users
DROP default_playlist_id;

DROP INDEX IF EXISTS playlist_idx;

DROP TABLE IF EXISTS playlists;

DROP TABLE IF EXISTS sessions;

DROP TABLE IF EXISTS users;