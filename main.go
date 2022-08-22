package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"output/asciiArtTemplates"
	"strconv"
	"strings"
)

func main() {
	outputBanner := flag.String("output", "output.txt", "creates a file with ascii art.")
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		log.Fatalln("Usage: go run . [OPTIONS] [STRING] [STYLE]")
	}
	asciiArtTemplates.ReadTemplates(&Store, args[1])

	input := strings.Split(args[0], "\\n")
	removeNewline(&input)

	var outputString string

	for _, s := range input {
		if s == "" {
			outputString = outputString + "\n"
			continue
		}
		outputString = outputString + PrintInput(s)
	}

	os.Create(*outputBanner)
	ioutil.WriteFile(*outputBanner, []byte(outputString), os.FileMode(os.O_RDONLY))
}

var (
	Store [128][8]string // Переменная для хранения символов из файла
)

// Выводит данную строку на консоль символами из файла
func PrintInput(s string) (result string) {

	for i := 0; i < 8; i++ {
		var tmp string

		for _, r := range s {
			if r < 0 || r > 127 || Store[int(r)][0] == "" {
				fmt.Println("A character is not available.")
				return
			}

			tmp = tmp + Store[int(r)][i]
		}

		result = result + tmp + "\n"
	}

	return result
}

// Когда в вводной строке нет слов, а только '\n' либо вообще ничего, создается лишняя пустая строка, из-за которой на консоль выводится лишшняя новая линия. Эта функция убирает эту строку.
// "\n" -> "", ""
// "\nHello" -> "", "Hello"
func removeNewline(input *[]string) {
	nowords := true
	for _, s := range *input {
		if len(s) > 0 {
			nowords = false
			break
		}
	}
	if nowords {
		*input = (*input)[1:]
	}
}

func ConsoleWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	asciiArtTemplates.Check("Error measuring console size:", err)

	outStr := string(out)
	outStr = strings.TrimSpace(outStr)
	heightWidth := strings.Split(outStr, " ")
	width, err := strconv.Atoi(heightWidth[1])
	asciiArtTemplates.Check("Error measuring console size:", err)

	return width
}
