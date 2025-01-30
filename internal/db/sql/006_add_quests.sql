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
  current_step INTEGER NOT NULL DEFAULT 0,
  world_id INTEGER NOT NULL,
  FOREIGN KEY (world_id)
    REFERENCES world(id) ON DELETE CASCADE
);

CREATE TABLE quest_steps (
  id INTEGER UNIQUE PRIMARY KEY,
  name TEXT,
  type TEXT,
  description TEXT,
  reward TEXT,
  is_started BOOLEAN NOT NULL DEFAULT FALSE,
  is_finished BOOLEAN NOT NULL DEFAULT FALSE,
  quest_id INTEGER NOT NULL,
  FOREIGN KEY (quest_id)
    REFERENCES quests(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE quests;
DROP TABLE quest_steps;
