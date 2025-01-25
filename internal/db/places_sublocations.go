package db

import (
	"context"
)

type Sublocation struct {
	Name       string
	Type       string
	Desc       string
	Id         int
	LocationId int
}

func (s *Sublocation) PlaceType() PlaceType   { return SUBLOCATION }
func (s *Sublocation) Inspect() (int, string) { return s.Id, s.Name }

type SublocationParams struct {
	Name       string
	Type       string
	Desc       string
	LocationId string
}

const createSublocationQuery = `
INSERT INTO sublocations (name, description, type, location_id)
	VALUES ($1, $2, $3, $4)
	RETURNING *
`

func (q *Queries) AddSublocation(
	ctx context.Context,
	params *SublocationParams,
) (*Sublocation, error) {
	sublocation := Sublocation{}
	row := q.Db.QueryRowContext(
		ctx,
		createSublocationQuery,
		params.Name,
		params.Desc,
		params.Type,
		params.LocationId,
	)

	err := row.Scan(
		&sublocation.Id,
		&sublocation.Name,
		&sublocation.Type,
		&sublocation.Desc,
		&sublocation.LocationId,
	)
	if err != nil {
		return &Sublocation{}, err
	}

	return &sublocation, nil
}

const getSublocationByNameQuery = `
SELECT * FROM locations WHERE LOWER(locations.name) LIKE LOWER($1) LIMIT 1
`

func (q *Queries) GetSublocationByName(
	ctx context.Context,
	name string,
) (*Sublocation, error) {
	sublocation := Sublocation{}
	row := q.Db.QueryRowContext(
		ctx,
		getSublocationByNameQuery,
		name,
	)

	err := row.Scan(
		&sublocation.Id,
		&sublocation.Name,
		&sublocation.Type,
		&sublocation.Desc,
		&sublocation.LocationId,
	)
	if err != nil {
		return nil, err
	}

	return &sublocation, nil
}
