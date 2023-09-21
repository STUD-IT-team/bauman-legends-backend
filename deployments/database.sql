CREATE TABLE teams (
	line_num	SERIAL,
	id		text	NOT NULL	PRIMARY KEY,
	title	text	NOT NULL
);

CREATE TABLE roles (
	line_num		SERIAL,
	id	INT		NOT NULL	PRIMARY KEY,
	title	text	NOT NULL
);

CREATE TABLE users (
	line_num				SERIAL,
	id 		text		NOT NULL	PRIMARY KEY,
	password text	NOT NULL,
	phone_number	text		NOT NULL,
	email			text		NOT NULL,
	email_confirmed	BOOLEAN		DEFAULT FALSE,
	telegram		text		NOT NULL,
	vk				text		NOT NULL,
	study_group		text		NOT NULL,
	fio		text		NOT NULL,
	team_id 		text 	    DEFAULT NULL,
	FOREIGN KEY (team_id) 
		REFERENCES teams (id) 
			ON DELETE SET DEFAULT,
	role_id			INT 		DEFAULT NULL,
	FOREIGN KEY (role_id) 
		REFERENCES roles (id) 
			ON DELETE SET DEFAULT,
	is_admin	BOOLEAN	DEFAULT FALSE
);

CREATE TABLE types_tasks (
	line_num	        SERIAL,
	id		INT		NOT NULL	PRIMARY KEY,
	title	text		NOT NULL
);

CREATE TABLE difficulties_tasks (
	line_num	        SERIAL,
	id		INT		NOT NULL	PRIMARY KEY,
	title	text		NOT NULL
);

CREATE TABLE tasks (
    line_num	        SERIAL,
	id 	text	NOT NULL PRIMARY KEY,
	title	text	NOT NULL,
    description	text,
	time_limit	text	NOT NULL,
	difficulty_id			INT 		DEFAULT NULL,
	FOREIGN KEY (difficulty_id) 
		REFERENCES difficulties_tasks (id)
			ON DELETE SET DEFAULT,
	type_id			INT 		DEFAULT NULL,
	FOREIGN KEY (type_id) 
		REFERENCES types_tasks (id)
			ON DELETE SET DEFAULT
);

CREATE TABLE team_tasks (
	line_num        SERIAL,
	id		text		NOT NULL,
	task_id			text 		NOT NULL,
	FOREIGN KEY (task_id) 
		REFERENCES tasks (id),
	team_id 		text 	    NOT NULL,
	FOREIGN KEY (team_id) 
		REFERENCES teams (id),
	start_time  	TIMESTAMPTZ DEFAULT NOW(),
	end_time    	TIMESTAMPTZ DEFAULT NULL,
	additional_points INT DEFAULT 0
);

CREATE TABLE types_answers (
	line_num	        SERIAL,
	id		INT		NOT NULL	PRIMARY KEY,
	title	text		NOT NULL
);

CREATE TABLE answers_tasks (
	line_num	        SERIAL,
	id		text		NOT NULL,
	task_id			text 		NOT NULL,
	FOREIGN KEY (task_id) 
		REFERENCES tasks (id),
	answer_type_id	INT 		NOT NULL,
	FOREIGN KEY (answer_type_id) 
		REFERENCES types_answers (id),
	data text NOT NULL
);

CREATE TABLE secrets (
	line_num	        SERIAL,
	id		text		NOT NULL	PRIMARY KEY,
	title	text NOT NULL,
	description	text
);

CREATE TABLE answers_secrets (
	line_num	        SERIAL,
	id		text		NOT NULL,
	secret_id			text 		NOT NULL,
	FOREIGN KEY (secret_id) 
		REFERENCES secrets (id),
	answer_type_id	INT 		NOT NULL,
	FOREIGN KEY (answer_type_id) 
		REFERENCES types_answers (id),
	data text NOT NULL
);

CREATE TABLE team_secrets (
	line_num	        SERIAL,
	id		text		NOT NULL,
	secret_id 		text 	    NOT NULL,
	FOREIGN KEY (secret_id) 
		REFERENCES secrets (id),
	team_id 		text 	    NOT NULL,
	FOREIGN KEY (team_id) 
		REFERENCES teams (id),
	start_time  	TIMESTAMPTZ DEFAULT NOW(),
	end_time    	TIMESTAMPTZ DEFAULT NULL
);
