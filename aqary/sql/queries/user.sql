-- name: CreateUser :exec
INSERT INTO "users"
    ("name", "phone_number")
Values ($1 , $2);

-- name: UpdateUser :exec
UPDATE "users" 
SET "otp" = $1, "otp_expiration_time" = $2 
WHERE "phone_number" = $3;

-- name: GetUser :one
SELECT "id", "name", "phone_number", "otp", "otp_expiration_time", "created_at", "updated_at"
FROM "users"
WHERE "phone_number" = $1
Limit 1;