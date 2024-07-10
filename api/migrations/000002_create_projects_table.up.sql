CREATE TABLE "projects" (
    "id" bigserial PRIMARY KEY,
    "title" varchar NOT NULL,
    "status" smallint NOT NULL,
    "deadline" timestamp NOT NULL,
    "memo" varchar,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
)