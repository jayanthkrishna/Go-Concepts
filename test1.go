package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
}

type employee struct {
	person
	company string
}

func (p person) speak() {
	fmt.Println("Hi My name is ", p.fname, p.lname)
	// p.fname = "Krishna"
	// fmt.Println("p val in speak method :", p.fname)
}
func (p *person) talk() {
	p.fname = "Krishna"
	fmt.Println("p val in talk method :", p.fname)
	fmt.Println("p in talk method :", *p)
}
func main() {
	var a int
	var b *int
	var arr1 = [5]int{1, 3, 4, 6, 7}
	// var arr2 = [5]int{5, 3, 5, 87, 8}

	fmt.Printf("a : %d\n", a)
	fmt.Printf("a address in main: %v\n", &a)
	c := f1(a)
	fmt.Printf("c address in main: %v\n", &c)
	fmt.Printf("a : %d\n", a)
	fmt.Printf("a val : %v\n", &a)
	b = f2(&a)
	fmt.Printf("a val :%v\n", a)
	fmt.Printf("b address in main: %v\n", b)
	fmt.Printf("b val in main: %v\n", *b)
	fmt.Printf("b's pointer address in main: %v\n", &b)
	fmt.Println("-----------------------------------------")
	fmt.Println("Arr1 Address ", &arr1[0], &arr1[1])
	f3(&arr1)

	fmt.Printf("Arr1 values after f3: %v\n", arr1)
	fmt.Println("-----------------------------------------")
	fmt.Println("-----------------------------------------")
	fmt.Println("-----------------------------------------")

	p := person{
		fname: "Jayanth", lname: "Yadav", age: 24,
	}

	e := &employee{
		person:  p,
		company: "tcs",
	}
	fmt.Println("e val ", e.person)
	e.speak()
	// fmt.Println("p val after speak method :", p.fname)
	// p.talk()
	// fmt.Println("p val after talk method :", p.fname)
}

func f1(inp int) int {

	inp = inp + 1
	fmt.Printf("a's value in f1 : %v\n", inp)
	fmt.Printf("a address in f1: %v\n", &inp)
	return inp
}

func f2(inp *int) *int {
	*inp = *inp + 2
	fmt.Printf("a's value in f2 : %v\n", *inp)
	fmt.Printf("a's address in f2 : %v\n", inp)
	fmt.Printf("a's pointer address in f2 : %v\n", &inp)

	return inp
}

func f3(inp *[5]int) {
	inp[0] = -12
	inp[1] = -32
	fmt.Println("inp val", inp)
	fmt.Println("inp dereferenced val", *inp)
	fmt.Println("inp Address", &inp)

}
