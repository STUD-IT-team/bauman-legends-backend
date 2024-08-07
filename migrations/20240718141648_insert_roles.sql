-- +goose Up
-- +goose StatementBegin
INSERT INTO "role" (type)
VALUES ('Участник'),
       ('Продавец'),
       ('Капитан'),
       ('Админ');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE ROLE;
-- +goose StatementEnd
