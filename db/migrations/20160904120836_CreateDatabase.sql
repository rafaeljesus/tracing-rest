-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create database tracing_rest_dev;
create database tracing_rest_test;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop database database tracing_rest_dev;
drop database tracing_rest_test;
