package main

type ClothingProduct struct {
    Product
    Size string
}

func (cp *ClothingProduct) GetSize() string {
    return cp.Size
}
