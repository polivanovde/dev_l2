package main

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	ps "github.com/mitchellh/go-ps"
)

func walk(path string) string {
	var out string
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		out += fmt.Sprintf("%s\n", file.Name())
	}

	return out
}

func pwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return pwd
}

func read() []string {
	in, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(strings.TrimSuffix(in, "\n"), " | ")
}

func pList() string {
	var out string
	processList, err := ps.Processes()
	if err != nil {
		log.Fatal("ps.Processes() Failed, are you using windows?")
	}

	for x := range processList {
		var process ps.Process
		process = processList[x]
		out += fmt.Sprintf("%d\t%s\n", process.Pid(), process.Executable())
	}

	return out
}

func execCommand(input []string) {
	var (
		pth    string = "."
		p, out string
	)
	for _, comm := range input {
		quoted := false
		c := strings.FieldsFunc(comm, func(r rune) bool {
			if r == '"' {
				quoted = !quoted
			}
			return !quoted && r == ' '
		})
		switch c[0] {
		case "cd":
			var fullPath string
			if len(c) > 1 {
				p = c[1]
			}
			if !strings.HasPrefix(p, "/") {
				fullPath = filepath.Join(pwd(), p)
			} else {
				fullPath = p
			}
			err := os.Chdir(fullPath)
			if err != nil {
				log.Fatal(err)
			}
		case "ls":
			out = walk(pth)
		case "pwd":
			out = pwd() + "\n"
		case "echo":
			if len(c) > 1 {
				out = c[1] + out
			}
		case "kill":
			var killPid string
			if len(c) > 1 {
				killPid = c[1]
			} else {
				killPid = out
			}

			pids := strings.Split(killPid, "\n")
			for _, pID := range pids {
				pid, err := strconv.Atoi(pID)
				if err != nil {
					log.Fatal(err)
				}

				proc, err := os.FindProcess(pid)
				if err != nil {
					log.Fatal(err)
				}

				err = proc.Kill()
				if err != nil {
					fmt.Println(err)
				}
			}
			out = ""
		case "ps":
			out = pList()
		default:
			cmd := exec.Command("bash", "-c", strings.Join(c, " "))
			cmd.Stdin = strings.NewReader(out)
			var stout bytes.Buffer
			cmd.Stdout = &stout
			cmd.Dir = pth
			err := cmd.Run()
			if err != nil {
				log.Println(err)
			}
			out = stout.String()
		}
	}
	fmt.Print(out)
	fmt.Printf("%s: ", pwd())
	execCommand(read())
}

func main() {
	fmt.Printf("%s: ", pwd())
	var input []string = read()

	execCommand(input)
}
