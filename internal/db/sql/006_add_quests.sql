-- +goose Up
CREATE TABLE quests (
  id INTEGER UNIQUE PRIMARY KEY,
  name TEXT,
  description TEXT,
  rewards TEXT,
  notes TEXT,
  level INTEGER,
  is_started BOOLEAN NOT NULL DEFAULT FALSE,
  is_finished BOOLEAN NOT NULL DEFAULT FALSE,
  world_id INTEGER NOT NULL,
  FOREIGN KEY (world_id)
    REFERENCES world(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE quests;
