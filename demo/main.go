package main

import (
	"fmt"

	"github.com/anton7r/bomlok/example"
)

func main() {
	// Create a Person instance
	person := &example.Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Email:     "john.doe@example.com",
	}

	// Use Bomlok interface methods
	fmt.Println("Person fields:", person.Bomlok_Fields())
	fmt.Println("FirstName:", person.Bomlok_GetValue("FirstName"))
	fmt.Println("Age:", person.Bomlok_GetValue("Age"))

	// Create an Address instance
	address := &example.Address{
		Street:  "123 Main St",
		City:    "New York",
		State:   "NY",
		ZipCode: "10001",
	}

	fmt.Println("\nAddress fields:", address.Bomlok_Fields())
	for _, field := range address.Bomlok_Fields() {
		fmt.Printf("%s: %v\n", field, address.Bomlok_GetValue(field))
	}

	// Create a Company instance
	company := &example.Company{
		Name:      "Tech Corp",
		Employees: 500,
		Revenue:   1000000.50,
		Public:    true,
	}

	fmt.Println("\nCompany fields:", company.Bomlok_Fields())
	for _, field := range company.Bomlok_Fields() {
		fmt.Printf("%s: %v\n", field, company.Bomlok_GetValue(field))
	}
}
