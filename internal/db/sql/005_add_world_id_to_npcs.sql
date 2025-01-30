-- +goose Up
ALTER TABLE npcs
  ADD world_id INTEGER
  REFERENCES worlds(id);

-- +goose Down
ALTER TABLE npcs DROP COLUMN world_id;
