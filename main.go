package main

import "fmt"

type listFunction func([]int) []int

func executar_estrategia(lista *[]int, estrategia listFunction) {
	*lista = estrategia(*lista)
}

func sortAscending(array []int) []int {
	var swapped bool = true
	var start int = 0
	var end int = len(array) - 1

	for swapped {
		swapped = false

		for i := start; i < end; i++ {
			if array[i] > array[i+1] {
				swap(&array[i], &array[i+1])
				swapped = true
			}
		}
		if !swapped {
			break
		}
		swapped = false
		end--

		for i := end - 1; i >= start; i-- {
			if array[i] > array[i+1] {
				swap(&array[i], &array[i+1])
				swapped = true
			}
		}
		start++
	}
	return array
}
  
func ordenarDecrescente(lista []int) []int {
	tamanho := len(lista)
	copia := make([]int, tamanho)
	copy(copia, lista)

	for i := 0; i < tamanho-1; i++ {
		for j := 0; j < tamanho-1-i; j++ {
			if copia[j] < copia[j+1] {
				copia[j], copia[j+1] = copia[j+1], copia[j]
			}
		}
	}
	return copia

}

func removeDuplicates(lista []int) []int {
	chavesVistas := make(map[int]bool)
	resultado := []int{}

	for _, item := range lista {
		if _, visto := chavesVistas[item]; !visto {
			chavesVistas[item] = true
			resultado = append(resultado, item)
		}
	}
	return resultado

}

func main() {
	primes := []int{11, 2, 2, 2, 5, 3, 3, 11, 5, 11, 13, 5, 2, 3, 11}
	executar_estrategia(&primes, removeDuplicates)
	fmt.Println(primes)

  lista1 := []int{11, 2, 2, 2, 5, 3, 3, 11, 5, 11, 13, 5, 2, 3, 11}
	executar_estrategia(&lista1, sortAscending)
	fmt.Println("Lista Crescente: ", lista1)
  
	lista2 := []int{11, 2, 2, 2, 5, 3, 3, 11, 5, 11, 13, 5, 2, 3, 11}
	executar_estrategia(&lista2, ordenarDecrescente)
	fmt.Println("Lista Decrescente: ", lista2)
}
