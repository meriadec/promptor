package promptor

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/k0kubun/go-ansi"
)

// Select method
func Select(items []string) int {

	selectedIndex := 0
	nbLines := printChoices(items, selectedIndex)

	// reset term for listen each keycode
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	b := make([]byte, 1)

	hasPressedArrow := false
	hasValidated := false

	for hasValidated == false {

		var direction string

		os.Stdin.Read(b)
		charCode := b[0]

		if charCode == 10 {
			break
		} else if charCode == 91 {
			hasPressedArrow = true
		} else if hasPressedArrow {
			switch charCode {
			case 65:
				direction = "top"
			case 66:
				direction = "bot"
			}
			hasPressedArrow = false
		}

		if direction != "" {
			eraseLines(nbLines)
			if direction == "top" {
				if selectedIndex > 0 {
					selectedIndex--
				} else {
					selectedIndex = len(items) - 1
				}
			} else if direction == "bot" {
				if selectedIndex < len(items)-1 {
					selectedIndex++
				} else {
					selectedIndex = 0
				}
			}
			printChoices(items, selectedIndex)
		}
	}

	return selectedIndex
}

func printChoices(items []string, selectedIndex int) int {

	nbLines := 0

	for i, choice := range items {

		choiceLines := strings.Split(choice, "\n")

		for y := 0; y < len(choiceLines); y++ {

			promptChar := "   "
			if y == 0 {
				promptChar = "[ ]"
				if i == selectedIndex {
					promptChar = "[x]"
				}
			}

			fmt.Println(promptChar, choiceLines[y])
			nbLines++
		}
	}

	return nbLines
}

func eraseLines(nb int) {
	for i := 0; i < nb; i++ {
		ansi.CursorUp(1)
		ansi.EraseInLine(2)
	}
}
