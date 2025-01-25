-- +goose Up
CREATE INDEX world_name_idx ON worlds (name);
CREATE INDEX area_name_idx ON areas (name);
CREATE INDEX loc_name_idx ON locations (name);
CREATE INDEX subloc_name_idx ON sublocations (name);

-- +goose Down
DROP INDEX world_name_idx ON worlds;
DROP INDEX area_name_idx ON areas;
DROP INDEX loc_name_idx ON locations;
DROP INDEX subloc_name_idx ON sublocations;
