package main

import (
	"flag"
	"fmt"
	"justify/asciiArtTemplates"
	"log"
	"strings"
)

var alignOption string

func main() {
	align := flag.String("align", "left", "Aligns ascii-art on a console to the left, right, center, or justifies it.")
	option := flag.String("option", "standard", "Ascii-art option: standard, thinkertoy, or shadow.")
	flag.Parse()

	*align = strings.ToLower(*align)
	if *align != "left" && *align != "right" && *align != "center" && *align != "justify" {
		log.Fatalln("No such aligning option.")
	}
	alignOption = *align

	args := flag.Args()
	if len(args) != 1 {
		log.Fatalln("Usage: go run . [BANNERS] [STRING]")
	}
	asciiArtTemplates.ReadTemplates(&Store, *option, &cWidth)

	input := strings.Split(args[0], "\\n")
	removeNewline(&input)

	if alignOption == "justify" {
		justify(&input)
		return
	}

	for _, s := range input {
		if s == "" {
			continue
		}
		fmt.Print(PrintInput(s))
	}

}

var (
	Store  [128][8]string // Переменная для хранения символов из файла
	cWidth int
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

		switch alignOption {
		case "right":
			tmp = strings.Repeat(" ", cWidth-len(tmp)-1) + tmp
		case "center":
			tmp = strings.Repeat(" ", (cWidth-len(tmp))/2) + tmp
		}

		if len(tmp) > cWidth {
			log.Fatalln("Ascii-art doesn't fit into console window.")
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
