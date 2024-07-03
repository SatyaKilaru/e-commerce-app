package main

type ElectronicProduct struct {
    Product
    Warranty int
}

func (ep *ElectronicProduct) GetWarranty() int {
    return ep.Warranty
}

