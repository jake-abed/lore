package db

import (
	"context"
	"fmt"
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
	hitpoints
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
	$10
) RETURNING *`

func (q *Queries) AddNpc(ctx context.Context, params *NpcParams) (*Npc, error) {
	npc := Npc{}
	row := q.Db.QueryRowContext(ctx, createNpcQuery, params.Name, params.Race,
		params.Class, params.Subclass, params.Alignment, params.Sex,
		params.Description, params.Languages, params.Level, params.Hitpoints)
	err := row.Scan(&npc.Id, &npc.Name, &npc.Race, &npc.Class, &npc.Subclass,
		&npc.Alignment, &npc.Level, &npc.Hitpoints, &npc.Sex, &npc.Description,
		&npc.Languages)
	if err != nil {
		fmt.Println(err)
		return &Npc{}, err
	}
	return &npc, nil
}
