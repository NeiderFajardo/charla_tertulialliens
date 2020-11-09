package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
)

type name struct{
	fname string
	lname string
}

func verificartamanio(s string)string{
	const tam = 20
	if len(s) > tam{
		return s[:tam]
	}
	return s
}

func leerarchivo(nom string){
	personas := []name{}
	data, err := os.Open(nom)
	if err != nil{
		fmt.Println("File reading error")
		return
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan(){
		aux := strings.Split(scanner.Text(), " ") 
		p := name{fname: verificartamanio(aux[0]), lname: verificartamanio(aux[1])}
		personas = append(personas, p)
	}
	for _, persona := range personas{
		fmt.Println("First Name: ",persona.fname," Last Name: ",persona.lname)
	}
}

func main(){
	var nombre string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter the name of the file: ")
	if scanner.Scan(){
		nombre = scanner.Text()
	}
	leerarchivo(nombre)
}