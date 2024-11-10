package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Text struct {
	Content string
}

func (t *Text) textModifier() {
	result := &strings.Builder{}
	sum := 0
	for _, r := range t.Content {
		if unicode.IsDigit(r) {
			sum += int(r) - 48
		} else {
			if r == ' ' {
				result.WriteString(string(" "))
			} else if r == '-' {

			} else if r == '+' {
				result.WriteString(string("!"))
			} else {
				result.WriteString(string(r))
			}
		}
	}
	result.WriteString(" ")
	result.WriteString(string(sum))
	fmt.Println(result)
}

func main() {
	test()
	text := &Text{}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите строку:")
	for scanner.Scan() {
		text.Content = scanner.Text()
		text.textModifier()
	}
}

func test() {
	t := Text{"   1-2   +"}
	t.textModifier()
	t1 := Text{"генрих1  играет+2   л-июбит0школу"}
	t2 := Text{"Я ю-лбю-л голанг   всем сердцем+"}
	t1.textModifier()
	t2.textModifier()

	fmt.Println("N1")
	fmt.Println("Было  - генрих1  играет+2   л-июбит0школу")
	fmt.Println("Стало -", t1.Content)
	fmt.Println("Нужно - генрих играет! илюбитшколу 3")

	fmt.Println("N2")
	fmt.Println("Было  - Я ю-лбю-л голанг   всем сердцем+")
	fmt.Println("Стало -", t2.Content)
	fmt.Println("Нужно - Я люблю голанг всем сердцем!")
}
