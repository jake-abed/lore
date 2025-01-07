package db

import (
	"context"
)

type PlaceType string

const (
	WORLD = "WORLD"
	REGION = "REGION"
	LOCATION = "LOCATION"
)

type Place interface {
	Type() PlaceType
	Inspect() (int, string)
}

type World struct {
	Name string
	Desc string
	Id   int
}

func (w *World) Type() PlaceType { return WORLD }
func (w *World) Inspect() (int, string) { return w.Id, w.Name }

type WorldParams struct {
	Name string
	Desc string
}

const createWorldQuery = `INSERT INTO worlds (name, description)
VALUES ($1, $2) RETURNING *`

func (q *Queries) AddWorld(
	ctx context.Context,
	params *WorldParams,
) (*World, error) {
	world := World{}
	row := q.Db.QueryRowContext(ctx, createWorldQuery, params.Name, params.Desc)
	err := row.Scan(&world.Id, &world.Name, &world.Desc)
	if err != nil {
		return &World{}, err
	}
	return &world, nil
}

const getWorldQuery = `SELECT * FROM worlds
WHERE LOWER(worlds.name) = $1 LIMIT 1`

func (q *Queries) GetWorldByName(
	ctx context.Context,
	name string,
) (*World, error) {
	world := World{}
	row := q.Db.QueryRowContext(ctx, getWorldQuery, name)
	err := row.Scan(&world.Id, &world.Name, &world.Desc)
	if err != nil {
		return &World{}, err
	}
	return &world, nil
}
