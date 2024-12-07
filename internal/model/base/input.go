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
