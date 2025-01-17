package db

import (
	"context"
)

type PlaceType string

const (
	WORLD       = "World"
	AREA        = "Area"
	LOCATION    = "Location"
	SUBLOCATION = "Sublocation"
)

type Place interface {
	PlaceType() PlaceType
	Inspect() (int, string)
}

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
		err := rows.Scan(&area.Id, &area.Name, &area.Desc)
		if err != nil {
			return nil, err
		}

		areas = append(areas, &area)
	}

	return areas, nil
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
