-- name: GetUserById :one
SELECT *
FROM users
WHERE id = $1;

-- name: GetUserByUsername :one
SELECT *
FROM users
WHERE username = $1;

-- name: ListUsers :many
SELECT *
FROM users;

-- name: CreateUser :execresult
INSERT INTO users (
  username, email
) VALUES (
  $1, $2
);

-- name: UpdateUser :one
UPDATE users
SET username = $2, email = $3
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
