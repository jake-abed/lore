-- +goose Up
CREATE INDEX npc_name_idx ON npcs (name);

-- +goose Down
DROP INDEX npc_name_idx ON npcs;
