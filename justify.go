package main

import (
	"fmt"
	"log"
	"strings"
)

func justify(input *[]string) {
	inputWords := splitToWords(input)

	for _, words := range inputWords {
		if len(words) == 0 {
			fmt.Println()
			continue
		}

		var spaces string
		if len(words) > 1 {
			spaces = getSpaces(&words)
		}

		for i := 0; i < 8; i++ {
			var tmp string

			for j, word := range words {
				for _, r := range word {
					if r < 0 || r > 127 || Store[int(r)][0] == "" {
						fmt.Println("A character is not available.")
						return
					}

					tmp = tmp + Store[int(r)][i]
				}

				if j < len(words)-1 {
					tmp = tmp + spaces
				}
			}

			fmt.Println(tmp)
		}
	}
}

func splitToWords(input *[]string) [][]string {
	var tmp [][]string
	for _, s := range *input {
		if onlySpaces(s) {
			tmp = append(tmp, []string{s})
			continue
		}

		tmp = append(tmp, strings.Fields(s))
	}

	return tmp
}

func getSpaces(words *[]string) (spaces string) {
	var tmp string
	for _, word := range *words {
		for _, r := range word {
			tmp = tmp + Store[int(r)][0]
		}
	}

	if len(tmp) > cWidth {
		log.Fatalln("Ascii-art doesn't fit into console window.")
	}

	count := (cWidth - len(tmp)) / (len(*words) - 1)

	spaces = strings.Repeat(" ", count)
	return spaces
}

func onlySpaces(s string) (onlySpaces bool) {
	onlySpaces = true
	for _, r := range s {
		if r != ' ' {
			onlySpaces = false
			break
		}
	}

	if len(s) == 0 {
		onlySpaces = false
	}

	return onlySpaces
}
