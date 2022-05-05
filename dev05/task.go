package main

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, word, isRev, isIgn, isCnt, isFull, isNum, after, before := args(os.Args)
	var (
		tCh        = make(chan map[int]string)
		rCh        = make(chan map[int]string)
		total uint = 0
	)
	go readFile(file, tCh)
	go grep(word, tCh, rCh, isRev, isIgn, isFull, after, before)

	for k, r := range <-rCh {
		total++
		if isNum {
			fmt.Printf("%d:", k)
		}
		fmt.Println(r)
	}
	if isCnt {
		fmt.Printf("Count, %d\n", total)
	}

}

func args(args []string) (file, word string, isRev, isIgn, isCnt, isFull, isNum bool, after, before int) {
	if len(args) < 3 {
		log.Fatal("usage: go run . <file_path> <pattern>")
	}

	for i := 1; i < len(args); i++ {
		if i == 1 {
			file = args[i]
		} else if i == 2 {
			word = args[i]
		} else {
			switch args[i] {
			case "-v":
				isRev = true
			case "-i":
				isIgn = true
			case "-c":
				isCnt = true
			case "-F":
				isFull = true
			case "-n":
				isNum = true
			}
			if strings.HasPrefix(args[i], "-A") {
				a, err := strconv.Atoi(strings.Trim(args[i], "-A"))
				if err != nil {
					log.Fatalln("no nums fund in -A")
				}
				after = a
			}
			if strings.HasPrefix(args[i], "-B") {
				b, err := strconv.Atoi(strings.Trim(args[i], "-B"))
				if err != nil {
					log.Fatalln("no nums fund in -B")
				}
				before = b
			}
			if strings.HasPrefix(args[i], "-C") {
				c, err := strconv.Atoi(strings.Trim(args[i], "-C"))
				if err != nil {
					log.Fatalln("no nums fund in -C")
				}
				after = c
				before = c
			}
		}
	}

	return
}

func readFile(file string, to chan<- map[int]string) {
	var (
		m      = make(map[int]string)
		rN int = 1
	)

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer close(to)
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		m[rN] = s.Text()
		rN++
	}
	to <- m
}

func grep(word string, from <-chan map[int]string, result chan<- map[int]string, isRev, isIgn, isFull bool, after, before int) {
	var (
		m                      = make(map[int]string)
		fromMap map[int]string = <-from
	)

	defer close(result)

	for key, line := range fromMap {
		func(l string) {
			var (
				searchL string
				searchW string
			)
			searchL = strings.TrimSuffix(l, "\n")
			searchW = word
			if isIgn {
				searchL = strings.ToLower(l)
				searchW = strings.ToLower(word)
			}

			if !isRev && !isFull && strings.Contains(searchL, searchW) {
				m[key] = l
			} else if isRev && ((!isFull && !strings.Contains(searchL, searchW)) || (isFull && searchL != searchW)) {
				m[key] = l
			} else if !isRev && isFull && searchL == searchW {
				m[key] = l
			}

			if before > 0 && fromMap[key-1] != "" && m[key] != "" {
				m[key-1] = fromMap[key-1]
			}
			if after > 0 && fromMap[key+1] != "" && m[key] != "" {
				m[key+1] = fromMap[key+1]
			}
		}(string(line))
	}

	result <- m
}
