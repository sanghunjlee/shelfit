package shelfit

import "time"

type Book struct {
	Id           int      `json:"id"`
	UUID         string   `json:"uuid"`
	Title        string   `json:"title"`
	Category     string   `json:"category"`
	Genres       []string `json:"genres"`
	Cover        int      `json:"cover"`
	Volumes      []int    `json:"volumes"`
	Link         string   `json:"link"`
	Status       Status   `json:"status"`
	FinishedDate string   `json:"finished_date"`
	StartedDate  string   `json:"started_date"`
	AddedDate    string   `json:"added_date"`
	Archived     bool     `json:"archived"`
	Notes        []string `json:"notes"`
}

func MakeBook(content *Content) (*Book, error) {
	book := &Book{
		Title:    content.Title,
		Category: content.Category,
		Genres:   content.Genres,
		Volumes:  content.Volumes,
		Status:   content.Status[0],
	}
	currTime := timeStamp(time.Now()).Format(ISO8601)
	switch book.Status {
	case Unread:
		book.AddedDate = currTime
		fallthrough
	case Started:
		book.StartedDate = currTime
		fallthrough
	case Finished:
		book.FinishedDate = currTime
	}
	return book, nil
}

func MakeVolumes(content *Content) ([]*Book, error) {
	var volumes []*Book
	var err error
	if content.HasVolumes {
		for _, volTitle := range content.VolumeTitles {
			volContent := content
			volContent.Title = volTitle
			volContent.Status = append(
				volContent.Status[:0],
				volContent.Status[1:]...,
			)
			var volume *Book
			volume, err = MakeBook(volContent)
			if err != nil {
				continue
			}
			volumes = append(volumes, volume)
		}
	}
	return volumes, err
}
