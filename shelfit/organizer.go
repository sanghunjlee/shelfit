package shelfit

import (
	"errors"
	"fmt"
)

type Organizer struct{}

type NeatShelf struct {
	Books map[string][]*Book
}

func (o *Organizer) GroupBy(shelf *Shelf, category string) (*NeatShelf, error) {
	groupedBooks := map[string][]*Book{}
	var categoryExist bool = false
	var bookKey string

	if len(shelf.Books) == 0 {
		return &NeatShelf{}, errors.New("empty shelf: there is no item to list")
	}

	for _, b := range shelf.Books {
		if category == "" || category == b.Category {
			categoryExist = true
			bookKey = b.Category
			if category == "" {
				bookKey = ""
			}
			groupedBooks[bookKey] = append(groupedBooks[bookKey], b)
		}
	}

	if !categoryExist {
		return &NeatShelf{}, fmt.Errorf("no category: there is no item with %s category", category)
	}

	return &NeatShelf{Books: groupedBooks}, nil
}
