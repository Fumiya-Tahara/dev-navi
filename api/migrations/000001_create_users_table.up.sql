CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL UNIQUE ,
    "password" varchar NOT NULL
)