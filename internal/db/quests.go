package db

import (
	"context"
	"database/sql"
)

type Quest struct {
	Name       string
	Desc       string
	Rewards    string
	Notes      string
	Level      int
	IsStarted  bool
	IsFinished bool
	WorldId    int
	Id         int
}

type QuestParams struct {
	Name       string
	Desc       string
	Rewards    string
	Notes      string
	Level      int
	IsStarted  bool
	IsFinished bool
	WorldId    int
}

const createQuestQuery = `
INSERT INTO quests (
	name,
	description,
	rewards,
	notes,
	level,
	is_started,
  is_finished,
  world_id
) VALUES (
	$1,
	$2,
	$3,
	$4,
	$5,
	$6,
	$7,
	$8
) RETURNING *
`

func (q *Queries) AddQuest(
	ctx context.Context,
	p *QuestParams,
) (*Quest, error) {
	quest := Quest{}
	row := q.Db.QueryRowContext(
		ctx,
		createQuestQuery,
		p.Name,
		p.Desc,
		p.Rewards,
		p.Notes,
		p.Level,
		p.IsStarted,
		p.IsFinished,
		p.WorldId,
	)

	err := scanQuest(row, &quest)
	if err != nil {
		return nil, err
	}

	return &quest, nil
}

const getQuestById = `
SELECT * FROM quests WHERE id = $1 LIMIT 1
`

func (q *Queries) GetQuestByIdQuery(
	ctx context.Context,
	id int,
) (*Quest, error) {
	quest := Quest{}
	row := q.Db.QueryRowContext(ctx, getQuestById, id)

	err := scanQuest(row, &quest)
	if err != nil {
		return nil, err
	}

	return &quest, nil
}

type UpdateQuestParams struct {
	Name       string
	Desc       string
	Rewards    string
	Notes      string
	Level      int
	IsStarted  bool
	IsFinished bool
	WorldId    int
	Id         int
}

const updateQuestByIdQuery = `
UPDATE quests SET
	name = $1,
	description = $2,
	rewards = $3,
	notes = $4,
	level = $5,
	is_started = $6,
	is_finished = $7,
	world_id = $8
	WHERE id = $9
	RETURNING *
`

func (q *Queries) UpdateQuestById(
	ctx context.Context,
	p UpdateQuestParams,
) (*Quest, error) {
	quest := Quest{}

	row := q.Db.QueryRowContext(
		ctx,
		updateQuestByIdQuery,
		p.Name,
		p.Desc,
		p.Rewards,
		p.Notes,
		p.Level,
		p.IsStarted,
		p.IsFinished,
		p.WorldId,
		p.Id,
	)

	err := scanQuest(row, &quest)
	if err != nil {
		return nil, err
	}

	return &quest, nil
}

const getXQuestsQuery = `
SELECT * FROM quests ORDER BY id ASC LIMIT $1 OFFSET $2
`

func (q *Queries) GetXQuests(
	ctx context.Context,
	x int,
	offset int,
) ([]*Quest, error) {
	quests := []*Quest{}

	rows, err := q.Db.QueryContext(ctx, getXQuestsQuery, x, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		quest := Quest{}

		err := scanQuestRows(rows, &quest)
		if err != nil {
			return nil, err
		}

		quests = append(quests, &quest)
	}

	return quests, nil
}

const getQuestsByNameQuery = `
SELECT * FROM quests WHERE LOWER(name) LIKE LOWER($1)
	ORDER BY name ASC
`

func (q *Queries) GetQuestsByName(
	ctx context.Context,
	name string,
) ([]*Quest, error) {
	quests := []*Quest{}

	rows, err := q.Db.QueryContext(ctx, getQuestsByNameQuery, name)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		quest := Quest{}

		err := scanQuestRows(rows, &quest)
		if err != nil {
			return nil, err
		}

		quests = append(quests, &quest)
	}

	return quests, nil
}

const deleteQuestByIdQuery = `
DELETE FROM quests WHERE id = $1
`

func (q *Queries) DeleteQuestById(ctx context.Context, id int) error {
	_, err := q.Db.ExecContext(ctx, deleteQuestByIdQuery, id)
	if err != nil {
		return err
	}

	return nil
}

// Scan Helpers

func scanQuest(row *sql.Row, q *Quest) error {
	err := row.Scan(
		&q.Id,
		&q.Name,
		&q.Desc,
		&q.Rewards,
		&q.Notes,
		&q.Level,
		&q.IsStarted,
		&q.IsFinished,
		&q.WorldId,
	)
	if err != nil {
		return err
	}

	return nil
}

func scanQuestRows(rows *sql.Rows, q *Quest) error {
	err := rows.Scan(
		&q.Id,
		&q.Name,
		&q.Desc,
		&q.Rewards,
		&q.Notes,
		&q.Level,
		&q.IsStarted,
		&q.IsFinished,
		&q.WorldId,
	)
	if err != nil {
		return err
	}

	return nil
}
