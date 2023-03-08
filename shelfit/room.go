package shelfit

import (
	"encoding/json"
	"fmt"
	"os"
)

type Room interface {
	Initialize()
	Load() ([]*Book, error)
	Save(books []*Book)
}

const ShelfFile = ".shelf.json"

type BedRoom struct {
	Found bool
}

func NewBedRoom() *BedRoom {
	return &BedRoom{Found: false}
}

func (b *BedRoom) Initialize() {
	f, err := os.OpenFile(ShelfFile, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		fmt.Println("Shelf file already exists!")
		os.Exit(0)
	}
	f.Write([]byte("[]"))
	if err1 := f.Close(); err1 != nil {
		fmt.Println("Error: Failed writing json file..", err)
		os.Exit(1)
	}
}

func (b *BedRoom) Load() ([]*Book, error) {
	data, err := os.ReadFile(ShelfFile)
	if err != nil {
		fmt.Println("Shelf is not found!")
		fmt.Println("Initialize a new shelf by typing 'shelfit init'.")
		os.Exit(0)
		return nil, err
	}
	var books []*Book
	jsonError := json.Unmarshal(data, &books)
	if jsonError != nil {
		fmt.Println("Error: Failed reading json data..", jsonError)
		os.Exit(1)
		return nil, jsonError
	}
	b.Found = true
	return books, nil
}

func (b *BedRoom) Save(books []*Book) {
	for _, book := range books {
		if book.UUID == "" {
			book.UUID = newUUID()
		}
	}
	data, _ := json.Marshal(books)
	if err := os.WriteFile(ShelfFile, []byte(data), 0644); err != nil {
		fmt.Println("Error: Failed writing json file..", err)
	}
}
