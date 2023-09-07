package product

import (
	"encoding/json"
	"os"
)

var Products []*Product

func GetAllProducts(path string) []*Product {
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

func Save(p Product, contr *ControllerProducts) {
	contr.Db = append(contr.Db, &p)
}
