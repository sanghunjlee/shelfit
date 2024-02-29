package shelfit

import "time"

type Book struct {
	Id           int      `json:"id"`
	UUID         string   `json:"uuid"`
	Title        string   `json:"title"`
	Category     string   `json:"category"`
	Tags         []string `json:"genres"`
	Link         string   `json:"link"`
	ModifiedDate string   `json:"modified_date"`
	AddedDate    string   `json:"added_date"`
	Archived     bool     `json:"archived"`
	Notes        []string `json:"notes"`
}

func MakeBook(content *Content) (*Book, error) {
	currTime := timeStamp(time.Now()).Format(ISO8601)
	book := &Book{
		Title:     content.Title,
		Category:  content.Category,
		Tags:      content.Tags,
		Notes:     []string{content.Note},
		AddedDate: currTime,
	}

	return book, nil
}
