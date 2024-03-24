-- name: GetUserById :one
SELECT * FROM tbl_user WHERE uid = $1 AND deleted_at IS NULL LIMIT 1;

-- name: GetUserByUsername :one
SELECT uid,username,password,is_admin FROM tbl_user WHERE username = $1 AND deleted_at IS NULL LIMIT 1;

-- name: CreateUser :exec
INSERT INTO tbl_user(username,password) VALUES($1,$2);

-- name: UpdateUser :exec
UPDATE tbl_user SET username = coalesce(sqlc.narg('username'),username),password = coalesce(sqlc.narg('password'),password) WHERE uid = $1;

-- name: SoftDeleteUser :exec
UPDATE tbl_user SET deleted_at= CURRENT_TIMESTAMP WHERE uid = $1;

-- name: DeleteUser :exec
DELETE FROM tbl_user WHERE uid = $1;
