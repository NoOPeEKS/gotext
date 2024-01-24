package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

func print_message(col int, row int, fg termbox.Attribute, bg termbox.Attribute, message string) {
	for _, ch := range message {
		termbox.SetCell(col, row, ch, fg, bg)
		col += runewidth.RuneWidth(ch)
	}
}

func run_editor() {
	err := termbox.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	print_message(
		25,
		11,
		termbox.ColorDefault,
		termbox.ColorDefault,
		"GoText - A simple and lightweight text editor in Go.",
	)
	termbox.Flush()
	termbox.PollEvent()
	termbox.Close()
}

func main() {
	run_editor()
}
