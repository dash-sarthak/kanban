-- noinspection SqlNoDataSourceInspectionForFile

-- Set Timezone
SET timezone = 'Asia/Kolkata';

-- Project authors table
CREATE TABLE authors (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    username TEXT NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now()),
    UNIQUE(username)
);

-- Projects table
CREATE TABLE projects (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    author UUID NOT NULL REFERENCES authors(id) ON DELETE CASCADE,
    description TEXT,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now())
);

-- States table
CREATE TABLE states (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    project UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now())
);
-- Components table
CREATE TABLE components (
    component UUID NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    project UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    state UUID NOT NULL REFERENCES states(id) ON DELETE CASCADE,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now())
);

