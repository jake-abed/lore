package db

import (
	"context"
)

type NpcQuest struct {
	NpcId   int
	QuestId int
}

type NpcWorld struct {
	NpcId   int
	WorldId int
}

type NpcArea struct {
	NpcId  int
	AreaId int
}

type NpcLocation struct {
	NpcId      int
	LocationId int
}

type QuestWorld struct {
	QuestId int
	WorldId int
}

type QuestArea struct {
	QuestId int
	AreaId  int
}

type QuestLocation struct {
	QuestId    int
	LocationId int
}

type ConnectionParams struct {
	FirstId  int
	SecondId int
}

const createNpcQuestConnectionQuery = `
INSERT INTO npcs_quests (npc_id, quest_id)
	VALUES ($1, $2)
	RETURNING npc_id, quest_id
`

func (q *Queries) CreateNpcQuestConnection(
	ctx context.Context,
	params ConnectionParams,
) (*NpcQuest, error) {
	npcQuest := NpcQuest{}
	row := q.Db.QueryRowContext(
		context.Background(),
		createNpcQuestConnectionQuery,
		params.FirstId,
		params.SecondId,
	)

	err := row.Scan(&npcQuest.NpcId, &npcQuest.QuestId)
	if err != nil {
		return nil, err
	}

	return &npcQuest, nil
}

const createNpcWorldConnectionQuery = `
INSERT INTO npcs_worlds (npc_id, world_id)
	VALUES ($1, $2)
	RETURNING npc_id, world_id
`

func (q *Queries) CreateNpcWorldConnection(
	ctx context.Context,
	params ConnectionParams,
) (*NpcWorld, error) {
	npcWorld := NpcWorld{}
	row := q.Db.QueryRowContext(
		context.Background(),
		createNpcWorldConnectionQuery,
		params.FirstId,
		params.SecondId,
	)

	err := row.Scan(&npcWorld.NpcId, &npcWorld.WorldId)
	if err != nil {
		return nil, err
	}

	return &npcWorld, nil
}

const createNpcAreaConnectionQuery = `
INSERT INTO npcs_areas(npc_id, area_id)
	VALUES ($1, $2)
	RETURNING npc_id, area_id 
`

func (q *Queries) CreateNpcAreaConnection(
	ctx context.Context,
	params ConnectionParams,
) (*NpcArea, error) {
	npcArea := NpcArea{}
	row := q.Db.QueryRowContext(
		context.Background(),
		createNpcAreaConnectionQuery,
		params.FirstId,
		params.SecondId,
	)

	err := row.Scan(&npcArea.NpcId, &npcArea.AreaId)
	if err != nil {
		return nil, err
	}

	return &npcArea, nil
}

const createNpcLocationConnectionQuery = `
INSERT INTO npcs_locations(npc_id, location_id)
	VALUES ($1, $2)
	RETURNING npc_id, location_id 
`

func (q *Queries) CreateNpcLocationConnection(
	ctx context.Context,
	params ConnectionParams,
) (*NpcLocation, error) {
	npcLocation := NpcLocation{}
	row := q.Db.QueryRowContext(
		context.Background(),
		createNpcLocationConnectionQuery,
		params.FirstId,
		params.SecondId,
	)

	err := row.Scan(&npcLocation.NpcId, &npcLocation.LocationId)
	if err != nil {
		return nil, err
	}

	return &npcLocation, nil
}

const createQuestWorldConnectionQuery = `
INSERT INTO quests_worlds(quest_id, world_id)
	VALUES ($1, $2)
	RETURNING quest_id, world_id 
`

func (q *Queries) CreateQuestWorldConnection(
	ctx context.Context,
	params ConnectionParams,
) (*QuestWorld, error) {
	questWorld := QuestWorld{}
	row := q.Db.QueryRowContext(
		context.Background(),
		createQuestWorldConnectionQuery,
		params.FirstId,
		params.SecondId,
	)

	err := row.Scan(&questWorld.QuestId, &questWorld.WorldId)
	if err != nil {
		return nil, err
	}

	return &questWorld, nil
}

const createQuestAreaConnectionQuery = `
INSERT INTO quests_areas(quest_id, area_id)
	VALUES ($1, $2)
	RETURNING quest_id, area_id 
`

func (q *Queries) CreateQuestAreaConnection(
	ctx context.Context,
	params ConnectionParams,
) (*QuestArea, error) {
	questArea := QuestArea{}
	row := q.Db.QueryRowContext(
		context.Background(),
		createQuestAreaConnectionQuery,
		params.FirstId,
		params.SecondId,
	)

	err := row.Scan(&questArea.QuestId, &questArea.AreaId)
	if err != nil {
		return nil, err
	}

	return &questArea, nil
}

const createQuestLocationConnectionQuery = `
INSERT INTO quests_locations(quest_id, location_id)
	VALUES ($1, $2)
	RETURNING quest_id, location_id 
`

func (q *Queries) CreateQuestLocationConnection(
	ctx context.Context,
	params ConnectionParams,
) (*QuestLocation, error) {
	questLocation := QuestLocation{}
	row := q.Db.QueryRowContext(
		context.Background(),
		createQuestLocationConnectionQuery,
		params.FirstId,
		params.SecondId,
	)

	err := row.Scan(&questLocation.QuestId, &questLocation.LocationId)
	if err != nil {
		return nil, err
	}

	return &questLocation, nil
}

// Get Functions

const getNpcConnectedQuestsQuery = `
SELECT q.* FROM quests AS q
	INNER JOIN npc_quests AS nq ON nq.npc_id
	WHERE nq.npc_id = $1
`

func (q *Queries) GetNpcConnectedQuests(
	ctx context.Context,
	npcId int,
) ([]*Quest, error) {
	rows, err := q.Db.QueryContext(
		context.Background(),
		getNpcConnectedQuestsQuery,
		npcId,
	)
	if err != nil {
		return nil, err
	}

	quests := []*Quest{}

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
