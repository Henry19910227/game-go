package round_info

type Cache interface {
	Save(gameId int, data []byte) (err error)
}
