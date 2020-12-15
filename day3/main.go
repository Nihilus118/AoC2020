package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type vec2 struct {
	x int
	y int
}

func main() {
	// Passwörter aus Datei
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var trees []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err != nil {
			log.Fatal(err)
		}
		// Karte in Stringslice
		trees = append(trees, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Wir haben die Karte, wie viel Platz ist überhaupt nach rechts?
	breite := len(trees[0])

	// Richtung variabel, löst also alle Aufgaben, auch von Teil2
	richtung := vec2{3, 1}

	// Los fahren
	unfaelle := 0
	pos := vec2{0, 0}
	// Solange bis wir am Ziel sind
	for pos.y < len(trees)-richtung.y {
		// In X-Richtung und immer wieder links anfangen
		pos.x = (pos.x + richtung.x) % breite
		// In Y-Richtung
		pos.y += richtung.y

		if string(trees[pos.y][pos.x]) == "#" {
			// Bum
			unfaelle++
		}
	}
	fmt.Println(unfaelle)
}
