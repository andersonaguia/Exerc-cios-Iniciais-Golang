package main

import (
	"fmt"
)

func ePrimo(num int)bool{
	if num <=1 {
		return false
	}
	if num == 2{
		return true
	}
	if num % 2 == 0{
		return false
	}
	for i := 3; i * i <= num; i++ {
		if num % i == 0 {
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
}
