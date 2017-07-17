package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var s []string

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter Num N: ")

	input, err := inputReader.ReadString('\n')
	s := strings.Split(input, "\n")
	in, err := strconv.Atoi(s[0])
	if err != nil {
		fmt.Print(err)
	}
	if in > 1 && in <= 12 {
		fmt.Println(fullassembly(in))
	} else {
		fmt.Println("Your input is wrong!")
	}
}

func fullassembly(h int) []string {
	if h == 1 {
		s = append(s, "1")
		return s
	} else {
		a := fullassembly(h - 1)
		g := strconv.Itoa(h)
		var b []string
		for i := 0; i < len(a); i++ {
			b = append(b, a[i]+","+g)
		}

		a = append(a, g)
		s = append(a, b...)
		return s
	}
}
