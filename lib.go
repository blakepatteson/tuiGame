package main

import (
	"fmt"
	"math/rand"
	"os"

	"golang.org/x/term"
)

func makeRawTerminal() *term.State {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	return oldState
}

func processInput(x, y int, running bool) (outX int, outY int, outRunning bool) {
	outX, outY, outRunning = x, y, running // Set the initial values for outX, outY, and outRunning

	switch getChar() {
	case 3:
		fmt.Print("\033[2J") // Clear the entire screen
		outRunning = false
	case 'w', '8':
		if y > 2 {
			outY = y - 1
		}
	case 'a', '4':
		if x > 2 {
			outX = x - 1
		}
	case 's', '2':
		if y < GAME_HEIGHT-1 {
			outY = y + 1
		}
	case 'd', '6':
		if x < GAME_WIDTH-1 {
			outX = x + 1
		}
	}
	return outX, outY, outRunning
}

// Read user input
func getChar() byte {
	var buf [1]byte
	_, err := os.Stdin.Read(buf[:])
	if err != nil {
		fmt.Println("err reading from stdin : ", err)
		return 0
	}
	return buf[0]
}

// func debugPrintPlayerPos(x, y int) {
// 	// Print debug information
// 	fmt.Printf("\033[%d;%dH", GAME_HEIGHT+1, GAME_WIDTH/2)
// 	fmt.Printf("Player position: (%d, %d)", x, y)
// }

func gameOver() {
	gameOverMsg := "game over!"
	fmt.Printf("\033[%d;%dH", GAME_HEIGHT+2, GAME_WIDTH/2-len(gameOverMsg)/2)
	fmt.Printf("%v\n", gameOverMsg)
}

// Draw the box around the game area
func drawBoard() {
	for i := 1; i <= GAME_WIDTH; i++ {
		fmt.Printf("\033[%d;%dH", 1, i)
		fmt.Print("-")
		fmt.Printf("\033[%d;%dH", GAME_HEIGHT, i)
		fmt.Print("-")
	}
	for i := 1; i <= GAME_HEIGHT; i++ {
		fmt.Printf("\033[%d;%dH", i, 1)
		fmt.Print("|")
		fmt.Printf("\033[%d;%dH", i, GAME_WIDTH)
		fmt.Print("|")
	}
}

func generateRandomRocks() []Rock {
	rocks := make([]Rock, NUM_ROCKS)
	for i := 0; i < NUM_ROCKS; i++ {
		rock := Rock{
			X: rand.Intn(GAME_WIDTH-2) + 2,  // Add 1 to avoid the left border, subtract 2 to avoid the right border
			Y: rand.Intn(GAME_HEIGHT-2) + 2, // Add 1 to avoid the top border, subtract 2 to avoid the bottom border
		}
		rocks[i] = rock
	}
	return rocks
}

func cls() {
	fmt.Print("\033[2J") // Clear the entire screen
}

func checkRockCollision(rocks []Rock, x int, y int, running bool) bool {
	for _, rock := range rocks {
		if rock.X == x && rock.Y == y {
			gameOver()
			running = false
			break
		}
	}
	return running
}

func drawPlayer(y int, x int) {
	fmt.Printf("\033[%d;%dH", y, x)
	fmt.Printf("#")
}

func clearPrevPos(prevY int, prevX int) {
	fmt.Printf("\033[%d;%dH", prevY, prevX)
	fmt.Printf(" ")
}

func initializeRocks(rocks []Rock) {
	for _, rock := range rocks {
		fmt.Printf("\033[%d;%dH", rock.Y, rock.X)
		fmt.Printf("O")
	}
}

func setupGame() (int, int, bool, []Rock) {
	x := 5
	y := 5
	running := true
	rocks := generateRandomRocks()
	drawBoard()
	initializeRocks(rocks)
	return x, y, running, rocks
}

func hideCursor() {
	fmt.Print("\033[?25l")
}

func showCursor() {
	fmt.Print("\033[?25h")
}
