package main

import (
	"fmt"
	"strings"
)

func ePrimo(num int) bool {
	if num <= 1 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	for i := 3; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	n1 := 3
	n2 := 3

	fmt.Println("A soma dos numeros e: ", n1+n2)
	fmt.Println("A diferença dos numeros e: ", n1-n2)

	text := "tamanho"

	fmt.Println("O tamanho da string e ", len(text))

	num := 100

	for i := 0; i <= num; i++ {
		numPrimo := ePrimo(i)

		if numPrimo {
			fmt.Printf("O numero (%d) é primo", i)
			fmt.Println("")
		}
	}

	text2 := "Este texto contém cinco palavras"
	totalSpace := strings.Count(text2, " ") + 1	

	fmt.Printf("O total de palavras é %d", totalSpace)
	fmt.Println("")

	text3 := "este é um texto a ser capitalizado"

	text3Capitalizado := strings.ToUpper(text3)

	fmt.Println("A soma dos numeros e: ", text3Capitalizado)

	value := 2
	exp := 8
	total := value

	for i := 1; i < exp; i++{
		total = total * value		
	}
	
	fmt.Printf("%d elevado a %d é: %d", value, exp, total)
	fmt.Println()

}
