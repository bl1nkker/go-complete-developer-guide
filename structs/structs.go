package structs

import "fmt"

type ContactInfo struct{
	Email string
	Number string
}
type Person struct{
	FirstName string
	LastName string
	ContactInfo
}

func (p Person) Print(){
	fmt.Println("Printing person...")
	fmt.Printf("%+v\n", p)
}


// !Go is pass by value language.
func (p *Person) UpdateName(newName string){
	fmt.Println("Pointer", p)
	p.FirstName = newName
}

func Run(){
	fmt.Println("Structs module")
	// alex := structs.Person{FirstName: "Alex", LastName: "Anderson"}
	var alex Person
	alexContact := ContactInfo{Email: "alex@gmail.com", Number: "111333"}
	alex.FirstName = "Alex"
	alex.LastName = "Anderson"
	alex.ContactInfo = alexContact
	alex.Print()

	jim := Person{
		FirstName: "Jim",
		LastName: "Fatherbeard",
		ContactInfo: ContactInfo{
			Email: "jim@gmail.com",
			Number: "228337",
		},
	}
	jim.UpdateName("Fuck")
	jim.Print()
	
}