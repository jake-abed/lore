-- +goose Up
CREATE TABLE npcs (
  id INTEGER PRIMARY KEY UNIQUE,
  name TEXT,
  race TEXT,
	class TEXT,
	subclass TEXT,
  alignment TEXT,
  level INTEGER,
	hitpoints INTEGER,
	sex TEXT,
	description TEXT,
	languages TEXT
);

-- +goose Down
DROP TABLE npcs; 
