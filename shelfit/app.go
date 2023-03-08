package shelfit

import "fmt"

const (
	// Current version of ShelfIt
	VERSION string = "0.0.1"
)

type App struct {
	Room  Room
	Shelf *Shelf
}

func NewApp() *App {
	app := &App{
		Room:  NewBedRoom(),
		Shelf: &Shelf{},
	}
	return app
}

func (a *App) Initialize() {
	a.Room.Initialize()
	fmt.Println("Shelf initialized.")
}

func (a *App) AddBook(text string) {
	a.load()
}

func (a *App) load() error {
	books, err := a.Room.Load()
	if err != nil {
		return err
	}
	a.Shelf.Load(books)
	return nil
}

func (a *App) save() error {
	a.Room.Save(a.Shelf.Books)
}
