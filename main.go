package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println()
	fmt.Println("Программа для:")
	fmt.Println("Резделение участка на макс. симетр. квадраты")
	fmt.Println("----------------------------------")
	fmt.Println()

	var side1, side2 int

	fmt.Print("Введите Высоту и Ширену: ")
	fmt.Scanln(&side1, &side2)

	min_side := separation(side1, side2)

	w_cube := side1 / min_side
	h_cube := side2 / min_side

	for h_cube > 0 {
		h_cube--

		repeatedString := strings.Repeat(" ☐ ", w_cube)
		fmt.Println(repeatedString)
	}

	fmt.Println()
	fmt.Println("----------------------------------")
	fmt.Print("Ответ:", min_side, "x", min_side, " в участках ", w_cube, "x", side2/min_side)

	var choose string

	fmt.Println()
	fmt.Println("Еще раз? д/н, y/n")
	fmt.Scanln(&choose)
	if choose == "д" || choose == "y" {
		main()
	}

}

func separation(w, h int) int {

	if w < h {
		big_side := h
		h = w
		w = big_side
		fmt.Println(w, "x", h)
	}

	if w%h == 0 {
		return h
	}

	return separation(w-w/h*h, h)

}
