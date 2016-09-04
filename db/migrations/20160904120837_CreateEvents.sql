-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table events (
  id serial primary key,
  name text not null,
  status text,
  payload jsonb,
  created_at timestamptz not null,
  updated_at timestamptz not null,
  deleted_at timestamptz
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table events;
