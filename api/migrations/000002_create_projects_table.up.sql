CREATE TABLE "projects" (
    "id" bigserial PRIMARY KEY,
    "title" varchar NOT NULL,
    "deadline" timestamp NOT NULL ,
    "memo" varchar
)