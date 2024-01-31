package main

import (
	"flag"
	"fmt"
	"math/rand"
	"slices"
)

func main() {
	notes := []string{"A", "B", "C", "D", "E", "F", "G"}
	var n int
	var remove string
	flag.IntVar(&n, "n", 1, "Number of note")
	flag.StringVar(&remove, "r", "", "Remove note from the list")
	flag.Parse()
	if n <= 0 {
		panic("Notes must be greater than 0")
	}
	for i := 1; i < n; i++ {
		notes = append(notes, notes...)
	}
	if remove != "" {
		for _, s := range remove {
			slices.DeleteFunc(notes, func(i string) bool {
				return i == string(s)
			})
		}
		notes = notes[:len(notes)-len(remove)*n]
	}
	rand.Shuffle(len(notes), func(i, j int) { notes[i], notes[j] = notes[j], notes[i] })
	fmt.Println(notes)
}
