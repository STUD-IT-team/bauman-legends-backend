-- +goose Up
-- +goose StatementBegin
CREATE TABLE master_class
(
    id         SERIAL   NOT NULL,
    started_at TIMESTAMP NULL    ,
    duration   INTEGER     NULL     DEFAULT 0,
    capacity   INTEGER  NULL     DEFAULT 0, --COMMENT 'количество тел участников',
    sec_id     INTEGER   NOT NULL,
    media_id   INTEGER   NOT NULL,  --COMMENT 'для навигации до точки',
    PRIMARY KEY (id)
);

CREATE TABLE media_obj
(
    id         SERIAL NOT NULL,
    uuid_media UUID   NOT NULL,
    type       TEXT   NULL    ,
    PRIMARY KEY (id)
);

CREATE TABLE point_task
(
    id          SERIAL  NOT NULL,
    description TEXT    NULL    ,
    media_id    INTEGER  NOT NULL,
    points  INTEGER NULL     DEFAULT 0,
    PRIMARY KEY (id)
);

CREATE TABLE role
(
    id   SERIAL NOT NULL,
    type TEXT   NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE sec
(
    id             SERIAL NOT NULL,
    name           TEXT   NULL    ,
    description    TEXT   NULL    ,
    responsible_id INTEGER NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE team
(
    id           SERIAL  NOT NULL,
    name         TEXT    NULL    ,
    delta_points INTEGER NULL    DEFAULT 0,
    final_video  BOOLEAN DEFAULT FALSE,
    PRIMARY KEY (id)
);

CREATE TABLE team_master_class
(
    id              SERIAL NOT NULL,
    team_id         INTEGER NOT NULL,
    master_class_id INTEGER NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE team_media_answer
(
    id            SERIAL   NOT NULL,
    team_id       INTEGER   NOT NULL,
    point_task_id INTEGER   NOT NULL,
    points        INTEGER  NULL     DEFAULT 0,
    comment       TEXT     NULL, ---     COMMENT 'коммент от проверяющего',
    media_id      SERIAL   NOT NULL,
    status        TEXT     NULL    ,
    date          TIMESTAMP NULL    ,
    PRIMARY KEY (id)
);

CREATE TABLE team_text_answer
(
    id           SERIAL   NOT NULL,
    team_id      INTEGER   NOT NULL,
    answer       TEXT     NOT NULL,
    text_task_id SERIAL   NOT NULL,
    points       INTEGER  NULL     DEFAULT 0,
    status       TEXT     NULL    ,
    date        TIMESTAMP NULL    ,
    PRIMARY KEY (id)
);

CREATE TABLE text_task
(
    id          SERIAL  NOT NULL,
    description TEXT    NOT NULL,
    answer      TEXT    NOT NULL,
    points  INTEGER NULL     DEFAULT 0,
    PRIMARY KEY (id)
);

CREATE TABLE "user"
(
    id           SERIAL NOT NULL,
    role_id      INTEGER NOT NULL DEFAULT 1,
    name         TEXT   NOT NULL,
    "group"      TEXT   NOT NULL,
    email        TEXT   UNIQUE NOT NULL, -- COMMENT ' unique',
    telegram     TEXT   NOT NULL,
    vk           TEXT   NOT NULL,
    password     TEXT   NOT NULL,
    phone_number TEXT   NOT NULL,
    team_id      INTEGER NULL    ,
    PRIMARY KEY (id)
);

ALTER TABLE "user"
    ADD CONSTRAINT FK_role_TO_user
        FOREIGN KEY (role_id)
            REFERENCES role (id);

ALTER TABLE point_task
    ADD CONSTRAINT FK_media_obj_TO_point_task
        FOREIGN KEY (media_id)
            REFERENCES media_obj (id);

ALTER TABLE team_text_answer
    ADD CONSTRAINT FK_team_TO_team_text_answer
        FOREIGN KEY (team_id)
            REFERENCES team (id);

ALTER TABLE team_text_answer
    ADD CONSTRAINT FK_text_task_TO_team_text_answer
        FOREIGN KEY (text_task_id)
            REFERENCES text_task (id);

ALTER TABLE team_media_answer
    ADD CONSTRAINT FK_team_TO_team_media_answer
        FOREIGN KEY (team_id)
            REFERENCES team (id);

ALTER TABLE team_media_answer
    ADD CONSTRAINT FK_point_task_TO_team_media_answer
        FOREIGN KEY (point_task_id)
            REFERENCES point_task (id);

ALTER TABLE "user"
    ADD CONSTRAINT FK_team_TO_user
        FOREIGN KEY (team_id)
            REFERENCES team (id);

ALTER TABLE team_media_answer
    ADD CONSTRAINT FK_media_obj_TO_team_media_answer
        FOREIGN KEY (media_id)
            REFERENCES media_obj (id);

ALTER TABLE master_class
    ADD CONSTRAINT FK_sec_TO_master_class
        FOREIGN KEY (sec_id)
            REFERENCES sec (id);

ALTER TABLE team_master_class
    ADD CONSTRAINT FK_team_TO_team_master_class
        FOREIGN KEY (team_id)
            REFERENCES team (id);

ALTER TABLE team_master_class
    ADD CONSTRAINT FK_master_class_TO_team_master_class
        FOREIGN KEY (master_class_id)
            REFERENCES master_class (id);

ALTER TABLE master_class
    ADD CONSTRAINT FK_media_obj_TO_master_class
        FOREIGN KEY (media_id)
            REFERENCES media_obj (id);

ALTER TABLE sec
    ADD CONSTRAINT FK_user_TO_sec
        FOREIGN KEY (responsible_id)
            REFERENCES "user" (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS role;
DROP TABLE IF EXISTS sec;
DROP TABLE IF EXISTS team_master_class;
DROP TABLE IF EXISTS team_media_answer;
DROP TABLE IF EXISTS point_task;
DROP TABLE IF EXISTS media_obj;
DROP TABLE IF EXISTS team_text_answer;
DROP TABLE IF EXISTS team;
DROP TABLE IF EXISTS text_task;
DROP TABLE IF EXISTS master_class;
-- +goose StatementEnd
