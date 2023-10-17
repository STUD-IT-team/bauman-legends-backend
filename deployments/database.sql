create
extension if not exists "uuid-ossp";

create
extension if not exists "pgcrypto";

drop table if exists "user";
drop table if exists "team_task";
drop table if exists "team_secret";
drop table if exists "team";
drop table if exists "role";
drop table if exists "answer";
drop table if exists "task";
drop table if exists "task_type";
drop table if exists "task_difficulty";
drop table if exists "answer_secret";
drop table if exists "answer_type";
drop table if exists "secret";

create table "team"
(
    id    uuid not null default uuid_generate_v4(),
    title text not null,

    primary key (id)
);

create table "role"
(
    id    int generated always as identity,
    title text not null,

    primary key (id)
);

create table "user"
(
    id              uuid    not null default uuid_generate_v4(),
    password        text    not null,
    phone_number    text    not null,
    email           text    not null unique,
    email_confirmed boolean not null default false,
    telegram        text    not null,
    vk              text    not null,
    "group"         text    not null,
    name            text    not null,
    team_id         uuid             default null,
    role_id         int              default null,
    is_admin        boolean          default false,

    primary key (id),

    foreign key (team_id)
        references "team" (id)
        on delete set default,
    foreign key (role_id)
        references "role" (id)
        on delete set default
);

create table "task_type"
(
    id    int generated always as identity,
    title text not null,

    primary key (id)
);

create table "task"
(
    id             uuid not null default uuid_generate_v4(),
    title          text not null,
    description    text,
    time_limit interval not null,
    type_id        int  not null,
    max_points     int  not null,
    min_points     int  not null,
    answer_type_id int  not null,
    primary key (id),

    foreign key (type_id)
        references "task_type" (id)
);

create table "team_task"
(
    id                uuid        not null default uuid_generate_v4(),
    task_id           uuid        not null,
    team_id           uuid        not null,
    start_time        timestamptz not null default now(),
    end_time          timestamptz          default null,
    additional_points int         not null default 0,
    answer_text       text                 default null,
    answer_image_url  text                 default null,
    result            bool                 default null,
    primary key (id),

    foreign key (task_id)
        references "task" (id),
    foreign key (team_id)
        references "team" (id)
);


create table "secret"
(
    id          uuid not null default uuid_generate_v4(),
    title       text not null,
    description text,

    primary key (id)
);

create table "answer_secret"
(
    id             uuid not null default uuid_generate_v4(),
    secret_id      uuid not null,
    answer_type_id int  not null,
    data           text not null,

    primary key (id),

    foreign key (secret_id)
        references "secret" (id),
    foreign key (answer_type_id)
        REFERENCES "answer_type" (id)
);

create table "team_secret"
(
    id         uuid not null default uuid_generate_v4(),
    secret_id  uuid not null,
    team_id    uuid not null,
    start_time timestamptz   default now(),
    end_time   timestamptz   default null,

    primary key (id),

    foreign key (secret_id)
        references "secret" (id),
    foreign key (team_id)
        references "team" (id)
);
