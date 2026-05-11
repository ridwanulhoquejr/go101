package main

import (
	"fmt"
)

// ============================================================
// TOPIC: Structs - Composition, Embedding, Nested Structs
// ============================================================

// --- EXAMPLE ---

type Address struct {
	City    string
	Country string
}

type Person struct {
	Name    string
	Age     int
	Address Address // nested struct
}

type Employee struct {
	Person // embedded struct (field promotion)
	Role   string
	Salary int
}

type User struct {
	Name    string
	Age     string
	Address string
}

// here in reveiver method the User object is copied
func (u *User) SetName(name string) error {

	if name == "" {
		return fmt.Errorf("Name is empty")
	}
	fmt.Printf("memeory address of user in rec method %p\n", u)

	fmt.Printf("initial name %v\n", u.Name)

	u.Name = name // override

	fmt.Printf("updated name %v\n", u.Name)

	return nil
}

func main() {
	// Struct literal

	// user := &User{
	// 	Name:    "Ridwan",
	// 	Age:     "12",
	// 	Address: "Chattogram",
	// }

	// user.SetName("updated name")
	// fmt.Printf("User name after update %v\n", user.Name)
	// fmt.Printf("memeory address of user %p", user)

	// pointer
	x := "42"

	p := &x // p = address of x
	fmt.Printf("P holds an address of x %p", p)

	// dereference: if we need to know about the value of a pointer variable (& address of)
	fmt.Printf("*P will prints the value of x %v", *p)

	// declation: pointer to
	var y *int // y is a pointer to an int; zero value is nil

	fmt.Printf("Y is %v (a declared *int has zero value nil)\n", y)
	// likewise
	// func (u *User) -> u is a pointer to a User

	// person := Person{
	// 	Name: "Rob",
	// 	Age:  45,
	// 	Address: Address{
	// 		City:    "San Francisco",
	// 		Country: "USA",
	// 	},
	// }
	// fmt.Printf("Person: %+v\n", person)
	// fmt.Println("City:", person.Address.City)

	// // Struct embedding - fields are promoted
	// employee := Employee{
	// 	Person: person,
	// 	Role:   "Engineer",
	// 	Salary: 160000,
	// }

	// // Access promoted fields directly
	// fmt.Println("Name:", employee.Name)         // promoted from Person
	// fmt.Println("City:", employee.Address.City) // promoted through Person
	// fmt.Printf("Full: %s, %s, $%d\n", employee.Name, employee.Role, employee.Salary)

	// // Structs are value types (copied on assignment)
	// p2 := person
	// p2.Name = "Alice"
	// fmt.Println("Original:", person.Name) // still "Rob"
	// fmt.Println("Copy:", p2.Name)         // "Alice"
}

// ============================================================
// PRACTICE QUESTIONS
// ============================================================
//
// Q1: Create a `Book` struct with Title, Author, and Pages.
//     Create an instance and print it.
//
// Q2: Create a `Library` struct that has a Name and a slice
//     of Books ([]Book). Add 3 books and loop through them.
//
// Q3: Create a `Student` struct that embeds `Person` and adds
//     a `GPA` field. Show that you can access Name directly.
//
// Q4: Write a function `newPerson(name string, age int) *Person`
//     that returns a pointer to a new Person. Why might you
//     return a pointer instead of a value?
//
// Q5: What happens when you compare two structs with ==?
//     Try comparing two Person values. Does it work?
//     (Hint: it works if all fields are comparable)
//
// Q6: Create an anonymous struct inline:
//     config := struct{ Host string; Port int }{"localhost", 8080}
//
// ============================================================
