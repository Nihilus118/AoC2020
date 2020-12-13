package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Teil 1
	// Zahlen aus Datei in eine Liste um damit zu Rechnen
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var zahlen []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		z, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		zahlen = append(zahlen, z)
	}

	fmt.Println(zahlen)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Zahlen finden die zusammen 2020 ergeben
	var z1, z2 int
	for i := 0; i < len(zahlen); i++ {
		for j := 1; j < len(zahlen)-1; j++ {
			if zahlen[i]+zahlen[j] == 2020 {
				z1, z2 = zahlen[i], zahlen[j]
				break
			}
		}
	}
	fmt.Println(z1, " + ", z2, " = ", z1+z2)
	fmt.Println(z1, " * ", z2, " = ", z1*z2)

	// Teil 2
	// Zahlen finden die zusammen 2020 ergeben
	var z3 int
	for i := 0; i < len(zahlen); i++ {
		for j := 1; j < len(zahlen)-1; j++ {
			for k := 1; k < len(zahlen)-2; k++ {
				if zahlen[i]+zahlen[j]+zahlen[k] == 2020 {
					z1, z2, z3 = zahlen[i], zahlen[j], zahlen[k]
					break
				}
			}
		}
	}
	fmt.Println(z1, " + ", z2, " + ", z3, " = ", z1+z2+z3)
	fmt.Println(z1, " * ", z2, " * ", z3, " = ", z1*z2*z3)
}
