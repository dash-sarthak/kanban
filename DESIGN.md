# DB Design

## authors

| id        | name         |
|-----------|--------------|
| UUID PKEY | TEXT NOTNULL |

## access_groups (TODO)
- An access group can have one or more authors.
- By default, an access group will have one author.
- If the author grants permission to another author to access a project, a new access group consisting of those 2 author will be created.
- Each project is assigned to one access_group.


## projects

| id        | name         | author                        | description | access_group                        |
|-----------|--------------|-------------------------------|-------------|-------------------------------------|
| UUID PKEY | TEXT NOTNULL | UUID NOTNULL FKEY authors(id) | TEXT        | UUID NOTNULL FKEY access_groups(id) |

## states

| id        | name         | project                         | order           |
|-----------|--------------|---------------------------------|-----------------|
| UUID PKEY | TEXT NOTNULL | UUID NOTNULL FKEY projects(id ) | INTEGER NOTNULL |

By default, each project will have the following states:

- todo
- doing
- done

User can delete existing states, but there needs to be **at least** 2 states in a project. User will have the option of performing the following actions:

- DELETE (restriction: if there are 2 states, deletion not possible)
- RENAME
- ADD NEW
- CHANGE ORDER (TBD)

## components

| id        | name         | project                        |
|-----------|--------------|--------------------------------|
| UUID PKEY | TEXT NOTNULL | UUID NOTNULL FKEY projects(id) |

- Each project can have as many components as possible.

## component_states

| component                        | project                   | state                   |
|----------------------------------|---------------------------|-------------------------|
| UUID NOTNULL FKEY components(id) | UUID NOTNULL projects(id) | UUID NOTNULL states(id) |

PRIMARY KEY (component, project)

- Each component must be part of a state that is defined for that project
- Following actions are premissible:
- CREATE
- UPDATE STATE
- RENAME
- DELETE

