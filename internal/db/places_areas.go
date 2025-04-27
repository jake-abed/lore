package db

import (
	"context"
	"database/sql"
)

type Area struct {
	Name    string
	Desc    string
	Type    string
	Id      int
	WorldId int
}

func (a *Area) PlaceType() PlaceType   { return AREA }
func (a *Area) Inspect() (int, string) { return a.Id, a.Name }

type AreaParams struct {
	Name    string
	Desc    string
	Type    string
	WorldId int
}

const createAreaQuery = `
INSERT INTO areas (name, description, type, world_id)
	VALUES ($1, $2, $3, $4)
	RETURNING *
`

func (q *Queries) AddArea(
	ctx context.Context,
	params *AreaParams,
) (*Area, error) {
	area := Area{}
	row := q.Db.QueryRowContext(
		ctx,
		createAreaQuery,
		params.Name,
		params.Desc,
		params.Type,
		params.WorldId,
	)

	err := row.Scan(
		&area.Id,
		&area.Name,
		&area.Type,
		&area.Desc,
		&area.WorldId,
	)
	if err != nil {
		return &Area{}, err
	}

	return &area, nil
}

const getAreaByNameQuery = `
SELECT * FROM areas WHERE LOWER(areas.name) LIKE LOWER($1) LIMIT 1
`

func (q *Queries) GetAreaByName(
	ctx context.Context,
	name string,
) (*Area, error) {
	area := Area{}
	row := q.Db.QueryRowContext(
		ctx,
		getAreaByNameQuery,
		name,
	)

	err := row.Scan(
		&area.Id,
		&area.Name,
		&area.Type,
		&area.Desc,
		&area.WorldId,
	)
	if err != nil {
		return nil, err
	}

	return &area, nil
}

const getAreaByIdQuery = `SELECT * FROM areas WHERE id = $1`

func (q *Queries) GetAreaById(
	ctx context.Context,
	id int,
) (*Area, error) {
	area := Area{}
	row := q.Db.QueryRowContext(
		ctx,
		getAreaByIdQuery,
		id,
	)

	err := row.Scan(
		&area.Id,
		&area.Name,
		&area.Type,
		&area.Desc,
		&area.WorldId,
	)
	if err != nil {
		return nil, err
	}

	return &area, nil
}

const getAllAreasQuery = `SELECT * FROM areas`

func (q *Queries) GetAllAreas(ctx context.Context) ([]*Area, error) {
	areas := []*Area{}

	rows, err := q.Db.QueryContext(ctx, getAllAreasQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		area := Area{}

		err := scanAreaRows(rows, &area)
		if err != nil {
			return nil, err
		}

		areas = append(areas, &area)
	}

	return areas, nil
}

const updateAreaByIdQuery = `UPDATE areas 
	SET name = $1, type = $2, description = $3, world_id = $4 WHERE id = $5
	RETURNING *
`

func (q *Queries) UpdateAreaById(
	ctx context.Context,
	area Area,
) (*Area, error) {
	updatedArea := Area{}
	row := q.Db.QueryRowContext(
		ctx,
		updateAreaByIdQuery,
		area.Name,
		area.Type,
		area.Desc,
		area.WorldId,
		area.Id,
	)
	err := row.Scan(
		&updatedArea.Id,
		&updatedArea.Name,
		&updatedArea.Type,
		&updatedArea.Desc,
		&updatedArea.WorldId,
	)
	if err != nil {
		return &Area{}, err
	}

	return &updatedArea, nil
}

const getXAreasQuery = `
SELECT * FROM areas WHERE areas.world_id LIKE $1
	ORDER BY areas.id ASC LIMIT $2 OFFSET $3
`

func (q *Queries) GetXAreas(
	ctx context.Context,
	worldId int,
	x int,
	offset int,
) ([]*Area, error) {
	areas := []*Area{}
	rows, err := q.Db.QueryContext(
		context.Background(),
		getXAreasQuery,
		worldId,
		x,
		offset*x,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		area := Area{}
		err := rows.Scan(
			&area.Id,
			&area.Name,
			&area.Type,
			&area.Desc,
			&area.WorldId,
		)
		if err != nil {
			return nil, err
		}

		areas = append(areas, &area)
	}

	return areas, nil
}

const deleteAreaByIdQuery = `DELETE FROM areas where id = $1`

func (q *Queries) DeleteAreaByIdQuery(ctx context.Context, id int) error {
	_, err := q.Db.ExecContext(ctx, deleteAreaByIdQuery, id)
	if err != nil {
		return err
	}

	return nil
}

func scanAreaRows(rows *sql.Rows, a *Area) error {
	return rows.Scan(
		&a.Id,
		&a.Name,
		&a.Type,
		&a.Desc,
		&a.WorldId,
	)
}
