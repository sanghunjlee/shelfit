package shelfit

type Status int16

const (
	Unread Status = iota
	Started
	Finished
	Unset Status = -1
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

func (s Status) RelatedStrings() []string {
	ret := []string{s.String()}
	switch s {
	case Unread:
		return append(ret, "unwatched", "planned")
	case Started:
		return append(ret, "watching", "reading")
	case Finished:
		return append(ret, "watched", "completed")
	}
	return ret
}
