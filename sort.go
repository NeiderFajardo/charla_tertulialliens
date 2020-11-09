package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func unir(numeros map[int][]int) {
	var la []int
	for _, i := range numeros {
		for _, j := range i {
			la = append(la, j)
		}
	}
	sort.Ints(la)
	fmt.Print("The entire sorted list: ")
	fmt.Println(la)
}

func organizar(nums []int, c chan<- []int, wg *sync.WaitGroup) {
	defer wg.Done()
	sort.Ints(nums)
	fmt.Println(nums)
	c <- nums
}

func main() {
	var temp []int
	var wg sync.WaitGroup
	const tamanio = 4
	c := make(chan []int, tamanio)

	numeros := make(map[int][]int)
	scanner := bufio.NewScanner(os.Stdin)
	wg.Add(4)
	fmt.Println("Please enter the numbers in one line, separate with single spaces:")
	if scanner.Scan() {
		linea := strings.Split(scanner.Text(), " ")
		for i, num := range linea {
			if aux, err := strconv.Atoi(num); err == nil {
				temp = numeros[i%tamanio]
				temp = append(temp, aux)
				numeros[i%tamanio] = temp
			}
		}
	}
	fmt.Print("Antes de organizar: ")
	fmt.Println(numeros)
	//The four goroutines to sorted the lists
	for j, arreglos := range numeros {
		go organizar(arreglos, c, &wg)
		x := <-c
		numeros[j] = x
	}
	fmt.Print("Después de organizar: ")
	fmt.Println(numeros)
	wg.Wait()
	//unir(numeros)
}
