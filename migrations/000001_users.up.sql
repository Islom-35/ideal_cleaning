CREATE TYPE user_role AS ENUM ('admin', 'user');

CREATE TABLE "users" (
    "id" SERIAL PRIMARY KEY,
    "full_name" VARCHAR,
    "email" VARCHAR,
    "phone" VARCHAR,
    "password" VARCHAR,
    "role" user_role,
    "created_at" TIMESTAMP DEFAULT current_timestamp,
    "updated_at" TIMESTAMP DEFAULT current_timestamp,
    "deleted_at" TIMESTAMP
);