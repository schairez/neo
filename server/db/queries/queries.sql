-- name: CreateUser :one
INSERT INTO users (user_name, pass_key, first_name, last_name)
VALUES ($1 $2 $3 $4)
RETURNING *;
-- name: CreateNB :one
INSERT INTO notebooks () -- name: CreateSection :one
INSERT INTO sections() -- name: CreateNote :one  
    -- name: CreateTag :one