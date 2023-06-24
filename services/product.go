package product

import (
	"fmt"
)

type Product struct {
	id          uint32
	name        string
	description string
	price       float32
}

func (p *Product) SetName(name string) {

	p.name = name
}

func (p *Product) SetDescription(description string) {

	p.description = description
}

func (p *Product) SetPrice(price float32) {

	p.price = price
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) Description() string {
	return p.description
}

func (p *Product) Price() float32 {
	return p.price
}

func (p Product) Store() {

	fmt.Printf("%+v", p)
}

func Index() {

}
