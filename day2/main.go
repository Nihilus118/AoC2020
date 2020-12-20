package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type passwort struct {
	text      string
	buchstabe string
	min       int
	max       int
}

func pwAusDatei(zeile string) (p passwort) {
	split := strings.Split(zeile, ":")
	p.text = strings.TrimSpace(split[1])
	split = strings.Split(split[0], " ")
	p.buchstabe = split[1]
	split = strings.Split(split[0], "-")
	var err error
	p.min, err = strconv.Atoi(split[0])
	if err != nil {
		log.Fatal(err)
	}
	p.max, err = strconv.Atoi(split[1])
	if err != nil {
		log.Fatal(err)
	}

	return p
}

// Teil 1
func (p passwort) validate1() bool {
	anz := strings.Count(p.text, p.buchstabe)
	if anz >= p.min && anz <= p.max {
		return true
	}
	return false
}

// Teil 2
func (p passwort) validate2() bool {
	// -1 weil es nicht mit Index 0 laufen soll
	zeichen1, zeichen2 := string(p.text[p.min-1]), string(p.text[p.max-1])
	if (zeichen1 == p.buchstabe) != (zeichen2 == p.buchstabe) {
		return true
	}
	return false
}

func main() {
	// Teil 1
	// Passwörter aus Datei
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	valide1 := 0
	valide2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Ist das Passwort ok?
		pw := pwAusDatei(scanner.Text())
		// Teil 1
		if pw.validate1() {
			valide1++
		}
		// Teil 2
		if pw.validate2() {
			valide2++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Teil 1:", valide1)
	fmt.Println("Teil 2:", valide2)
}
