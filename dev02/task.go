package main

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
    "fmt"
    "strings"
    "unicode"
    "strconv"
)

//UnPack is switch string elements
func UnPack(s string) string {
    var (
        newS string
        prev rune
        skip bool
    )
    rs := []rune(s)
    for _,cur := range rs {
        if unicode.IsDigit(cur) && !skip {
            num, err := strconv.Atoi(string(cur))
            if err == nil {
                newS += strings.Repeat(string(prev), num-1)
            }
        } else if cur == 92 && !skip {
            skip = true
        } else {
            skip = false
            prev = cur
            newS += string(cur)
        }
    }

    return newS
}

func main() {
    var in string
    fmt.Scan(&in)
    fmt.Println(UnPack(in))
}
