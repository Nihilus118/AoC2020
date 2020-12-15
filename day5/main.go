package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type platz struct {
	reihe  int64
	spalte int64
}

func (p platz) getSitzplatzID() int64 {
	return p.reihe*8 + p.spalte
}

func main() {
	// Plätze aus Datei
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var plaetze []int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		temp := platz{}
		if err != nil {
			log.Fatal(err)
		}
		// Buchstaben zu binär, danach zu dezimal um schöner damit zu rechnen
		temp.reihe, _ = strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(string(scanner.Text()[:7]), "F", "0"), "B", "1"), 2, 64)
		temp.spalte, _ = strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(string(scanner.Text()[len(scanner.Text())-3:]), "L", "0"), "R", "1"), 2, 64)
		// Jetzt nurnoch den Sitzplatz ausrechnen und in Slice
		plaetze = append(plaetze, temp.getSitzplatzID())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Absteigend sortieren um den letzten Sitzplatz zu finden
	sort.Slice(plaetze, func(i, j int) bool { return plaetze[i] > plaetze[j] })
	fmt.Println("Teil 1:", plaetze[0])

	// Teil 2
	for i := 0; i < len(plaetze)-1; i++ {
		// Ab wann wird nichtmehr normal weiter gezählt?
		if plaetze[i] != plaetze[0]-int64(i) {
			fmt.Println("Teil 2:", plaetze[i]+1)
			break
		}
	}
}
