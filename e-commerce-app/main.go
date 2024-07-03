package main

func main() {
    electronicProduct := &ElectronicProduct{
        Product: Product{
            Name:        "Smartphone",
            Price:       799.99,
            Description: "A high-end smartphone with advanced features",
        },
        Warranty: 2,
    }

    clothingProduct := &ClothingProduct{
        Product: Product{
            Name:        "T-Shirt",
            Price:       29.99,
            Description: "A comfortable cotton t-shirt",
        },
        Size: "M",
    }

    println("Electronic Product:")
    println("Name:", electronicProduct.GetName())
    println("Price:", electronicProduct.GetPrice())
    println("Description:", electronicProduct.GetDescription())
    println("Warranty:", electronicProduct.GetWarranty(), "years")

    println("\nClothing Product:")
    println("Name:", clothingProduct.GetName())
    println("Price:", clothingProduct.GetPrice())
    println("Description:", clothingProduct.GetDescription())
    println("Size:", clothingProduct.GetSize())
}
