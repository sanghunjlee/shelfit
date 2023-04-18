package shelfit

type Volume struct {
	Id           int      `json:"id"`
	UUID         string   `json:"uuid"`
	Title        string   `json:"title"`
	Link         string   `json:"link"`
	Status       Status   `json:"status"`
	FinishedDate string   `json:"finished_date"`
	StartedDate  string   `json:"started_date"`
	AddedDate    string   `json:"added_date"`
	Notes        []string `json:"notes"`
}
