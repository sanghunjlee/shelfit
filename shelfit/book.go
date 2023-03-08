package shelfit

type Book struct {
	Id           int      `json:"id"`
	UUID         string   `json:"uuid"`
	Title        string   `json:"title"`
	Category     string   `json:"category"`
	Genres       []string `json:"genres"`
	Cover        int      `json:"volumes`
	Volumes      []int    `json:"volumes"`
	Link         string   `json:"link"`
	Status       string   `json:"status"`
	Finished     bool     `json:"finished"`
	FinishedDate string   `json:"finished_date"`
	Started      bool     `json:"started"`
	StartedDate  string   `json:"started_date"`
	AddedDate    string   `json:"added_date"`
	Archived     bool     `json:"archived"`
	Notes        []string `json:"notes"`
}
