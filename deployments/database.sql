set search_path = "public";

create extension if not exists "uuid-ossp";

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

create table "team" (
	id		uuid	not null	default uuid_generate_v4(),
	title	text	not null,

	primary key (id)
);

create table "role" (
	id		int		GENERATED ALWAYS AS IDENTITY,
	title	text	not null,

	primary key (id)
);

-- insert into "role" (title) values ('Участник');
-- insert into "role" (title) values ('Зам');
-- insert into "role" (title) values ('Капитан');

create table "user" (
    id		        uuid	not null	default uuid_generate_v4(),
	password 	    text    not null,
	phone_number	text	not null,
	email		    text	not null,
	email_confirmed	boolean not null    default false,
	telegram	    text	not null,
	vk		        text	not null,
	"group"	        text	not null,
	name		    text	not null,
	team_id 	    uuid 	            default null,
	role_id		    int		            default null,
	is_admin	    boolean		        default false,

    primary key (id),

    foreign key (team_id)
        references "team" (id)
            on delete set default,
    foreign key (role_id)
        references "role" (id)
            on delete set default
);

create table "task_type" (
	id		int	     not null serial,
	title	text	not null,

	primary key (id)
);
insert into "task_type" (title) values ("текст");
insert into "task_type" (title) values ("текст + фото");
insert into "task_type" (title) values ("фото");
insert into "task_type" (title) values ("qr-code");
insert into "task_type" (title) values ("любое другое");

create table "task_difficulty" (
	id		int		generated always as identity,
	title	text	not null,

	primary key (id)
);

create table "task" (
    id		        uuid	    not null	default uuid_generate_v4(),
	title		    text	    not null,
   	description	    text,
	time_limit	    interval    not null,
	difficulty_id	int 	    not null,
    type_id		    int	        not null,

    primary key (id),

	foreign key (difficulty_id)
		references "task_difficulty" (id),
	foreign key (type_id)
		references "task_type" (id)
);

create table "team_task" (
    id		            uuid        not null	default uuid_generate_v4(),
	task_id		        uuid	    not null,
    team_id 	        uuid	    not null,
	start_time	        timestamptz not null 	default now(),
	end_time    	    timestamptz 	        default null,
	additional_points   int 		not null    default 0,

	primary key (id),

    foreign key (task_id)
        references "task" (id),
    foreign key (team_id)
        references "team" (id)
);

create table "answer_type" (
	id		int		generated always as identity,
	title	text	not null,

	primary key (id)
);

create table "answer" (
    id		        uuid    not null	default uuid_generate_v4(),
    task_id		    uuid    not null,
    answer_type_id	int		not null,
    data 		    text 	not null,

    primary key (id),

	foreign key (task_id)
		references "task" (id),
	foreign key (answer_type_id)
		references "answer_type" (id)
);

create table "secret" (
	id		    uuid    not null	default uuid_generate_v4(),
	title		text    not null,
	description	text,

	primary key (id)
);

create table "answer_secret" (
    id		        uuid    not null	default uuid_generate_v4(),
	secret_id	    uuid 	not null,
    answer_type_id	int		not null,
    data 		    text 	not null,

    primary key (id),

	foreign key (secret_id)
		references "secret" (id),
	foreign key (answer_type_id)
		REFERENCES "answer_type" (id)
);

create table "team_secret" (
	id		    uuid        not null	default uuid_generate_v4(),
	secret_id 	uuid	    not null,
    team_id 	uuid 	    not null,
    start_time  timestamptz 	        default now(),
    end_time    timestamptz 	        default null,

    primary key (id),

	foreign key (secret_id)
		references "secret" (id),
	foreign key (team_id)
		references "team" (id)
);
