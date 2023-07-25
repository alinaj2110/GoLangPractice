package main

import "fmt"

type contactInfo struct {
	email   string
	pincode int
}
type person struct {
	firstname string
	lastname  string
	contact   contactInfo
}

func main() {
	sample := person{
		firstname: "Anjali",
		lastname:  "Praveen",
		contact: contactInfo{
			email:   "anjalipraveen2110@gmail.com",
			pincode: 12345678,
		},
	}
	sample.updateName("Anne")
	sample.print()

	books := map[string]int{"book1": 2023, "book2": 2021, "book3": 2022}
	print(books)

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(fname string) {
	(*p).firstname = fname
}

func print(m map[string]int) {

	for k, v := range m {
		fmt.Println("Key:", k, " Value:", v)
	}

}
