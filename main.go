package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

const (
	gameWidth  = 100
	gameHeight = 50
)

type Rock struct {
	X int
	Y int
}

var rocks = []Rock{
	{10, 10},
	{20, 15},
	{40, 5},
}

func main() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	x := 5
	y := 5
	running := true

	for running {
		drawBoard(x, y)
		initializeRocks(rocks)

		// Draw player
		fmt.Printf("\033[%d;%dH", y, x)
		fmt.Printf("#")

		// Check for collision with rocks
		for _, rock := range rocks {
			if rock.X == x && rock.Y == y {
				running = false
				break
			}
		}

		x, y, running = processInput(x, y, running)
	}
}

func processInput(x, y int, running bool) (outX int, outY int, outRunning bool) {
	outX, outY, outRunning = x, y, running // Set the initial values for outX, outY, and outRunning

	switch getChar() {
	case 3:
		outRunning = false
	case 'w':
		if y > 1 {
			outY = y - 1
		}
	case 'a':
		if x > 1 {
			outX = x - 1
		}
	case 's':
		if y < 24 {
			outY = y + 1
		}
	case 'd':
		if x < 79 {
			outX = x + 1
		}
	}
	return outX, outY, outRunning
}

func initializeRocks(rocks []Rock) {
	for _, rock := range rocks {
		fmt.Printf("\033[%d;%dH", rock.Y, rock.X)
		fmt.Printf("O")
	}
}

// Read from the terminal
func getChar() byte {
	var buf [1]byte
	_, err := os.Stdin.Read(buf[:])
	if err != nil {
		fmt.Println("err reading from stdin : ", err)
		return 0
	}
	return buf[0]
}

func drawBoard(x, y int) {
	fmt.Print("\033[2J") // Clear the entire screen

	// Draw the box around the game area
	for i := 1; i <= gameWidth; i++ {
		fmt.Printf("\033[%d;%dH", 1, i)
		fmt.Print("-")
		fmt.Printf("\033[%d;%dH", gameHeight, i)
		fmt.Print("-")
	}
	for i := 1; i <= gameHeight; i++ {
		fmt.Printf("\033[%d;%dH", i, 1)
		fmt.Print("|")
		fmt.Printf("\033[%d;%dH", i, gameWidth)
		fmt.Print("|")
	}

	// Print debug information
	fmt.Printf("\033[%d;%dH", gameHeight+1, 1)
	fmt.Printf("Player position: (%d, %d)", x, y)
}
