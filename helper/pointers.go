package helper

import "fmt"

func printPointer(namePointer *string) {
    fmt.Println("C: ", namePointer) // C:  0xc0000101b0   --> value  (Go copy it from A)
    fmt.Println("D: ", &namePointer) // D:  0xc00007c018   --> address of C
}


func Run(){
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

	i, j := 1, 1
	fmt.Printf("i is %v, j is %v\n", &i, &j)

	name := "Bill"
	namePointer := &name

    fmt.Println("A: ", namePointer) // A:  0xc0000101b0   --> value
    fmt.Println("B: ", &namePointer) // B:  0xc00000e018   --> address of A
    printPointer(namePointer)	
}

func updateSlice(s []string) {
	s[0] = "Fuck the"
}

func PointersGotcha(){
	slice := []string{"Hello", "World", "asdadawdasdsa"}
	fmt.Println("Before:", slice)
	updateSlice(slice)
	fmt.Println("After:", slice)
}