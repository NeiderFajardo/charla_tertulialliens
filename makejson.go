package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func mostrarjson(m map[string]interface{}) {
	if pjson, err := json.MarshalIndent(m, "", "	"); err == nil {
		fmt.Println(string(pjson))
	}

	qjson, _ := json.Marshal(m)
	fmt.Println(string(qjson))
}

func main() {
	q := make(map[string]interface{})
	var name, addr string
	var age int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter a name: ")
	if scanner.Scan() {
		name = scanner.Text()
	}
	fmt.Print("Please enter an address: ")
	if scanner.Scan() {
		addr = scanner.Text()
	}
	fmt.Print("Please enter your age: ")
	if scanner.Scan() {
		age, _ = strconv.Atoi(scanner.Text())
	}
	q["address"] = addr
	q["name"] = name
	q["age"] = age

	mostrarjson(q)

}
