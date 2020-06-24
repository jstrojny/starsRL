package main

import(
	term "github.com/jstrojny/starsRL/bearlibterminal"
)

const(
	screenWidth int = 80
	screenHeight int = 50
)

func main() {
	term.Open()
	term.Set("window.title=starsRL; window.size=80x50;")
	term.Clear()
	term.Print(25, 10, "@")
	term.Refresh()
	closed := 0
	charX := 25
	charY := 10
	for closed != 1 {
		input := term.Read()
		if input == term.TK_CLOSE {
			term.Close()
			closed = 1
		}
		if input == term.TK_ESCAPE {
			term.Close()
			closed = 1
		}
		if input == term.TK_KP_8 {
			charY--
			term.Clear()
			term.Print(charX, charY, "@")
			term.Refresh()
		} else if input == term.TK_KP_2 {
			charY++
			term.Clear()
			term.Print(charX, charY, "@")
			term.Refresh()
		} else if input == term.TK_KP_4 {
			charX--
			term.Clear()
			term.Print(charX, charY, "@")
			term.Refresh()
		} else if input == term.TK_KP_6 {
			charX++
			term.Clear()
			term.Print(charX, charY, "@")
			term.Refresh()
		} else if input == term.TK_KP_7 {
			charX--
			charY--
			term.Clear()
			term.Print(charX, charY, "@")
			term.Refresh()
		} else if input == term.TK_KP_9 {
			charX++
			charY--
			term.Clear()
			term.Print(charX, charY, "@")
			term.Refresh()
		} else if input == term.TK_KP_1 {
			charX--
			charY++
			term.Clear()
			term.Print(charX, charY, "@")
			term.Refresh()
		} else if input == term.TK_KP_3 {
			charX++
			charY++
			term.Clear()
			term.Print(charX, charY, "@")
			term.Refresh()
		}

	}
}