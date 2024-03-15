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