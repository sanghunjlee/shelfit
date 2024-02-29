package shelfit

type Content struct {
	Title    string
	Archived bool

	Category string
	Tags     []string
	Note     string
	Link     []string

	HasArchived bool

	HasCategory bool
	HasTag      bool
	HasVolumes  bool
	HasStatus   bool
	HasLink     bool
}
