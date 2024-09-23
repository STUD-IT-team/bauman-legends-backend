-- +goose Up
-- +goose StatementBegin
alter table master_class drop CONSTRAINT fk_media_obj_to_master_class;
delete from point_task;
delete from media_obj;
alter table media_obj alter column uuid_media type text;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table media_obj alter column uuid_media type uuid;
-- +goose StatementEnd
