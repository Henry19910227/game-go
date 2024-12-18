package bet

type Output struct {
	Table
}

func (Output) TableName() string {
	return "bets"
}
