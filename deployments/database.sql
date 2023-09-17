CREATE TABLE teams (
	id	        SERIAL      	PRIMARY KEY,
	team_name	VARCHAR(30)	NOT NULL
);

CREATE TABLE roles (
	id		SERIAL      	PRIMARY KEY,
	role_name	VARCHAR(30)	NOT NULL
);

CREATE TABLE users (
	id			SERIAL 			PRIMARY KEY,
	phone_number		VARCHAR(20)		NOT NULL,
	email			VARCHAR(255)		NOT NULL,
	email_confirmed		BOOLEAN			DEFAULT FALSE,
	telegram		VARCHAR(32)		NOT NULL,
	vk			VARCHAR(32)		NOT NULL,
	study_group		text			NOT NULL,
	full_name		text			NOT NULL,
	team_id 		INT	 	    	DEFAULT NULL,
	FOREIGN KEY (team_id) 
		REFERENCES teams (id) 
			ON DELETE SET DEFAULT,
	role_id			INT 			DEFAULT NULL,
	FOREIGN KEY (role_id) 
		REFERENCES roles (id) 
			ON DELETE SET DEFAULT,
	admin_permissions	BOOLEAN			DEFAULT FALSE
);

CREATE TABLE sessions (
    	user_id			INT,
   	FOREIGN KEY (user_id) 
		REFERENCES users (id),
   	token			text,
   	end_session		date,
    	start_session       	date,
    	client_ip           	inet,
    	client_browser      	VARCHAR(32),
    	client_os           	VARCHAR(32),
    	geodata             	VARCHAR(20)
);

CREATE TABLE types_tasks (
	id	        SERIAL     	PRIMARY KEY,
	task_type	VARCHAR(20)	NOT NULL
);

CREATE TYPE type_difficulty AS ENUM ('1', '2', '3', '4', '5', '6', '7', '8', '9', '10');

CREATE TABLE tasks (
    	id	        SERIAL           PRIMARY KEY,
	task_name	VARCHAR(30)	 NOT NULL,
    	description	text,
    	correct_ans     text[]           DEFAULT array[]::text[],
    	qr_or_text_resp smallint         NOT NULL,
    	execution_time  VARCHAR(10)      NOT NULL,
    	task_difficulty type_difficulty  NOT NULL,
    	task_type_id    INT,
    	FOREIGN KEY (task_type_id) 
		REFERENCES types_tasks (id)
);

CREATE TABLE tasks_for_teams (
	id          	SERIAL      PRIMARY KEY,
   	task_id	    	INT,
    	FOREIGN KEY (task_id) 
		REFERENCES tasks (id),
   	 team_id     	INT,
	FOREIGN KEY (team_id) 
		REFERENCES teams (id),
    	start_time  	TIMESTAMPTZ DEFAULT NOW(),
    	end_time    	TIMESTAMPTZ DEFAULT NULL,
    	extra_point 	INT         DEFAULT 0
);

CREATE TABLE secrets (
  	id              SERIAL           PRIMARY KEY,
    	secret_name     VARCHAR(30)      NOT NULL,
    	description     text,
    	correct_ans     text[]           DEFAULT array[]::text[],
    	qr_or_text_resp smallint         NOT NULL
);

CREATE TABLE secrets_for_teams (
    	id          	SERIAL      PRIMARY KEY,
    	secret_id	INT,
    	FOREIGN KEY (secret_id) 
		REFERENCES secrets (id),
    	team_id     	INT,
	FOREIGN KEY (team_id) 
		REFERENCES teams (id),
    	start_time  	TIMESTAMPTZ DEFAULT NOW(),
    	end_time    	TIMESTAMPTZ DEFAULT NULL
)
