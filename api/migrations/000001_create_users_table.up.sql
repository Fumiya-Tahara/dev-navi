CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL UNIQUE ,
    "password" varchar NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
)