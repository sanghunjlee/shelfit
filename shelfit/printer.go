package shelfit

import (
	"fmt"
	"strconv"
	"strings"
)

type Printer interface {
	Print(*NeatShelf) error
}

type ShellPrinter struct {
	Queue [][]string
}

func NewShellPrinter() *ShellPrinter {
	return &ShellPrinter{Queue: [][]string{}}
}

func (p *ShellPrinter) Print(neatShelf *NeatShelf) error {
	for mapKey, neatBooks := range neatShelf.Books {
		for index, book := range neatBooks {
			if index == 0 {
				p.AddQueue(mapKey)
				p.AddQueue(
					"Id",
					"Title",
					"Category",
					"Tags",
				)
			}
			p.AddQueue(
				strconv.Itoa(book.Id),
				book.Title,
				book.Category,
				strings.Join(addPrefix("#", book.Tags), " "),
			)
		}
	}
	p.PrintQueue()
	return nil
}

func (p *ShellPrinter) AddQueue(args ...string) {
	p.Queue = append(p.Queue, args)
}

func (p *ShellPrinter) PrintQueue() {
	maxRow := len(p.Queue)
	if maxRow == 0 {
		return
	}
	widths := map[int]int{0: 1}
	for i := 0; i < maxRow; i++ {
		for index, value := range p.Queue[i] {
			if widths[index] < len(value) {
				widths[index] = len(value)
			}
		}
	}
	var padding int
	var totalWidth int
	for _, v := range widths {
		totalWidth += v
	}
	switch {
	case totalWidth > 100:
		padding = 1
	case totalWidth < 20:
		padding = 5
	default:
		padding = 100 / totalWidth
	}

	for _, q := range p.Queue {
		var line string
		for index, word := range q {
			line += ljust(word, widths[index]+padding)
		}
		fmt.Println(line)
	}
}
