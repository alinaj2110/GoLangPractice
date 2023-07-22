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
	fmt.Printf("%+v", sample)
}
