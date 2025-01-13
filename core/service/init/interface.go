package init

type Service interface {
	InitBetAreaCache() (err error)
}
