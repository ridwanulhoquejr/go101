package practices

import (
	"fmt"
)

// interface
type Speaker interface {
	Speak()
}

// struct
type Person struct {
	Name string
}

type User struct {
	Person
	Id string
}

type Robot struct {
	Id string
}

// interface implementations
// user struct
func (u *User) Speak() {
	fmt.Println("Hello i am User, implementing Speaker interface!")
}

// Robot struct
func (r *Robot) Speak() {
	fmt.Println("Hello i am Robot, implementing Speaker interface!")
}

func checkInterface(s Speaker) {
	// assertion
	fmt.Printf("I am in checkInterface and my type is %T\n", s)

	// now call its own implementation
	s.Speak()

	// switch assert
	switch v := s.(type) {
	case *User:
		fmt.Printf("Hello i am switch assertion and my type %T", v)
	case *Robot:
		fmt.Printf("Hello i am switch assertion and my type %T", v)
	default:
		fmt.Printf("Hello i am switch assertion and i am default type")
	}
}

func main() {

	var i interface{} = 42
	if v, ok := i.(int); ok {
		fmt.Printf("I am interface assert and my type is %T and value %v", i, v)
	}

	//
	user := &User{
		Person: Person{
			Name: "Ridwan",
		},
		Id: "U-1",
	}

	robot := &Robot{
		Id: "R-1",
	}

	var s Speaker = user
	if v, ok := s.(*User); ok {
		fmt.Printf("i am type assertion and my type is %T", v)
	}

	checkInterface(user)
	checkInterface(robot)

}
