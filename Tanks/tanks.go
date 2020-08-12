package tanks

import (
	"bufio"
	// "fmt"
	"github.com/eiannone/keyboard"
	"os"
	"os/exec"
	// "runtime"
)

func main() {
	// initial values
	gameStarted := false
	vl := 20
	hl := 80
	px := 10
	py := 5

	tank1 := tank{}
	tank1.posX = 10
	tank1.posY = 5

	tank2 := tank{}
	tank2.posX = 25
	tank2.posY = 7

	keyPressed, _, err := keyboard.GetSingleKey()

	if err != nil {
		panic(err)
	}

	println(keyPressed)

	scanner := bufio.NewScanner(os.Stdin)
	println("type start to begin or exit to quit.")

	for scanner.Scan() {
		if scanner.Text() == "start" {
			gameStarted = true
			printLevel(vl, hl, tank1.posX, tank1.posY, gameStarted)
		}
		if scanner.Text() == "w" && gameStarted == true {
			if py > 1 {
				py--
				printLevel(vl, hl, px, py, gameStarted)
			}
		}
		if scanner.Text() == "s" && gameStarted == true {
			if py < vl-2 {
				py++
				printLevel(vl, hl, px, py, gameStarted)
			}
		}
		if scanner.Text() == "a" && gameStarted == true {
			if px > 0 {
				px--
				printLevel(vl, hl, px, py, gameStarted)
			}
		}
		if scanner.Text() == "d" && gameStarted == true {
			if px < hl-1 {
				px++
				printLevel(vl, hl, px, py, gameStarted)
			}
		}
		if scanner.Text() == "f" && gameStarted == true {
			println("fire")
		}
		if scanner.Text() == "exit" {
			os.Exit(2)
		}
	}
}

func printLevel(ly int, lx int, px int, py int, gameStarted bool) {
	var lineToRender string
	if gameStarted == true {
		clearConsole()
		for y := 0; y < ly; y++ {
			lineToRender = "*"
			for x := 0; x < lx; x++ {
				if y == 0 || y == ly-1 {
					lineToRender += "*"
				} else {
					if x == px && y == py {
						lineToRender += "T"
					} else {
						lineToRender += " "
					}
				}
			}
			lineToRender += "*"
			println(lineToRender)
			lineToRender = ""
		}
	} else {
		println("game not started")
	}
}

func clearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
