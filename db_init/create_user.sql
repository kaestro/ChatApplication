CREATE TABLE "users" (
    "id" SERIAL PRIMARY KEY,
    "user_name" VARCHAR(255) NOT NULL UNIQUE,
    "email_address" VARCHAR(255) NOT NULL UNIQUE,
    "password" VARCHAR(255) NOT NULL
);