package main

import "fmt"

type person struct {
	Name string
	Age  uint
}

func main() {

	// a := person{
	// 	Name: "Jayanth",
	// 	Age:  25,
	// }
	// b := &person{
	// 	Name: "Jayanth",
	// 	Age:  25,
	// }

	// fmt.Println("a : ", a)
	// fmt.Println("b :", b.Name)
	// fmt.Println("a==*b :", a == *b)

	var f map[int]int

	f = make(map[int]int)
	for i := 0; i < 10; i++ {
		f[i] = i * i
	}
	fmt.Println("f :", f)

	q := map[int]int{}

	for i := 0; i < 10; i++ {
		q[i] = i + 1
	}
	fmt.Println("q :", q)
	change(q)
	fmt.Println("q :", q)

	fmt.Println("---------------------")
	fmt.Println("---------------------")

}

func change(d map[int]int) {

	for i := 0; i < 10; i++ {
		d[i] = d[i] + 3

	}

}
