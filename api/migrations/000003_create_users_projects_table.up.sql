-- api/migrations/000003_create_users_projects_table.up.sql
CREATE TABLE "users_projects" (
    "user_id" bigint NOT NULL,
    "project_id" bigint NOT NULL,
    PRIMARY KEY ("user_id", "project_id"),
    FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("project_id") REFERENCES "projects" ("id") ON DELETE CASCADE
);