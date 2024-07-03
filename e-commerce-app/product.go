package main

type Product struct {
    Name        string
    Price       float64
    Description string
}

func (p *Product) GetName() string {
    return p.Name
}

func (p *Product) GetPrice() float64 {
    return p.Price
}

func (p *Product) GetDescription() string {
    return p.Description
}
