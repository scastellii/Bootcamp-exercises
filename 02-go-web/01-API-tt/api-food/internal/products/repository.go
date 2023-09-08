package products

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-bases/02-go-web/01-API-tt/api-food/internal/domain"
	"os"
)

func NewRepositoryController(path string, lastId int) *RepositoryController {
	cp := FillProducts(path, lastId)
	return cp
}

// RepositoryController is an struct that represents a controller for allProducts
// exposing methods to handle allProducts
type RepositoryController struct {
	Db     []*domain.Product
	LastId int
}

type Repository interface {
	GetAll() ([]domain.Product, error)
	GetByPrice(price float64) ([]domain.Product, error)
	GetById(id int) (domain.Product, error)
	Store(product *domain.Product) (domain.Product, error)
	Update(id int, product *domain.Product) (domain.Product, error)
	Delete(id int) (domain.Product, error)
}

func (c *RepositoryController) GetAll() ([]domain.Product, error) {
	var result []domain.Product
	for _, product := range c.Db {
		result = append(result, *product) //TODO si elimino con el delete, al volver a consultar por todos los prductos me devuelve aca un null pointer.
		//Pareciera que c.db por más de eliminar mantiene el tamaño y la que elimina queda en nil.
		//Ver si estoy eliminando mal o si se puede redimensionar el slice.
	}
	return result, nil
}

func (c *RepositoryController) GetByPrice(price float64) (result []domain.Product, err error) {
	for _, product := range c.Db {
		if product.Price > price {
			result = append(result, *product)
		}
	}
	if len(result) == 0 {
		err = errors.New("no se encontraron productos con esas caracteristicas")
	}
	return result, err
}

func (c *RepositoryController) GetById(id int) (prd domain.Product, err error) {
	for _, product := range c.Db {
		if product.Id == id {
			return *product, nil
		}
	}
	err = errors.New("No se encontró un producto con id " + string(id))
	return
}

func (c *RepositoryController) Store(product *domain.Product) (prd domain.Product, err error) {
	err = c.validateUniqueCode(product)
	if err != nil {
		return
	}
	product.Id = c.LastId + 1
	c.Db = append(c.Db, product)
	c.LastId++
	fmt.Println("Se guardo con exito el producto", product)
	prd = *product
	return prd, nil
}

func (c *RepositoryController) Update(id int, product *domain.Product) (prod domain.Product, err error) {
	for _, prd := range c.Db {
		if prd.Id == id {
			err = c.validateUniqueCode(product)
			if err != nil {
				return
			}
			prd.CodeValue = product.CodeValue
			prd.Price = product.Price
			prd.Name = product.Name
			prd.Quantity = product.Quantity
			prd.IsPublished = product.IsPublished
			prd.Expiration = product.Expiration
			prod = *prd
			return
		}
	}
	err = errors.New("no existe producto para actualizar")
	return
}

func (c *RepositoryController) Delete(id int) (product domain.Product, err error) {
	if id == len(c.Db) {
		for _, prd := range c.Db {
			if prd.Id == id {
				c.Db = c.Db[0 : id+1]
				product = *prd
				return
			}
		}
	}
	for _, prd := range c.Db {
		if prd.Id == id {
			c.Db = append(c.Db[:id+1], c.Db[id+1+1:]...) //TODO si es el ultimo seguro tenga error out of memory. VERIFICAR
			product = *prd
			return
		}
	}
	err = errors.New("no existe articulo para borrar")
	return
}

func FillProducts(path string, lastId int) *RepositoryController {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	var products []*domain.Product
	err = json.Unmarshal(file, &products)
	c := RepositoryController{
		Db:     products,
		LastId: lastId,
	}
	if err != nil {
		return nil
	}
	return &c
}

func (c *RepositoryController) validateUniqueCode(p *domain.Product) (err error) {
	for _, product := range c.Db {
		if product.CodeValue == p.CodeValue {
			err = errors.New("el código del producto debe ser unico")
			return
		}
	}
	return nil
}
