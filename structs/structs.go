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

func (p *Person) UpdateName(newName string){
	p.FirstName = newName
}