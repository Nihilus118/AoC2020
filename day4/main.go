package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func validate(daten string) bool {
	// Erstmal eine Map aus dem String erstellen um sinnvoll damit zu arbeiten
	split := strings.Split(daten, " ")
	sort.Strings(split)

	var temp []string
	attribute := map[string]string{}

	for i := 0; i < len(split); i++ {
		if split[i] != "" {
			temp = strings.Split(split[i], ":")
			attribute[temp[0]] = temp[1]
		}
	}

	// Prüfen ob alle nötigen Schlüssel existieren und die Werte sinnvoll sind (Teil 2)
	if val, vorhanden := attribute["byr"]; vorhanden {
		jahr, err := strconv.Atoi(val)
		if err != nil {
			return false
		}
		if !(jahr >= 1920 && jahr <= 2002) {
			fmt.Println("byr falsch")
			return false
		}
	} else {
		return false
	}

	if val, vorhanden := attribute["iyr"]; vorhanden {
		jahr, err := strconv.Atoi(val)
		if err != nil {
			return false
		}
		if !(jahr >= 2010 && jahr <= 2020) {
			fmt.Println("iyr falsch")
			return false
		}
	} else {
		return false
	}

	if val, vorhanden := attribute["eyr"]; vorhanden {
		jahr, err := strconv.Atoi(val)
		if err != nil {
			return false
		}
		if !(jahr >= 2020 && jahr <= 2030) {
			fmt.Println("eyr falsch")
			return false
		}
	} else {
		return false
	}

	if val, vorhanden := attribute["hgt"]; vorhanden {
		temp := string(val[0 : len(val)-2])
		zahl, err := strconv.Atoi(temp)
		if err != nil {
			return false
		}

		if strings.HasSuffix(val, "in") {
			if !(zahl >= 59 && zahl <= 76) {
				fmt.Println("hgt falsch")
				return false
			}
		} else if strings.HasSuffix(val, "cm") {
			if !(zahl >= 150 && zahl <= 193) {
				fmt.Println("hgt falsch")
				return false
			}
		} else {
			fmt.Println("hgt einheit falsch")
			return false
		}
	} else {
		return false
	}

	if val, vorhanden := attribute["hcl"]; vorhanden {
		matched, err := regexp.MatchString(`^#[a-f0-9]{6}`, val)
		if err != nil {
			log.Fatal(err)
		}
		if !matched {
			fmt.Println("hcl falsch")
			return false
		}
	} else {
		return false
	}

	if val, vorhanden := attribute["ecl"]; vorhanden {
		if !contains(eyecolors, val) {
			fmt.Println("ecl falsch")
			return false
		}
	} else {
		return false
	}

	if val, vorhanden := attribute["pid"]; vorhanden {
		matched, err := regexp.MatchString(`^[0-9]{9}`, val)
		if err != nil {
			log.Fatal(err)
		}
		if !matched {
			fmt.Println("pid falsch")
			return false
		}
	} else {
		return false
	}

	return true
}

var eyecolors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

func main() {
	// Pässe aus Datei
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var paesse []string
	var temp string = ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err != nil {
			log.Fatal(err)
		}

		if scanner.Text() == "" {
			// Passdaten zuende, also in Liste
			paesse = append(paesse, temp)
			temp = ""
		} else {
			// Passdaten noch nicht am Ende also erstmal temporär in String
			temp += scanner.Text() + " "
		}
	}

	// Falls die Datei nicht mit einer Leerzeile endet darf der letzte Pass nicht vernachlässigt werden!
	if temp != "" {
		paesse = append(paesse, temp)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Jetzt alle Pässe prüfen
	valide := 0
	for _, pass := range paesse {
		if validate(pass) {
			valide++
		}
	}

	fmt.Println(valide)
}
