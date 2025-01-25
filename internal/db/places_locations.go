package db

import (
	"context"
)

type Location struct {
	Name   string
	Type   string
	Desc   string
	Id     int
	AreaId int
}

func (l *Location) PlaceType() PlaceType   { return LOCATION }
func (l *Location) Inspect() (int, string) { return l.Id, l.Name }

type LocationParams struct {
	Name   string
	Type   string
	Desc   string
	AreaId int
}

const createLocationQuery = `
INSERT INTO locations (name, description, type, area_id)
	VALUES ($1, $2, $3, $4)
	RETURNING *
`

func (q *Queries) AddLocation(
	ctx context.Context,
	params *LocationParams,
) (*Location, error) {
	location := Location{}
	row := q.Db.QueryRowContext(
		ctx,
		createLocationQuery,
		params.Name,
		params.Desc,
		params.Type,
		params.AreaId,
	)

	err := row.Scan(
		&location.Id,
		&location.Name,
		&location.Type,
		&location.Desc,
		&location.AreaId,
	)
	if err != nil {
		return &Location{}, err
	}

	return &location, nil
}

const getLocationByNameQuery = `
SELECT * FROM locations WHERE LOWER(locations.name) LIKE LOWER($1) LIMIT 1
`

func (q *Queries) GetLocationByName(
	ctx context.Context,
	name string,
) (*Location, error) {
	location := Location{}
	row := q.Db.QueryRowContext(
		ctx,
		getLocationByNameQuery,
		name,
	)

	err := row.Scan(
		&location.Id,
		&location.Name,
		&location.Type,
		&location.Desc,
		&location.AreaId,
	)
	if err != nil {
		return nil, err
	}

	return &location, nil
}

const updateLocationByIdQuery = `UPDATE locations
	SET name = $1, type = $2, description = $3, area_id = $4 WHERE id = $5
	RETURNING *
`

func (q *Queries) UpdateLocationById(
	ctx context.Context,
	location Location,
) (*Location, error) {
	updatedLocation := Location{}
	row := q.Db.QueryRowContext(
		ctx,
		updateLocationByIdQuery,
		location.Name,
		location.Type,
		location.Desc,
		location.AreaId,
		location.Id,
	)
	err := row.Scan(
		&updatedLocation.Id,
		&updatedLocation.Name,
		&updatedLocation.Type,
		&updatedLocation.Desc,
		&updatedLocation.AreaId,
	)
	if err != nil {
		return &Location{}, err
	}

	return &updatedLocation, nil
}

const getXLocationsQuery = `
SELECT * FROM locations ORDER BY locations.id ASC LIMIT $1 OFFSET $2
`

func (q *Queries) GetXLocations(
	ctx context.Context,
	x int,
	offset int,
) ([]*Location, error) {
	locations := []*Location{}
	rows, err := q.Db.QueryContext(
		context.Background(),
		getXLocationsQuery,
		x,
		offset*x,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		location := Location{}
		err := rows.Scan(&location.Id, &location.Name, &location.Desc)
		if err != nil {
			return nil, err
		}

		locations = append(locations, &location)
	}

	return locations, nil
}

const deleteLocationByIdQuery = `DELETE FROM locations WHERE id = $1`

func (q *Queries) DeleteLocationByIdQuery(ctx context.Context, id int) error {
	_, err := q.Db.ExecContext(ctx, deleteLocationByIdQuery, id)
	if err != nil {
		return err
	}

	return nil
}
