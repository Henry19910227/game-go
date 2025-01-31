package game

type HandlerFunc func(c *Context)

type Status int

const (
	Unknown = -1
	Betting = 0
	Deal    = 1
	Settle  = 2
)

type StageLinkedList struct {
	head    *Stage
	tail    *Stage
	current *Stage
}

func (s *StageLinkedList) Next() {
	s.current = s.current.next
}

func (s *StageLinkedList) Current() *Stage {
	return s.current
}

type Stage struct {
	Countdown int
	Handler   HandlerFunc
	next      *Stage
}
