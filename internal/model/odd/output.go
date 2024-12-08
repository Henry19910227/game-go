package odd

type Output struct {
	Table
}

func (Output) TableName() string {
	return "odds"
}
