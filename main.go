package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	switch n {
	case 1:
		fmt.Printf("Yes\nNo\nNo\n")
	case 0:
		fmt.Printf("No\nNo\nNo\n")
	case 3:
		fmt.Printf("No\nYes\nNo\n")
	case 5:
		fmt.Printf("No\nNo\nYes\n")
	case 6:
		fmt.Printf("Yes\nNo\nYes\n")
	case 4:
		fmt.Printf("Yes\nYes\nNo\n")
	case 8:
		fmt.Printf("No\nYes\nYes\n")
	case 9:
		fmt.Printf("Yes\nYes\nYes\n")
	}

}
