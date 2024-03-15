package main

import (
	"fmt"
	"go-complete-developer-guide/cards"
	"go-complete-developer-guide/structs"
)

// Cards
func main() {
	// runCards()
	// runStructs()
	// runPointers()
	pointersGotcha()
}

func runCards(){
	deck := cards.NewDeck()
	_, deck = cards.Deal(deck, 5)

	// I/O
	deck.SaveToFile("cards.txt")
	newD := cards.NewDeckFromFile("cards.txt")
	newD.Print()

	// Shuffle
	deck.Print()
	deck.Shuffle()
	deck.Print()
}

func runStructs(){
	fmt.Println("Structs module")
	// alex := structs.Person{FirstName: "Alex", LastName: "Anderson"}
	var alex structs.Person
	alexContact := structs.ContactInfo{Email: "alex@gmail.com", Number: "111333"}
	alex.FirstName = "Alex"
	alex.LastName = "Anderson"
	alex.ContactInfo = alexContact
	alex.Print()

	jim := structs.Person{
		FirstName: "Jim",
		LastName: "Fatherbeard",
		ContactInfo: structs.ContactInfo{
			Email: "jim@gmail.com",
			Number: "228337",
		},
	}
	jim.UpdateName("Fuck")
	jim.Print()
	
}

func runPointers(){
	var x, y int;
	fmt.Println(x == 0, x == y) // true true
	fmt.Println(&x == &y) // false
	fmt.Printf("x is %T, pointerX is %T\n", x, &x) // x is int, pointerX is *int

	var p *int // p is type of pointer of int
	fmt.Println(p == nil) // true

	p = &x
	fmt.Println(p, p == &x, x == *p) // <p_address> true true
	fmt.Printf("p is %T, pointerP is %T\n", p, &p) // p is *int, pointerP is **int

	var q **int
	fmt.Println(q == nil) // true
	q = &p
	fmt.Println(q, q == &p, p == *q) // <q_address> true true
	fmt.Printf("q is %T, pointerQ is %T\n", q, &q) // q is **int, pointerQ is ***int
}

func updateSlice(s []string) {
	s[0] = "Fuck the"
}

func pointersGotcha(){
	slice := []string{"Hello", "World", "asdadawdasdsa"}
	fmt.Println("Before:", slice)
	updateSlice(slice)
	fmt.Println("After:", slice)
}