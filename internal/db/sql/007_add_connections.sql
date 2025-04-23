-- +goose Up
CREATE TABLE npcs_quests (
  npc_id INTEGER NOT NULL,
  quest_id INTEGER NOT NULL,
  FOREIGN KEY (npc_id) REFERENCES npcs(id) ON DELETE CASCADE,
  FOREIGN KEY (quest_id) REFERENCES quests(id) ON DELETE CASCADE,
  UNIQUE(npc_id, quest_id)
);

CREATE TABLE npcs_worlds (
  npc_id INTEGER NOT NULL,
  world_id INTEGER NOT NULL,
  FOREIGN KEY (npc_id) REFERENCES npcs(id) ON DELETE CASCADE,
  FOREIGN KEY (world_id) REFERENCES worlds(id) ON DELETE CASCADE,
  UNIQUE(npc_id, world_id)
);

CREATE TABLE npcs_areas (
  npc_id INTEGER NOT NULL,
  area_id INTEGER NOT NULL,
  FOREIGN KEY (npc_id) REFERENCES npcs(id) ON DELETE CASCADE,
  FOREIGN KEY (area_id) REFERENCES areas(id) ON DELETE CASCADE,
  UNIQUE(npc_id, area_id)
);

CREATE TABLE npcs_locations (
  npc_id INTEGER NOT NULL,
  location_id INTEGER NOT NULL,
  FOREIGN KEY (npc_id) REFERENCES npcs(id) ON DELETE CASCADE,
  FOREIGN KEY (location_id) REFERENCES locations(id) ON DELETE CASCADE,
  UNIQUE(npc_id, location_id)
);

CREATE TABLE quests_worlds (
  quest_id INTEGER NOT NULL,
  world_id INTEGER NOT NULL,
  FOREIGN KEY (quest_id) REFERENCES quests(id) ON DELETE CASCADE,
  FOREIGN KEY (world_id) REFERENCES worlds(id) ON DELETE CASCADE,
  UNIQUE(quest_id, world_id)
);

CREATE TABLE quests_areas (
  quest_id INTEGER NOT NULL,
  area_id INTEGER NOT NULL,
  FOREIGN KEY (quest_id) REFERENCES quests(id) ON DELETE CASCADE,
  FOREIGN KEY (area_id) REFERENCES areas(id) ON DELETE CASCADE,
  UNIQUE(quest_id, area_id)
);

CREATE TABLE quests_locations (
  quest_id INTEGER NOT NULL,
  location_id INTEGER NOT NULL,
  FOREIGN KEY (quest_id) REFERENCES quests(id) ON DELETE CASCADE,
  FOREIGN KEY (location_id) REFERENCES locations(id) ON DELETE CASCADE,
  UNIQUE(quest_id, location_id)
);


-- +goose Down
DROP TABLE npcs_quests;
DROP TABLE npcs_worlds;
DROP TABLE npcs_areas;
DROP TABLE npcs_locations;
DROP TABLE quests_worlds;
DROP TABLE quests_areas;
DROP TABLE quests_locations;
