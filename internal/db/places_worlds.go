package db

import (
	"context"
)

type World struct {
	Name string
	Desc string
	Id   int
}

func (w *World) PlaceType() PlaceType   { return WORLD }
func (w *World) Inspect() (int, string) { return w.Id, w.Name }

type WorldParams struct {
	Name string
	Desc string
}

const worldCountQuery = `SELECT COUNT(id) FROM worlds`

func (q *Queries) WorldCount(ctx context.Context) (int, error) {
	var count int
	row := q.Db.QueryRowContext(ctx, worldCountQuery)
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
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

const getWorldByIdQuery = `SELECT * FROM worlds
WHERE worlds.id = $1 LIMIT 1`

func (q *Queries) GetWorldById(
	ctx context.Context,
	id int,
) (*World, error) {
	world := World{}
	row := q.Db.QueryRowContext(ctx, getWorldByIdQuery, id)
	err := row.Scan(&world.Id, &world.Name, &world.Desc)
	if err != nil {
		return &World{}, err
	}
	return &world, nil
}

const updateWorldByIdQuery = `UPDATE worlds
	SET name = $1, description = $2 WHERE id = $3
	RETURNING *
`

func (q *Queries) UpdateWorldById(
	ctx context.Context,
	world World,
) (*World, error) {
	updatedWorld := World{}
	row := q.Db.QueryRowContext(
		ctx,
		updateWorldByIdQuery,
		world.Name,
		world.Desc,
		world.Id,
	)
	err := row.Scan(&updatedWorld.Id, &updatedWorld.Name, &updatedWorld.Desc)
	if err != nil {
		return &World{}, err
	}

	return &updatedWorld, nil
}

const deleteWorldByIdQuery = `DELETE FROM worlds WHERE id = $1`

func (q *Queries) DeleteWorldByIdQuery(ctx context.Context, id int) error {
	_, err := q.Db.ExecContext(ctx, deleteWorldByIdQuery, id)
	if err != nil {
		return err
	}

	return nil
}

const getXWorldsQuery = `
SELECT * FROM worlds ORDER BY worlds.id ASC LIMIT $1 OFFSET $2
`

func (q *Queries) GetXWorlds(
	ctx context.Context,
	x int,
	offset int,
) ([]*World, error) {
	worlds := []*World{}
	rows, err := q.Db.QueryContext(
		context.Background(),
		getXWorldsQuery,
		x,
		offset*x,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		world := World{}
		err := rows.Scan(&world.Id, &world.Name, &world.Desc)
		if err != nil {
			return nil, err
		}

		worlds = append(worlds, &world)
	}

	return worlds, nil
}
