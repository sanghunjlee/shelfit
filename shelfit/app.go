package shelfit

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// Current version of ShelfIt
	VERSION string = "0.0.1"
	ISO8601 string = "2006-01-02T15:04:05Z07:00"
)

type App struct {
	Room    Room
	Shelf   *Shelf
	Printer Printer
}

func NewApp() *App {
	app := &App{
		Room:    NewBedRoom(),
		Shelf:   &Shelf{},
		Printer: NewShellPrinter(),
	}
	return app
}

func (a *App) Initialize() {
	a.Room.Initialize()
	fmt.Println("Shelf initialized.")
}

func (a *App) AddBook(input string) {
	a.load()

	author := &Author{}

	content, err := author.Write(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	book, err := MakeBook(content)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	a.Shelf.Add(book)

	a.save()
	fmt.Printf("Book %d added to the shelf.", book.Id)
}

func (a *App) DeleteBook(input string) {
	a.load()
	ids := a.getIds(input)
	if len(ids) == 0 {
		return
	}
	a.Shelf.Delete(ids...)
	a.save()
	for _, id := range ids {
		fmt.Printf("Book %d deleted.", id)
	}
}

func (a *App) ClearBooks(input string) {
	a.load()
	a.Shelf.Clear()
	a.save()
}

func (a *App) ListBooks(input string, expand bool, groupBy string) {
	a.load()

	org := &Organizer{}

	var neatShelf *NeatShelf

	neatShelf, err := org.GroupBy(a.Shelf, expand, groupBy)
	if err != nil {
		fmt.Println(err.Error())
	}

	a.Printer.Print(neatShelf)
}

func (a *App) load() error {
	books, err := a.Room.Load()
	if err != nil {
		return err
	}
	a.Shelf.Load(books)
	return nil
}

func (a *App) save() {
	a.Room.Save(a.Shelf.Books)
}

func (a *App) getIds(text string) (ids []int) {
	textFrags := strings.Split(text, ",")
	for _, frag := range textFrags {
		if rangeIds := parseRangeInt(frag); len(rangeIds) > 0 {
			ids = append(ids, rangeIds...)
		} else if id, err := strconv.Atoi(frag); err == nil {
			ids = append(ids, id)
		}
	}
	return ids
}
