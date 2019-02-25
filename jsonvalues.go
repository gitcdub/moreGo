//How to read and decode JSON values from an input stream.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	const jsonStream = `
	{"Name": "John", "City": "Bangkok", "Salary": 51000}
	{"Name": "Mark", "City": "London", "Salary": 72430}
	{"Name": "Jane", "City": "Prague", "Salary": 46250}
	{"Name": "Haro", "City": "Tokyo", "Salary": 63460}
	`
	type Employee struct {
		Name, City string
		Salary     int32
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var emp Employee
		if err := dec.Decode(&emp); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s: %d\n", emp.Name, emp.City, emp.Salary)
	}
}
