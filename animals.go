package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Animal struct {
	food 		string
	locomotion 	string
	noise 		string
}


func (a *Animal) Eat(){
	fmt.Println("Food eaten: ",a.food)
}

func (a *Animal) Move(){
	fmt.Println("Locomotion method: ",a.locomotion)
}

func (a *Animal) Speak(){
	fmt.Println("Spoken sound: ", a.noise)
}

func responder(animal *Animal, pregunta string){
	switch pregunta{
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Enter a valid request")
	}
}


func main(){
	cow := Animal{food:"grass",locomotion:"walk",noise:"moo"}
	bird := Animal{food: "worms",locomotion: "fly",noise: "peep"}
	snake := Animal{food: "mice",locomotion: "slither",noise: "hsss"}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">")
		if scanner.Scan(){
			res := strings.Split(scanner.Text(), " ")
			if res[0] == "cow"{
				responder(&cow,res[1])
			}else if res[0] == "bird"{
				responder(&bird, res[1])
			}else if res[0] == "snake"{
				responder(&snake, res[1])
			}
		}

	}
}