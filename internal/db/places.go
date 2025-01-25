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

/*
All search goes here. I'm still debating how to implement them.
Subject to change. Very grug brain implementation of it.
*/

type SearchParams struct {
	Name   string
	Limit  int
	Offset int
}

const searchWorldByNameQuery = `
SELECT * FROM worlds WHERE LOWER(name) LIKE LOWER($1)
	ORDER BY name ASC LIMIT $2 OFFSET $3
`

func (q *Queries) SearchWorldsByName(
	ctx context.Context,
	params SearchParams,
) ([]*World, error) {
	worlds := []*World{}
	rows, err := q.Db.QueryContext(
		ctx,
		searchWorldByNameQuery,
		params.Name,
		params.Limit,
		params.Offset,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		w := World{}
		err = rows.Scan(&w.Id, &w.Name, &w.Desc)
		if err != nil {
			return nil, err
		}
		worlds = append(worlds, &w)
	}

	return worlds, nil
}

const searchAreaByNameQuery = `
SELECT * FROM areas WHERE LOWER(name) LIKE LOWER($1)
	ORDER BY name ASC LIMIT $2 OFFSET $3
`

func (q *Queries) SearchAreasByName(
	ctx context.Context,
	params SearchParams,
) ([]*Area, error) {
	areas := []*Area{}
	rows, err := q.Db.QueryContext(
		ctx,
		searchAreaByNameQuery,
		params.Name,
		params.Limit,
		params.Offset,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		a := Area{}
		err = rows.Scan(&a.Id, &a.Name, &a.Type, &a.Desc, &a.WorldId)
		if err != nil {
			return nil, err
		}
		areas = append(areas, &a)
	}

	return areas, nil
}

func (q *Queries) SearchLocationsByName(
	ctx context.Context,
	params SearchParams,
) ([]*Location, error) {
	locations := []*Location{}
	rows, err := q.Db.QueryContext(
		ctx,
		searchAreaByNameQuery,
		params.Name,
		params.Limit,
		params.Offset,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		l := Location{}
		err = rows.Scan(&l.Id, &l.Name, &l.Type, &l.Desc, &l.AreaId)
		if err != nil {
			return nil, err
		}
		locations = append(locations, &l)
	}

	return locations, nil
}
