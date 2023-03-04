package main

import "fmt"

type Person struct {
	name      string
	age       int
	netIncome int
	debt      int
}

type Employee struct {
	Person
	salary int
}

func (p Person) NetWorth() int {
	return p.netIncome - p.debt
}

func main() {
	fmt.Println("Hello World!")
	p := Person{name: "John", age: 30, netIncome: 50000, debt: 10000}
	e := Employee{Person: p, salary: 50000}
	fmt.Println(e.Person.NetWorth())
}
