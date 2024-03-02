package shelfit

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sanghunjlee/shelfit/window"
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
	Window  Window
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

func (a *App) EditBook(input string) {
	a.load()

	id, err := a.getId(input)
	book := a.Shelf.FindById(id)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	options := []string{
		"Title",
		"Category",
		"Tags",
	}

	selection, err := window.Select(options)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var (
		oldValue string
		value    string
	)

	// r := reflect.ValueOf(book)
	// f := reflect.Indirect(r).FieldByName(selection)
	// oldValue = f.String()

	switch selection {
	case options[0]: // Title
		ex := book.Title
		value, err = window.Prompt(fmt.Sprintf("Enter a new value for \"%s\"", selection), ex)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		oldValue = ex
		book.Title = value
	case options[1]: // Category
		ex := book.Category
		value, err = window.Prompt(fmt.Sprintf("Enter a new value for \"%s\"", selection), ex)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		oldValue = ex
		book.Category = value
	case options[2]: // Tags
		ex := book.Tags
		value, err = window.TagInput(ex)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		oldValue = strings.Join(ex, " ")
		book.Tags = strings.Split(value, " ")
	default:
		fmt.Println(fmt.Errorf("selection issue: %s is not an option", selection))
		return
	}

	a.save()
	fmt.Printf("Item (Id=%d) value changed for \"%s\":\n\tOld: %s\n\tNew: %s", book.Id, selection, oldValue, value)
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

func (a *App) getId(text string) (id int, err error) {
	id, err = strconv.Atoi(text)
	if err != nil {
		return -1, err
	}

	for _, b := range a.Shelf.Books {
		if b.Id == id {
			return id, nil
		}
	}

	return id, fmt.Errorf("no item found: there is no item with id: %d", id)
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
