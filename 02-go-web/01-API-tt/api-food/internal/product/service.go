package product

import (
	"errors"
	"fmt"
	"time"
)

type Product struct {
	Id          int
	Name        string
	Quantity    int
	CodeValue   string
	IsPublished bool
	Expiration  string
	Price       float64
}

func validateProduct(product Product) (err error) {
	if product.Price == 0.0 {
		err = &EmptyError{Field: "precio"}
		return
	}
	if product.Name == "" {
		err = &EmptyError{Field: "nombre"}
		return
	}
	if product.Expiration == "" {
		err = &EmptyError{Field: "venicimiento"}
		return
	}
	if product.CodeValue == "" {
		err = &EmptyError{Field: "codigo"}
		return
	}
	if product.Quantity == 0 {
		err = &EmptyError{Field: "cantidad"}
		return
	}
	if err = validateUniqueCode(product); err != nil {
		return
	}
	if err = validateDate(product); err != nil {
		return
	}

	return nil
}

func validateUniqueCode(p Product) (err error) {
	for _, product := range Products {
		if product.CodeValue == p.CodeValue {
			err = errors.New("el código del producto debe ser unico")
			return
		}
	}
	return nil
}

func validateDate(p Product) (err error) {
	layout := "02/01/2006"
	_, err = time.Parse(layout, p.Expiration)
	if err != nil {
		return
	}
	return nil
}

type EmptyError struct {
	Field string
}

func (e *EmptyError) Error() string {
	return fmt.Sprintf("El %s del producto no debe estar vacio", e.Field)
}
