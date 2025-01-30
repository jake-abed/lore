package db

import (
	"context"
)

type Npc struct {
	Name        string
	Race        string
	Class       string
	Subclass    string
	Alignment   string
	Sex         string
	Description string
	Languages   string
	Id          int
	Level       int
	Hitpoints   int
	WorldId     int
}

type NpcParams struct {
	Name        string
	Race        string
	Class       string
	Subclass    string
	Alignment   string
	Sex         string
	Description string
	Languages   string
	Level       int
	Hitpoints   int
	WorldId     int
}

const createNpcQuery = `INSERT INTO npcs (
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
) VALUES (
	$1,
	$2,
	$3,
	$4,
	$5,
	$6,
	$7,
	$8,
	$9,
	$10,
	$11
) RETURNING *`

func (q *Queries) AddNpc(ctx context.Context, params *NpcParams) (*Npc, error) {
	npc := Npc{}
	row := q.Db.QueryRowContext(
		ctx,
		createNpcQuery,
		params.Name,
		params.Race,
		params.Class,
		params.Subclass,
		params.Alignment,
		params.Sex,
		params.Description,
		params.Languages,
		params.Level,
		params.Hitpoints,
		params.WorldId,
	)
	err := row.Scan(
		&npc.Id,
		&npc.Name,
		&npc.Race,
		&npc.Class,
		&npc.Subclass,
		&npc.Alignment,
		&npc.Level,
		&npc.Hitpoints,
		&npc.Sex,
		&npc.Description,
		&npc.Languages,
		&npc.WorldId,
	)

	if err != nil {
		return &Npc{}, err
	}
	return &npc, nil
}

const editNpcByIdQuery = `UPDATE npcs
SET name = $1,
		race = $2, 
		class = $3, 
		subclass = $4, 
		alignment = $5,
		level = $6,
		hitpoints = $7,
		sex = $8,
		description = $9,
		languages = $10,
		world_id = $11
WHERE id = $12
RETURNING *`

func (q *Queries) EditNpcById(
	ctx context.Context,
	npc *Npc,
) (*Npc, error) {
	row := q.Db.QueryRowContext(
		ctx,
		editNpcByIdQuery,
		npc.Name,
		npc.Race,
		npc.Class,
		npc.Subclass,
		npc.Alignment,
		npc.Level,
		npc.Hitpoints,
		npc.Sex,
		npc.Description,
		npc.Languages,
		npc.WorldId,
		npc.Id,
	)

	updated := Npc{}
	err := row.Scan(
		&updated.Id,
		&updated.Name,
		&updated.Race,
		&updated.Class,
		&updated.Subclass,
		&updated.Alignment,
		&updated.Level,
		&updated.Hitpoints,
		&updated.Sex,
		&updated.Description,
		&updated.Languages,
		&updated.WorldId,
	)
	if err != nil {
		return nil, err
	}

	return &updated, nil

}

const viewNpcQueryByName = `SELECT * FROM npcs WHERE name LIKE $1`

func (q *Queries) ViewNpcByName(
	ctx context.Context,
	name string,
) (*Npc, error) {
	npc := Npc{}
	row := q.Db.QueryRowContext(ctx, viewNpcQueryByName, name)
	err := row.Scan(
		&npc.Id,
		&npc.Name,
		&npc.Race,
		&npc.Class,
		&npc.Subclass,
		&npc.Alignment,
		&npc.Level,
		&npc.Hitpoints,
		&npc.Sex,
		&npc.Description,
		&npc.Languages,
		&npc.WorldId,
	)
	if err != nil {
		return nil, err
	}
	return &npc, nil
}

const searchNpcsByNameQuery = `SELECT * FROM npcs WHERE name LIKE $1`

func (q *Queries) SearchNpcsByName(
	ctx context.Context,
	name string,
) ([]*Npc, error) {
	rows, err := q.Db.QueryContext(ctx, searchNpcsByNameQuery, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var npcs []*Npc
	for rows.Next() {
		var new Npc
		err := rows.Scan(
			&new.Id,
			&new.Name,
			&new.Race,
			&new.Class,
			&new.Subclass,
			&new.Alignment,
			&new.Level,
			&new.Hitpoints,
			&new.Sex,
			&new.Description,
			&new.Languages,
			&new.WorldId,
		)
		if err != nil {
			return nil, err
		}
		npcs = append(npcs, &new)
	}
	return npcs, nil
}
