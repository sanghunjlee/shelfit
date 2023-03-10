package shelfit

import (
	"errors"
	"reflect"
)

type Organizer struct{}

type NeatShelf struct {
	Books map[string][]*Book
}

func (o *Organizer) GroupBy(shelf *Shelf, expand bool, groupBy string) (*NeatShelf, error) {
	groupedBooks := map[string][]*Book{}
	var key string = ""
	var err error = nil

	_, groupBy = correctFieldName(groupBy, Book{})

	for _, b := range shelf.Books {
		if b.Cover == -1 {
			switch groupBy {
			case "", "All", "all", "Any", "any":
				key = ""
			default:
				keyValue := reflect.ValueOf(b).Elem().FieldByName(groupBy)
				if keyValue == (reflect.Value{}) {
					err = errors.New("The groupBy keyword " + groupBy + " is not a field")
					key = ""
				} else {
					key = keyValue.String()
				}
			}
			groupedBooks[key] = append(groupedBooks[key], b)
			if expand {
				for _, v := range b.Volumes {
					groupedBooks[key] = append(groupedBooks[key], shelf.FindById(v))
				}
			}
		}
	}

	return &NeatShelf{Books: groupedBooks}, err
}
