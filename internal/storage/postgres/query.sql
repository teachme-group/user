-- name: CreateUser :one
insert into "user" (id, login, email, created_at)
values ($1, $2, $3, $4)
    returning id, login, email, created_at;