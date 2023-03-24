CREATE TABLE IF NOT EXISTS users
(
    id            uuid         not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
    );


CREATE TABLE IF NOT EXISTS todo_lists
(
    id          uuid         not null unique,
    title       varchar(255) not null,
    description varchar(255)
    );

CREATE TABLE IF NOT EXISTS users_list
(
    id      uuid                                              not null unique,
    user_id uuid references users (id) on delete cascade      not null,
    list_id uuid references todo_lists (id) on delete cascade not null
    );

CREATE TABLE IF NOT EXISTS todo_items
(
    id          uuid         not null unique,
    title       varchar(255) not null,
    description varchar(255),
    done        boolean      not null default false
    );

CREATE TABLE IF NOT EXISTS lists_items
(
    id      uuid                                              not null unique,
    item_id uuid references todo_items (id) on delete cascade not null,
    list_id uuid references todo_lists (id) on delete cascade not null
    );