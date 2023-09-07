package product

import (
	"encoding/json"
	"os"
)

var Products []Product

func getAllProducts(path string) []Product {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(file, &Products)
	if err != nil {
		return nil
	}
	return Products
}
