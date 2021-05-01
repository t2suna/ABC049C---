package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	s := nextLine()

	if len(s) == 0 {
		fmt.Print("NO")
		os.Exit(0)
	}

	for len(s) > 0 {
		if strings.HasPrefix(s, "dream") {
			s = strings.Replace(s, "dream", "", 1)
			if strings.HasPrefix(s, "eraser") {
				s = strings.Replace(s, "eraser", "", 1)
			} else if strings.HasPrefix(s, "erase") {
				s = strings.Replace(s, "erase", "", 1)
			} else if strings.HasPrefix(s, "er") {
				s = strings.Replace(s, "er", "", 1)
			}
		} else if strings.HasPrefix(s, "eraser") {
			s = strings.Replace(s, "eraser", "", 1)
		} else if strings.HasPrefix(s, "erase") {
			s = strings.Replace(s, "erase", "", 1)
		} else {
			fmt.Print("NO")
			os.Exit(0)
		}
	}
	fmt.Print("YES")
}
