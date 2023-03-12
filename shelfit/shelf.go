package shelfit

type Shelf struct {
	Name  string  `json:"name"`
	UUID  string  `json:"uuid"`
	Books []*Book `json:"book_items_attrs"`
}

func (s *Shelf) Load(books []*Book) {
	s.Books = books
}

func (s *Shelf) Add(book *Book) int {
	book.Id = s.getNextId()
	s.Books = append(s.Books, book)
	return book.Id
}

func (s *Shelf) Delete(ids ...int) {
	for _, id := range ids {
		index := s.GetIndexById(id)
		if index == -1 {
			continue
		}
		s.Books = append(s.Books[:index], s.Books[index+1:]...)
	}
}

func (s *Shelf) FindIdByTitle(title string) int {
	for _, b := range s.Books {
		if compareStrings(title, b.Title) {
			return b.Id
		}
	}
	return -1
}

func (s *Shelf) FindById(id int) *Book {
	for _, book := range s.Books {
		if book.Id == id {
			return book
		}
	}
	return nil
}

func (s *Shelf) GetIndexById(id int) int {
	for index, b := range s.Books {
		if b.Id == id {
			return index
		}
	}
	return -1
}

func (s *Shelf) getMaxId() int {
	maxId := 0
	for _, book := range s.Books {
		if book.Id > maxId {
			maxId = book.Id
		}
	}
	return maxId
}

func (s *Shelf) getNextId() int {
	if len(s.Books) == 0 {
		return 0
	}
	var found bool
	maxId := s.getMaxId()
	for i := 0; i < maxId; i++ {
		found = false
		for _, book := range s.Books {
			if book.Id == i {
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
