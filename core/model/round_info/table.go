package round_info

type Table struct {
	IDField
	GameIdField
	TypeField
	ElementsField
	PatternsField
	ResultsField
}

func (Table) TableName() string {
	return "round_infos"
}
