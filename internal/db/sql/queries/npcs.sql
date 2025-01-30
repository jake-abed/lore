-- name: CreateNPC :one
INSERT INTO npcs (
	name,
	race,
	class,
	subclass,
	alignment,
	sex,
	description,
	languages,
	level,
	hitpoints,
	world_id
) VALUES (?1,	?2,	?3,	?4,	?5,	?6,	?7,	?8,	?9,	?10,	?11)
RETURNING *;
