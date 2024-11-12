package dndapi

type MonsterSearchResp struct {
	Count   int                   `json:"count"`
	Results []MonsterSearchResult `json:"results"`
}

type MonsterSearchResult struct {
	Index string `json:"index"`
	Name  string `json:"name"`
}

type Monster struct {
	Index                 string             `json:"index"`
	Name                  string             `json:"name"`
	Size                  string             `json:"size"`
	Type                  string             `json:"type"`
	Alignment             string             `json:"alignment"`
	ArmorClass            []ArmorClass       `json:"armor_class"`
	HitPoints             int                `json:"hit_points"`
	HitDice               string             `json:"hit_dice"`
	HitPointsRoll         string             `json:"hit_points_roll"`
	Speed                 Speed              `json:"speed"`
	Strength              int                `json:"strength"`
	Dexterity             int                `json:"dexterity"`
	Constitution          int                `json:"constitution"`
	Intelligence          int                `json:"intelligence"`
	Wisdom                int                `json:"wisdom"`
	Charisma              int                `json:"charisma"`
	Proficiencies         []Proficiencies    `json:"proficiencies"`
	DamageVulnerabilities []any              `json:"damage_vulnerabilities"`
	DamageResistances     []any              `json:"damage_resistances"`
	DamageImmunities      []string           `json:"damage_immunities"`
	ConditionImmunities   []any              `json:"condition_immunities"`
	Senses                Senses             `json:"senses"`
	Languages             string             `json:"languages"`
	ChallengeRating       float32                `json:"challenge_rating"`
	ProficiencyBonus      int                `json:"proficiency_bonus"`
	Xp                    int                `json:"xp"`
	SpecialAbilities      []SpecialAbilities `json:"special_abilities"`
	Actions               []Actions          `json:"actions"`
	LegendaryActions      []LegendaryActions `json:"legendary_actions"`
	Image                 string             `json:"image"`
	URL                   string             `json:"url"`
}

type ArmorClass struct {
	Type  string `json:"type"`
	Value int    `json:"value"`
}

type Speed struct {
	Walk string `json:"walk"`
	Fly  string `json:"fly"`
	Swim string `json:"swim"`
}

type Proficiency struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type Proficiencies struct {
	Value       int         `json:"value"`
	Proficiency Proficiency `json:"proficiency"`
}

type Senses struct {
	Blindsight        string `json:"blindsight"`
	Darkvision        string `json:"darkvision"`
	PassivePerception int    `json:"passive_perception"`
}

type Usage struct {
	Type      string `json:"type"`
	Times     int    `json:"times"`
	RestTypes []any  `json:"rest_types"`
}

type SpecialAbilities struct {
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Usage Usage  `json:"usage,omitempty"`
}

type SubActions struct {
	ActionName string `json:"action_name"`
	Count      int    `json:"count"`
	Type       string `json:"type"`
}

type DamageType struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type Damage struct {
	DamageType DamageType `json:"damage_type"`
	DamageDice string     `json:"damage_dice"`
}

type DcType struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type Dc struct {
	DcType      DcType `json:"dc_type"`
	DcValue     int    `json:"dc_value"`
	SuccessType string `json:"success_type"`
}

type ActionUsage struct {
	Type     string `json:"type"`
	Dice     string `json:"dice"`
	MinValue int    `json:"min_value"`
}

type Actions struct {
	Name            string    `json:"name"`
	MultiattackType string    `json:"multiattack_type,omitempty"`
	Desc            string    `json:"desc"`
	Actions         []Actions `json:"actions"`
	AttackBonus     int       `json:"attack_bonus,omitempty"`
	Damage          []Damage  `json:"damage,omitempty"`
	Dc              Dc        `json:"dc,omitempty"`
	Usage           Usage     `json:"usage,omitempty"`
}

type LegendaryActions struct {
	Name   string   `json:"name"`
	Desc   string   `json:"desc"`
	Dc     Dc       `json:"dc,omitempty"`
	Damage []Damage `json:"damage,omitempty"`
}
