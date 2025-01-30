-- name: CreateWorld :one
INSERT INTO worlds (name, description) VALUES (?1, ?2) RETURNING *;
