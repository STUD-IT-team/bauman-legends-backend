CREATE TABLE teams (
	line_num	SERIAL,
	team_id		UUID	NOT NULL	PRIMARY KEY,
	title	text	NOT NULL
);

CREATE TABLE roles (
	id		SERIAL,
	role_id	UUID		NOT NULL	PRIMARY KEY,
	role_name	text	NOT NULL
);

CREATE TABLE users (
	id				SERIAL,
	user_id 		UUID		NOT NULL	PRIMARY KEY,
	phone_number	text		NOT NULL,
	email			text		NOT NULL,
	email_confirmed	BOOLEAN		DEFAULT FALSE,
	telegram		text		NOT NULL,
	vk				text		NOT NULL,
	study_group		text		NOT NULL,
	full_name		text		NOT NULL,
	team_id 		UUID 	    DEFAULT NULL,
	FOREIGN KEY (team_id) 
		REFERENCES teams (team_id) 
			ON DELETE SET DEFAULT,
	role_id			UUID 		DEFAULT NULL,
	FOREIGN KEY (role_id) 
		REFERENCES roles (role_id) 
			ON DELETE SET DEFAULT,
	admin_permissions	BOOLEAN	DEFAULT FALSE
);

CREATE TABLE types_tasks (
	id	        SERIAL,
	type_id		UUID		NOT NULL	PRIMARY KEY,
	task_type	text		NOT NULL
);

CREATE TYPE type_difficulty AS ENUM ('1', '2', '3', '4', '5', '6', '7', '8', '9', '10');

CREATE TABLE tasks (
    id	        SERIAL,
	task_id 	UUID	NOT NULL PRIMARY KEY,
	task_name	text	NOT NULL,
    description	text,
	correct_ans     text[]    DEFAULT array[]::text[],
	qr_or_text_resp smallint  NOT NULL,
    execution_time  text      NOT NULL,
    task_difficulty type_difficulty  NOT NULL,
    task_type_id    UUID,
    FOREIGN KEY (task_type_id) 
			REFERENCES types_tasks (type_id)
);

CREATE TABLE ongoing_tasks (
	id          	SERIAL,
	ongoing_task_id	UUID		NOT NULL	PRIMARY KEY,
   	task_id	    	UUID,
    FOREIGN KEY (task_id) 
		REFERENCES tasks (task_id),
   	team_id     	UUID,
	FOREIGN KEY (team_id) 
		REFERENCES teams (team_id),
    start_time  	TIMESTAMPTZ DEFAULT NOW(),
    end_time    	TIMESTAMPTZ DEFAULT NULL,
    extra_point 	INT         DEFAULT 0
);

CREATE TABLE secrets (
  	id              SERIAL,
	secret_id		UUID	NOT NULL	PRIMARY KEY,
    secret_name     text	NOT NULL,
    description     text,
    correct_ans     text[]           DEFAULT array[]::text[],
    qr_or_text_resp smallint         NOT NULL
);

CREATE TABLE ongoing_secrets (
    id	SERIAL,
	ongoing_secret_id	UUID	NOT NULL	PRIMARY KEY,
    secret_id	UUID,
    FOREIGN KEY (secret_id) 
		REFERENCES secrets (secret_id),
    team_id		UUID,
	FOREIGN KEY (team_id) 
		REFERENCES teams (team_id),
    start_time  	TIMESTAMPTZ DEFAULT NOW(),
    end_time    	TIMESTAMPTZ DEFAULT NULL
)
