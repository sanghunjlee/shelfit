package shelfit

import "time"

type Book struct {
	Id           int       `json:"id"`
	UUID         string    `json:"uuid"`
	Title        string    `json:"title"`
	Category     string    `json:"category"`
	Genres       []string  `json:"genres"`
	Volumes      []*Volume `json:"volumes"`
	Link         string    `json:"link"`
	Status       Status    `json:"status"`
	FinishedDate string    `json:"finished_date"`
	StartedDate  string    `json:"started_date"`
	AddedDate    string    `json:"added_date"`
	Archived     bool      `json:"archived"`
	Notes        []string  `json:"notes"`
}

func MakeBook(content *Content) (*Book, error) {
	book := &Book{
		Title:    content.Title,
		Category: content.Category,
		Genres:   content.Genres,
		Status:   content.Status[0],
	}
	currTime := timeStamp(time.Now()).Format(ISO8601)
	switch book.Status {
	case Finished:
		book.FinishedDate = currTime
		fallthrough
	case Started:
		book.StartedDate = currTime
		fallthrough
	case Unread:
		book.AddedDate = currTime
	}

	for i, v := range content.VolumeTitles {
		volume := &Volume{
			Title:  v,
			Status: content.Status[i+1],
		}
		switch volume.Status {
		case Finished:
			volume.FinishedDate = currTime
			fallthrough
		case Started:
			volume.StartedDate = currTime
			fallthrough
		case Unread:
			volume.AddedDate = currTime
		}
		volume.Id = book.getNextId()
		book.Volumes = append(book.Volumes, volume)
	}

	return book, nil
}

func (b *Book) getMaxId() int {
	maxId := 0
	for _, volume := range b.Volumes {
		if volume.Id > maxId {
			maxId = volume.Id
		}
	}
	return maxId
}

func (b *Book) getNextId() int {
	if len(b.Volumes) == 0 {
		return 0
	}
	var found bool
	maxId := b.getMaxId()
	for i := 0; i < maxId; i++ {
		found = false
		for _, volume := range b.Volumes {
			if volume.Id == i {
				found = true
				break
			}
		}
		if !found {
			return i
		}
	}
	return maxId + 1
}
