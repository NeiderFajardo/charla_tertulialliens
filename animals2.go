package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
)


type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food 		string
	locomotion 	string
	noise 		string
}

type Bird struct {
	food 		string
	locomotion 	string
	noise 		string
}

type Snake struct {
	food 		string
	locomotion 	string
	noise 		string
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (c Cow) Move(){
	fmt.Println(c.locomotion)
}

func (b Bird) Move(){
	fmt.Println(b.locomotion)
}

func (s Snake) Move(){
	fmt.Println(s.locomotion)
}

func (c Cow) Speak(){
	fmt.Println(c.noise)
}

func (b Bird) Speak(){
	fmt.Println(b.noise)
}

func (s Snake) Speak(){
	fmt.Println(s.noise)
}

func CrearAnimal(animales map[string]Animal, name, tipo string) string {
	var animal Animal
	switch strings.ToLower(tipo){
	case "cow":
		animal = Cow{food:"grass",locomotion:"walk",noise:"moo"}
	case "bird":
		animal = Bird{food: "worms",locomotion: "fly",noise: "peep"}
	case "snake":
		animal = Snake{food: "mice",locomotion: "slither",noise: "hsss"}
	default:
		return "Enter a valid animal type"
	}
	animales[name] = animal
	return "Created it!"
}

func ProcesarQuery(animales map[string]Animal, name, accion string) {
	if an, ok := animales[name]; ok{
		switch strings.ToLower(accion){
		case "eat":
			an.Eat()
		case "move":
			an.Move()
		case "speak":
			an.Speak()
		default:
			fmt.Println("Enter a valid query")
		}
	} else {
		fmt.Println("Animal not found")
	}
}


func main(){
	animales :=make(map[string]Animal)
	scanner := bufio.NewScanner(os.Stdin)

	for{
		fmt.Print(">")
		if scanner.Scan(){
			request := strings.Split(scanner.Text(), " ")
			if request[0] == "newanimal"{
				fmt.Println(CrearAnimal(animales, request[1], request[2]))
			}else if request[0] == "query"{
				ProcesarQuery(animales, request[1], request[2])
			}
		}
	}

}