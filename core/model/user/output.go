package user

type Output struct {
	Table
}

func (Output) TableName() string {
	return "users"
}
