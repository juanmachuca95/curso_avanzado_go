package main

import "fmt"

/*
Patron de diseño factory

Metodo creacional
*/

type IProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

type Computer struct {
	name  string
	stock int
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c *Computer) getStock() int {
	return c.stock
}

func (c *Computer) getName() string {
	return c.name
}

type Laptop struct {
	Computer
}

type Desktop struct {
	Computer
}

func newLaptop() IProduct {
	return &Laptop{
		Computer: Computer{
			name:  "Laptop one",
			stock: 25,
		},
	}
}

func newDesktop() IProduct {
	return &Desktop{
		Computer: Computer{
			name:  "Desktop one",
			stock: 35,
		},
	}
}

func printNameIProduct(p IProduct) {
	fmt.Printf("Product name: %s - stock: %d\n", p.getName(), p.getStock())
}

func GetComputerFactory(computer string) (IProduct, error) {
	if computer == "laptop" {
		return newLaptop(), nil
	}
	if computer == "desktop" {
		return newDesktop(), nil
	}

	return nil, fmt.Errorf("Invalid computer")
}

func main() {
	laptop, _ := GetComputerFactory("laptop")
	desktop, _ := GetComputerFactory("desktop")

	printNameIProduct(laptop)
	printNameIProduct(desktop)
}
