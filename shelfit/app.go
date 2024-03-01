package shelfit

import (
	"bufio"
	"fmt"
	"os"
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

func (a *App) AddBook(input string, note string) {
	a.load()

	author := &Author{}

	content, err := author.Write(input, note)
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
		fmt.Println(fmt.Errorf("no item found"))
		return
	}
	strIds := convertIntArrayToStringArr(ids)

	fmt.Printf("This will delete the following %d items:\n", len(ids))
	a.ListBooks(strings.Join(strIds, ","), "")

	fmt.Println("Continue deleting? [y/N]")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err.Error())
	}

	if char == 121 || char == 89 {
		a.Shelf.Delete(ids...)
		a.save()
		for _, id := range ids {
			fmt.Printf("Book %d deleted.\n", id)
		}
	}
}

func (a *App) ClearBooks(input string) {
	a.load()

	if len(a.Shelf.Books) > 0 {
		fmt.Printf("Clear %d item? [y/N]", len(a.Shelf.Books))
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println(err.Error())
		}

		if char == 121 || char == 89 {
			a.Shelf.Clear()
			a.save()
		}
	} else {
		fmt.Print("No items to clear")
	}
}

func (a *App) ListBooks(input string, category string) {
	var (
		neatShelf *NeatShelf
		books     []*Book
	)

	a.load()

	org := &Organizer{}
	ids := a.getIds(input)

	if len(ids) > 0 {
		books = make([]*Book, 0)

		for _, b := range a.Shelf.Books {
			for _, id := range ids {
				if id == b.Id {
					books = append(books, b)
					break
				}
			}
		}
	} else {
		books = a.Shelf.Books
	}

	neatShelf, err := org.GroupBy(books, category)
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
	var parsedIds []int
	textFrags := strings.Split(text, ",")
	for _, frag := range textFrags {
		if rangeIds := parseRangeInt(frag); len(rangeIds) > 0 {
			parsedIds = append(parsedIds, rangeIds...)
		} else if id, err := strconv.Atoi(frag); err == nil {
			parsedIds = append(parsedIds, id)
		}
	}
	for _, id := range parsedIds {
		for _, b := range a.Shelf.Books {
			if b.Id == id {
				ids = append(ids, id)
				break
			}
		}
	}
	return ids
}
