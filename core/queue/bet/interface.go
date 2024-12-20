package bet

type Queue interface {
	Write(data []byte) (err error)
}
