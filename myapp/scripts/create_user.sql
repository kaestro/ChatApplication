CREATE TABLE "user" (
    "id" SERIAL PRIMARY KEY,
    "userName" VARCHAR(255) NOT NULL UNIQUE,
    "emailAddress" VARCHAR(255) NOT NULL UNIQUE,
    "password" VARCHAR(255) NOT NULL
);

ALTER TABLE "user" RENAME TO "users";

ALTER TABLE "users" RENAME COLUMN "userName" TO "user_name";
ALTER TABLE "users" RENAME COLUMN "emailAddress" TO "email_address";