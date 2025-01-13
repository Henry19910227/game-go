package init

type Adapter interface {
	InitBetAreaCache() (err error)
}
