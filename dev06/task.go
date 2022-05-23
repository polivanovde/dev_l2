package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	input, deli, fields, sep := args(os.Args)

	if sep && !strings.Contains(input, deli) {
		log.Fatalf("input not contains '%s'", deli)
	}

	stSl := strings.Split(input, deli)

	for k, col := range stSl {
		if fields > 0 && k == fields {
			fmt.Println(col)
		} else if fields == 0 {
			fmt.Println(col)
		}
	}
}

func args(args []string) (input, deli string, fields int, sep bool) {
	if len(args) < 1 {
		log.Fatal("usage: go run . <file_path> <pattern>")
	}

	deli = "\t"

	for i := 1; i < len(args); i++ {
		if i == 1 {
			input = args[i]
		} else {
			switch args[i] {
			case "-f":
				if len(args) > i+1 {
					fields, _ = strconv.Atoi(args[i+1])
				}
			case "-d":
				if len(args) > i+1 {
					deli = args[i+1]
				}
			case "-s":
				sep = true
			}
		}
	}

	return
}
