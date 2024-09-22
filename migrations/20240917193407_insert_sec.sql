-- +goose Up
-- +goose StatementBegin
ALTER TABLE master_class
    ADD ended_at TIMESTAMP;

ALTER TABLE master_class
    DROP duration;


INSERT into "user" (name, "group", email, telegram, vk, password, phone_number)
VALUES ('Монгилёв Андрей Валентинович', 'sec', '@mongilyov', '@mongilyov', '', '', '79895021950') RETURNING id;

INSERT into "user" (name, "group", email, telegram, vk, password, phone_number)
VALUES ('Чернышев Владимир Юрьевич', 'sec', '@SirLambor', '@SirLambor', '', '', '79287645771') RETURNING id;

INSERT into "user" (name, "group", email, telegram, vk, password, phone_number)
VALUES ('Токарев Иван Дмитриевич', 'sec', '@ivantokarev', '@ivantokarev', '', '', '79026532930') RETURNING id;

INSERT into "user" (name, "group", email, telegram, vk, password, phone_number)
VALUES ('Короткова Ульяна Андреевна', 'sec', '@uly_lapula', '@uly_lapula', '', '', '79046869628') RETURNING id;

INSERT into "user" (name, "group", email, telegram, vk, password, phone_number)
VALUES ('Сухова Юлия Дмитриевна', 'sec', '@ssowhattt', '@ssowhattt', '', '', '79251446548') RETURNING id;

INSERT into "user" (name, "group", email, telegram, vk, password, phone_number)
VALUES ('Балакало Максим Сергеевич', 'sec', '@maks_balakalo', '@maks_balakalo', '', '', '79065689673') RETURNING id;

INSERT INTO sec (name, description, responsible_id)
VALUES ('ISCRA', '', (SELECT id FROM "user" WHERE email = '@mongilyov'));

INSERT INTO sec (name, description, responsible_id)
VALUES ('BRT', '', (SELECT id FROM "user" WHERE email = '@SirLambor'));

INSERT INTO sec (name, description, responsible_id)
VALUES ('ITS BMSTU', '', (SELECT id FROM "user" WHERE email = '@ivantokarev'));

INSERT INTO sec (name, description, responsible_id)
VALUES ('МКЦ', '', (SELECT id FROM "user" WHERE email = '@uly_lapula'));

INSERT INTO sec (name, description, responsible_id)
VALUES ('Гидронавтика', '', (SELECT id FROM "user" WHERE email = '@ssowhattt'));

INSERT INTO sec (name, description, responsible_id)
VALUES ('ЦМР', '', (SELECT id FROM "user" WHERE email = '@maks_balakalo'));

INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 16:00:00', '2024-09-24 16:30:00', 18, (SELECT id FROM sec WHERE name = 'ISCRA'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 16:30:00', '2024-09-24 17:00:00', 18, (SELECT id FROM sec WHERE name = 'ISCRA'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 17:00:00', '2024-09-24 17:30:00', 18, (SELECT id FROM sec WHERE name = 'ISCRA'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 17:30:00', '2024-09-24 18:00:00', 18, (SELECT id FROM sec WHERE name = 'ISCRA'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 18:00:00', '2024-09-24 18:30:00', 18, (SELECT id FROM sec WHERE name = 'ISCRA'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 18:30:00', '2024-09-24 19:00:00', 18, (SELECT id FROM sec WHERE name = 'ISCRA'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 19:00:00', '2024-09-24 19:30:00', 18, (SELECT id FROM sec WHERE name = 'ISCRA'), 1);

INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 14:00', '2024-09-24 14:20', 8, (SELECT id FROM sec WHERE name = 'BRT'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 14:20', '2024-09-24 14:40', 8, (SELECT id FROM sec WHERE name = 'BRT'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 14:40', '2024-09-24 15:00', 8, (SELECT id FROM sec WHERE name = 'BRT'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 15:00', '2024-09-24 15:20', 8, (SELECT id FROM sec WHERE name = 'BRT'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 15:20', '2024-09-24 15:40', 8, (SELECT id FROM sec WHERE name = 'BRT'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 15:40', '2024-09-24 16:00', 8, (SELECT id FROM sec WHERE name = 'BRT'), 1);

INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 16:00', '2024-09-24 16:45', 50, (SELECT id FROM sec WHERE name = 'ITS BMSTU'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 16:45', '2024-09-24 17:30', 50, (SELECT id FROM sec WHERE name = 'ITS BMSTU'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 17:30', '2024-09-24 18:15', 50, (SELECT id FROM sec WHERE name = 'ITS BMSTU'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 18:15', '2024-09-24 19:00', 50, (SELECT id FROM sec WHERE name = 'ITS BMSTU'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 19:00', '2024-09-24 19:45', 50, (SELECT id FROM sec WHERE name = 'ITS BMSTU'), 1);

INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 12:00', '2024-09-24 13:00', 30, (SELECT id FROM sec WHERE name = 'МКЦ'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 13:00', '2024-09-24 14:00', 30, (SELECT id FROM sec WHERE name = 'МКЦ'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 14:00', '2024-09-24 15:00', 30, (SELECT id FROM sec WHERE name = 'МКЦ'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 15:00', '2024-09-24 16:00', 30, (SELECT id FROM sec WHERE name = 'МКЦ'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 16:00', '2024-09-24 17:00', 30, (SELECT id FROM sec WHERE name = 'МКЦ'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 17:00', '2024-09-24 18:00', 30, (SELECT id FROM sec WHERE name = 'МКЦ'), 1);

INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 16:00', '2024-09-24 16:30', 8, (SELECT id FROM sec WHERE name = 'Гидронавтика'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 16:30', '2024-09-24 17:00', 8, (SELECT id FROM sec WHERE name = 'Гидронавтика'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 17:00', '2024-09-24 17:30', 8, (SELECT id FROM sec WHERE name = 'Гидронавтика'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 17:30', '2024-09-24 18:00', 8, (SELECT id FROM sec WHERE name = 'Гидронавтика'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 18:00', '2024-09-24 18:30', 8, (SELECT id FROM sec WHERE name = 'Гидронавтика'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 18:30', '2024-09-24 19:00', 8, (SELECT id FROM sec WHERE name = 'Гидронавтика'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 19:00', '2024-09-24 19:30', 8, (SELECT id FROM sec WHERE name = 'Гидронавтика'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 19:30', '2024-09-24 20:00', 8, (SELECT id FROM sec WHERE name = 'Гидронавтика'), 1);

INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 16:00', '2024-09-24 16:25', 10, (SELECT id FROM sec WHERE name = 'ЦМР'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 16:25', '2024-09-24 16:50', 10, (SELECT id FROM sec WHERE name = 'ЦМР'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 16:50','2024-09-24 17:15', 10, (SELECT id FROM sec WHERE name = 'ЦМР'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 17:15', '2024-09-24 17:40', 10, (SELECT id FROM sec WHERE name = 'ЦМР'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 17:40', '2024-09-24 18:05', 10, (SELECT id FROM sec WHERE name = 'ЦМР'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 18:05','2024-09-24 18:30', 10, (SELECT id FROM sec WHERE name = 'ЦМР'), 1);
INSERT INTO master_class (started_at, ended_at, capacity, sec_id, media_id)
VALUES ('2024-09-24 18:30', '2024-09-24 18:55', 10, (SELECT id FROM sec WHERE name = 'ЦМР'), 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE master_class
    DROP ended_at;

ALTER TABLE master_class
    ADD duration INTEGER;

TRUNCATE master_class;

DELETE FROM "user" WHERE email in ('@mongilyov', '@SirLambor','@ivantokarev',  '@uly_lapula', '@ssowhattt', '@maks_balakalo');

-- +goose StatementEnd
