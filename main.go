package main

import (
	"io/ioutil"
	"log"

	"github.com/syedakmall/go-protobuf-example/protos"
	"google.golang.org/protobuf/proto"
)

func main() {

	// write binary file

	max := protos.Employee{
		Id:     1001,
		Name:   "Max",
		Salary: 32000,
	}
	lewis := protos.Employee{
		Id:     1002,
		Name:   "Lewis",
		Salary: 32000,
	}
	employees := protos.Employees{}
	employees.Employees = append(employees.Employees, &lewis, &max)
	log.Printf("Employees : %v", employees.GetEmployees())

	out, err := proto.Marshal(&employees)

	if err != nil {
		log.Fatalf("Serialization error: %s", err.Error())
	}

	if err := ioutil.WriteFile("employees.bin", out, 0644); err != nil {
		log.Fatalf("Write File Error: %s ", err.Error())
	}
	log.Println("Write Success")

	in, err := ioutil.ReadFile("employees.bin")

	if err != nil {
		log.Fatalf("Read File Error: %s ", err.Error())
	}

	// read from binary file

	employees = protos.Employees{}

	err = proto.Unmarshal(in, &employees)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Employees : %v", employees.GetEmployees())
}
