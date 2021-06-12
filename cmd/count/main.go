package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func check(e error) {
	if e != nil {
		fmt.Println("error:", e)
	}
}

func main() {
	args := os.Args[1:]
	path := args[0]

	f, err := os.Open(path)
	check(err)

	r := bufio.NewReader(f)

	m := make(map[string]uint)

	count := 0

	for c, _, err := r.ReadRune(); err == nil; c, _, err = r.ReadRune() {
		// ignore non chinese characters
		if !unicode.Is(unicode.Han, c) {
			continue
		}

		// if char has not been seen yet, add to count
		if _, ok := m[string(c)]; !ok {
			count++
			m[string(c)] = 1
		}
	}

	fmt.Println(count)
}
