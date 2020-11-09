package main

import(
	"fmt"
	"strconv"
	"sync"
)

/*
A Race Condition occurs when two or more goroutines can access to the same data and chage it at the same time.
In other words, race condition occurs when two parts of code that run concurrently access and modify shared data,
so inconsistencies can be generated.

To appreciate the race conditions, run the program with the following instruction:

go run -race racecondition.go
*/


func sumar(wg *sync.WaitGroup) int {
	i := 0
	/*
	In this code there are two goroutins that access the variable i at the same time and modify it,
	this makes the value of i different in the second goroutine for different executions of the code.
	*/
	go func() {
		for j := 0; j<10;j++{
			fmt.Println(i)
			i++
		}
		wg.Done()
	}()
	go func() {
		i = i+2
		fmt.Println("Goroutine 2: ",strconv.Itoa(i))
		wg.Done()
	}()
	return i
}

//Posiblemente, el hecho que el main termine su ejecución impide que se observe que las otras goroutine se están ejecutando
func main() {
//If you run this program the output will be 0, because the return of the sum() function is executed before that the two goroutines
	var wg sync.WaitGroup
	wg.Add(2)
	res := sumar(&wg)
	wg.Wait()
	fmt.Println("Result: ",strconv.Itoa(res))

}