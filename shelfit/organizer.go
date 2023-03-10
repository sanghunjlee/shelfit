package shelfit

type Organizer struct{}

type NeatShelf struct {
	Books map[string][]*Book
}

func (o *Organizer) GroupByCategory(shelf *Shelf) *NeatShelf {
	groupedBooks := map[string][]*Book{}

	for _, b := range shelf.Books {
		groupedBooks[b.Category] = append(groupedBooks[b.Category], b)
	}

	return &NeatShelf{Books: groupedBooks}
}

func (o *Organizer) GroupByGenre(shelf *Shelf) *NeatShelf {
	groupedBooks := map[string][]*Book{}

	for _, b := range shelf.Books {
		for _, g := range b.Genres {
			groupedBooks[g] = append(groupedBooks[g], b)
		}
	}
	return &NeatShelf{Books: groupedBooks}
}
