package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Olexander63/goFundamentals/08/packages/menu"
)

var in = bufio.NewReader(os.Stdin)

func main() {

loop:
	for {
		fmt.Println("Please select an option")
		fmt.Println("1) Print menu")
		fmt.Println("2) Add item")
		fmt.Println("q) quit")
		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			menu.Print()
		case "2":
			menu.Add()
		case "q":
			break loop
		default:
			fmt.Println("Unknown option")
		}
	}
}
