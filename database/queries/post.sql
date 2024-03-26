-- name: FindPostById :one
SELECT pid,title,content FROM tbl_post WHERE pid=$1;

-- name: ListPosts :many
SELECT pid,title FROM tbl_post ORDER BY created_at DESC LIMIT 25 OFFSET $1;

-- name: CreatePost :one
INSERT INTO tbl_post(title,content) VALUES($1,$2) RETURNING pid;

-- name: UpdatePost :exec
UPDATE tbl_post SET title=coalesce(sqlc.narg('title'),title) AND content=coalesce(sqlc.narg('content'),content) WHERE pid=$1;

-- name: DeletePost :exec
DELETE FROM tbl_post WHERE pid = $1;
