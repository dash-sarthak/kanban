# DB Design

## authors
**[ COMPLETED ]**

| id        | name         | username            |
|-----------|--------------|---------------------|
| UUID PKEY | TEXT NOTNULL | TEXT NOTNULL UNIQUE |

## access_groups
**[ TODO ]**
- An access group can have one or more authors.
- By default, an access group will have one author.
- If the author grants permission to another author to access a project, a new access group consisting of those 2 author will be created.
- Each project is assigned to one access_group.


## projects
**[ COMPLETED ]**

| id        | name         | author                        | description |
|-----------|--------------|-------------------------------|-------------|
| UUID PKEY | TEXT NOTNULL | UUID NOTNULL FKEY authors(id) | TEXT        |

## states
**[ IN-PROGRESS ]**

| id        | name         | project                         |
|-----------|--------------|---------------------------------|
| UUID PKEY | TEXT NOTNULL | UUID NOTNULL FKEY projects(id ) |

- [ ] By default, each project will have the following states:
    - todo
    - doing
    - done

- [ ] User can delete existing states, but there needs to be **at least** 2 states in a project. User will have the option of performing the following actions:

    - DELETE (restriction: if there are 2 states, deletion not possible)
    - RENAME
    - ADD NEW
    - CHANGE ORDER (TBD)

## components
**[ IN-PROGRESS ]**

| id        | name         | project                        | state                   |
|-----------|--------------|--------------------------------|-------------------------|
| UUID PKEY | TEXT NOTNULL | UUID NOTNULL FKEY projects(id) | UUID NOTNULL states(id) |


-[ ] Each component must be part of a state that is defined for that project
-[ ] Following actions are permissible:
    - CREATE
    - UPDATE STATE
    - RENAME
    - DELETE
    
