package asciiArtTemplates

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

// Эта функция запускается перед main и считывает файл в Store
func ReadTemplates(Store *[128][8]string, style string) {
	if !TxtFileCheck(style) {
		log.Fatalln("The file was changed.")
	}

	f, err := os.Open("asciiArtTemplates/" + style + ".txt")
	Check("Error opening file:", err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	order := " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~" // это порядок символов в файле, нужен чтобы правильно все считать в Store
	for _, r := range order {
		scanner.Scan() // Так как перед каждым символом в файле идет пустая строка, ее нужно пропускать, перед считыванием в Store
		for i := 0; i < 8; i++ {
			scanner.Scan()                    // здесь уже начинается считывание из файла в Store
			Store[int(r)][i] = scanner.Text() // Каждый символ помещается в ячейку под номером своего ascii значенения
		}
	}
}

// Проверяет ошибку, выводит ее на консоль с сопроводительным сообщеением в строке, если она есть
func Check(s string, e error) {
	if e != nil {
		log.Fatalln(s, e.Error())
	}
}

// Проверяет файл на какие либо изменения, и выдает ошибку, если они были
func TxtFileCheck(style string) bool {
	var hash string

	switch style {
	case "standard":
		hash = "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	case "shadow":
		hash = "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	case "thinkertoy":
		hash = "a57beec43fde6751ba1d30495b092658a064452f321e221d08c3ac34a9dc1294"
	default:
		log.Fatalln("No such style.")
	}

	file, err := os.Open("asciiArtTemplates/" + style + ".txt")
	Check("Error opening file:", err)
	defer file.Close()
	buf := make([]byte, 30*1024)
	sha256 := sha256.New()
	for {
		n, err := file.Read(buf)
		if n > 0 {
			_, err := sha256.Write(buf[:n])
			if err != nil {
				log.Fatal(err)
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Read %d bytes: %v", n, err)
			break
		}
	}
	sum := fmt.Sprintf("%x", sha256.Sum(nil))

	return sum == hash
}
