package main

import (
	"os"

	"golang.org/x/term"
)

func main() {
	cls()
	defer cls()
	hideCursor()
	defer showCursor()
	defer term.Restore(int(os.Stdin.Fd()), makeRawTerminal())

	x, y, running, rocks := setupGame()
	prevX, prevY := x, y
	for running {
		clearPrevPos(prevY, prevX)
		drawPlayer(y, x)

		running = checkRockCollision(rocks, x, y, running)

		prevX, prevY = x, y // Update previous position
		x, y, running = processInput(x, y, running)
	}
}
