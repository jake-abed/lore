-- +goose Up
CREATE TABLE worlds (
  id INTEGER UNIQUE PRIMARY KEY,
  name TEXT UNIQUE,
  description TEXT
);

CREATE TABLE areas (
  id INTEGER UNIQUE PRIMARY KEY,
  name TEXT,
  type TEXT,
  description TEXT,
  world_id INTEGER NOT NULL,
  FOREIGN KEY (world_id)
    REFERENCES worlds(id) ON DELETE CASCADE
);

CREATE TABLE locations (
  id INTEGER UNIQUE PRIMARY KEY,
  name TEXT,
  type TEXT,
  description TEXT,
  area_id INTEGER NOT NULL,
  FOREIGN KEY (area_id)
    REFERENCES areas(id) ON DELETE CASCADE
);

CREATE TABLE sublocations (
  id INTEGER UNIQUE PRIMARY KEY,
  name TEXT,
  type TEXT,
  description TEXT,
  location_id INTEGER NOT NULL,
  FOREIGN KEY (location_id)
    REFERENCES locations(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE worlds;
DROP TABLE areas;
DROP TABLE locations;
DROP TABLE sublocations;
