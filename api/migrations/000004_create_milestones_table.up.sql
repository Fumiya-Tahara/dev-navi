CREATE TABLE "milestones" (
    "id" bigserial PRIMARY KEY,
    "project_id" bigint NOT NULL,
    "title" varchar NOT NULL,
    "deadline" timestamp,
    CONSTRAINT fk_milestones_projects FOREIGN KEY (project_id) REFERENCES projects(id)
)