-- name: GetPastes :many
SELECT * FROM pastes;

-- name: GetPaste :one
SELECT * FROM pastes WHERE file_id = $1 LIMIT 1;

-- name: CreatePaste :exec
INSERT INTO pastes (
    file_id, file_content
) VALUES ($1, $2);