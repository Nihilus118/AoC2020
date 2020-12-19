package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Gruppen aus Datei
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var gruppen []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err != nil {
			log.Fatal(err)
		}

		gruppen = append(gruppen, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Durch Gruppen und zählen (Teil 1)
	summe := 0
	var temp string
	for _, gruppe := range gruppen {
		if gruppe == "" {
			// Gruppe zuende, also in Liste
			summe += len(temp)
			temp = ""
		} else {
			for _, buchstabe := range gruppe {
				// Duplikate verhindern
				if strings.Index(temp, string(buchstabe)) == -1 {
					temp += string(buchstabe)
				}
			}
		}
	}

	// Letzte Gruppe nicht vernachlässigen
	if temp != "" {
		summe += len(temp)
	}

	fmt.Println("Teil 1:", summe)

	// Anzahl des Buchstaben muss gleich Anzahl der Personen in der Gruppe sein
	anzPersonen := 0
	buchstabenZaehler := map[rune]int{
		'a': 0,
		'b': 0,
		'c': 0,
		'd': 0,
		'e': 0,
		'f': 0,
		'g': 0,
		'h': 0,
		'i': 0,
		'j': 0,
		'k': 0,
		'l': 0,
		'm': 0,
		'n': 0,
		'o': 0,
		'p': 0,
		'q': 0,
		'r': 0,
		's': 0,
		't': 0,
		'u': 0,
		'v': 0,
		'w': 0,
		'x': 0,
		'y': 0,
		'z': 0,
	}
	summe = 0
	for _, gruppe := range gruppen {
		if gruppe == "" {
			// Gruppe zuende also Ergebnis merken
			for i := range buchstabenZaehler {
				if buchstabenZaehler[i] == anzPersonen {
					summe++
				}
			}

			// Neue Gruppe also neu Zählen
			anzPersonen = 0
			for i := range buchstabenZaehler {
				buchstabenZaehler[i] = 0
			}
		} else {
			// Jeden Buchstaben zählen
			anzPersonen++
			for _, buchstabe := range gruppe {
				buchstabenZaehler[buchstabe]++
			}
		}
	}

	fmt.Println("Teil 2:", summe)
}
