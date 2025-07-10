

-- name: SignUpAndCreateDefaultPlaylist :exec
WITH
  new_playlist AS (
    INSERT INTO playlists (name, description)
    VALUES ('default', 'This is your default playlist. you can save songs here.')
    RETURNING id
  )
    INSERT INTO users (username, password_hash, default_playlist_id, path_to_pfp)
    VALUES ($1, $2, (SELECT id FROM new_playlist), $3)
    RETURNING id;


-- name: SignUpAndCreateDefaultPlaylistNoPfp :exec
WITH
    new_playlist AS (
    INSERT INTO playlists (name, description)
    VALUES ('default', 'This is your default playlist. you can save songs here.')
    RETURNING id
  ) INSERT INTO users (username, password_hash, default_playlist_id)
    VALUES ($1, $2, (SELECT id FROM new_playlist))
    RETURNING id;

-- name: GetUserCredentials :one
SELECT id, password_hash FROM users WHERE username=$1;

-- name: GetUsers :many
SELECT username, path_to_pfp  FROM users;


-- name: CreateSession :exec
INSERT INTO sessions (id, user_id) VALUES($1,$2);

-- name: DeleteSession :exec
DELETE FROM sessions WHERE id = $1;

-- name: GetSessionValues :one
SELECT id, user_id FROM sessions WHERE id = $1;

-- name: GetUserInfo :one
SELECT username, path_to_pfp, default_playlist_id FROM users WHERE id = $1;