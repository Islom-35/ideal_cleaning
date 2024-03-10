CREATE TABLE "products" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR,
    "price" int,
    "count" int,
    "created_at" TIMESTAMP DEFAULT current_timestamp,
    "updated_at" TIMESTAMP DEFAULT current_timestamp,
    "deleted_at" TIMESTAMP
);