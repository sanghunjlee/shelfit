package shelfit

type Status int16

const (
	Unset  Status = -1
	Unread Status = iota
	Started
	Finished
)

func (s Status) String() string {
	switch s {
	case Unread:
		return "unread"
	case Started:
		return "started"
	case Finished:
		return "finished"
	case Unset:
		return "unset"
	}
	return "nil"
}
