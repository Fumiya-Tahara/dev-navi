CREATE TABLE "milestones" (
    "id" bigserial PRIMARY KEY,
    "project_id" bigint NOT NULL,
    "title" varchar NOT NULL,
    "status" smallint NOT NULL,
    "deadline" timestamp,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_milestones_projects FOREIGN KEY (project_id) REFERENCES projects(id)
)