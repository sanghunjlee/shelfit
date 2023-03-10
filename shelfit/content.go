package shelfit

type Content struct {
	Title    string
	Archived bool

	Category     string
	Genres       []string
	Cover        int
	Volumes      []int
	VolumeTitles []string
	Status       []Status
	Link         []string

	HasArchived bool

	HasCategory bool
	HasGenres   bool
	HasVolumes  bool
	HasStatus   bool
	HasLink     bool
}

func (c *Content) PrevStatus() Status {
	if len(c.Status) == 0 {
		return Unset
	}
	return c.Status[len(c.Status)-1]
}
