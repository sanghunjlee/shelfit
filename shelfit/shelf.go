package shelfit

type Shelf struct {
	Name  string  `json:"name"`
	UUID  string  `json:"uuid"`
	Books []*Book `json:"book_items_attrs"`
}

func (s *Shelf) Load(books []*Book) {
	s.Books = books
}

func (s *Shelf) Add(book *Book) {
	book.Id = s.NextId()
	s.Books = append(s.Books, book)
}
