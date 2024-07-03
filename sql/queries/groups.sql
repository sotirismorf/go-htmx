-- name: CreateGroup :exec
INSERT INTO groups (name, location) VALUES ($1, $2);
