package db

import (
	"context"
)

type Quest struct {
	Name        string
	Desc        string
	Rewards     string
	Notes       string
	Level       int
	IsStarted   bool
	IsFinished  bool
	CurrentStep int
	WorldId     int
	Id          int
}

type QuestParams struct {
	Name        string
	Desc        string
	Rewards     string
	Notes       string
	Level       int
	IsStarted   bool
	IsFinished  bool
	CurrentStep int
	WorldId     int
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
  current_step,
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
	$9
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
		p.CurrentStep,
		p.WorldId,
	)

	err := row.Scan(
		&quest.Id,
		&quest.Name,
		&quest.Desc,
		&quest.Rewards,
		&quest.Notes,
		&quest.Level,
		&quest.IsStarted,
		&quest.IsFinished,
		&quest.CurrentStep,
		&quest.WorldId,
	)
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

	err := row.Scan(
		&quest.Id,
		&quest.Name,
		&quest.Desc,
		&quest.Rewards,
		&quest.Notes,
		&quest.Level,
		&quest.IsStarted,
		&quest.IsFinished,
		&quest.CurrentStep,
		&quest.WorldId,
	)
	if err != nil {
		return nil, err
	}

	return &quest, nil
}

type QuestStep struct {
	Name       string
	Type       string
	Desc       string
	Reward     string
	IsStarted  bool
	IsFinished bool
	QuestId    int
	WorldId    int
}

type QuestStepParams struct {
	Name       string
	Type       string
	Desc       string
	Reward     string
	IsStarted  bool
	IsFinished bool
	QuestId    int
	WorldId    int
}
