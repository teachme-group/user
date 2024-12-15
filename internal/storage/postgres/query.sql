-- name: CreateUser :one
insert into "user" (login, email, password, created_at)
values ($1, $2, $3, $4)
    returning id, login, email, password, created_at;

-- name: ValidateUserSignUp :one
select exists(select 1 from "user" where login = $1 or email = $2) as exists;