package structs
import "fmt"

type ContactInfo struct{
	Email string
	Number string
}
type Person struct{
	FirstName string
	LastName string
	Contact ContactInfo
}

func (p *Person) Print(){
	fmt.Println("Printing person...")
	fmt.Printf("%+v\n", p)
}