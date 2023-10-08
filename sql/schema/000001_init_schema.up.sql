-- A static table for states
CREATE TABLE states (
    id SERIAL PRIMARY KEY,
    state TEXT NOT NULL
)
-- Static states
INSERT INTO states (state)
VALUES
('todo'),
('in_progress'),
('done');

-- Project authors table
CREATE TABLE authors (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    username TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Projects table
CREATE TABLE projects (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    author UUID NOT NULL REFERENCES authors(id) ON DELETE CASCADE,
    description TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Components table
-- All components across all tables
CREATE TABLE components (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    project UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Component state table
CREATE TABLE component_states (
    component UUID NOT NULL REFERENCES components(id) ON DELETE CASCADE,
    project UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    state INTEGER REFERENCES states(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY (component, project)
);

