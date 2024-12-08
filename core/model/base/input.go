package base

type ListInput struct {
	IsDeletedField
	CreateAtField
	UpdateAtField
}

type PagingInput struct {
	PageField
	SizeField
}

type Preload struct {
	Field      string
	Conditions []interface{}
}

type PreloadInput struct {
	Preloads []*Preload
}
