package main

import(
	"fmt"
	"sync"
	"strconv"
)

type Chopstick struct{
	mux sync.Mutex
}

type Philosopher struct{
	num int
	comida int
	leftStick *Chopstick
	rightStick *Chopstick
}


func (p Philosopher) eat(ch chan<- int, wg *sync.WaitGroup) {	
		for p.comida < 3{
			if ok := pedirPermiso(ch, p.num); ok{
				p.leftStick.mux.Lock()
				p.rightStick.mux.Lock()
				numero := strconv.Itoa(p.num)
				fmt.Println("starting to eat ",numero)	
				p.leftStick.mux.Unlock()
				p.rightStick.mux.Unlock()
				fmt.Println("finishing eating ",numero)
			}
			p.comida = p.comida +1
		}
	wg.Done()
}

func pedirPermiso(ch chan<- int, n int) bool {
	ch <- n
	return true
}

func host(ch <-chan int) {
	for{
		<-ch
	}
}


func crearPhilosphes(numero int, sticks []*Chopstick) []Philosopher {
	var philosophes []Philosopher
	for i := 0; i < numero; i++{
		p := Philosopher{num:i+1, comida:0, leftStick:sticks[i], rightStick:sticks[(i+1)%numero]}
		philosophes = append(philosophes, p)
	}
	return philosophes
} 

func crearSticks(numero int) []*Chopstick {
	var csticks []*Chopstick
	for i := 0;i < numero;i++{
		s := new(Chopstick)
		csticks = append(csticks, s)
	}
	return csticks
}


func main(){
	var wg sync.WaitGroup
	philosophes := crearPhilosphes(5, crearSticks(5))
	canal := make(chan int, 2)

	wg.Add(5)
	go host(canal)
	for p := range philosophes{
		go philosophes[p].eat(canal, &wg)
	}
	wg.Wait()
}