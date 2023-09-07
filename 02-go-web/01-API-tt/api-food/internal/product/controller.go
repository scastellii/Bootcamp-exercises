package product

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func NewControllerProducts(db []*Product, lastId int) *ControllerProducts {
	return &ControllerProducts{
		Db:     db,
		LastId: lastId,
	}
}

type RequestProduct struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_Value"`
	IsPublished bool    `json:"is_Published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type ResponseProduct struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_Value"`
	IsPublished bool    `json:"is_Published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

// ControllerProducts is an struct that represents a controller for products
// exposing methods to handle products
type ControllerProducts struct {
	Db     []*Product
	LastId int
}

func productToResponseProduct(product Product) ResponseProduct {
	return ResponseProduct{
		Id:          product.Id,
		Name:        product.Name,
		Quantity:    product.Quantity,
		CodeValue:   product.CodeValue,
		IsPublished: product.IsPublished,
		Expiration:  product.Expiration,
		Price:       product.Price,
	}
}

func (contr *ControllerProducts) GetProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		var resProducts []ResponseProduct
		if qParam, ok := c.GetQuery("priceGt"); ok {
			priceToSearch, _ := strconv.Atoi(qParam)
			for _, product := range contr.Db {
				if product.Price > float64(priceToSearch) {
					resPrd := productToResponseProduct(*product)
					resProducts = append(resProducts, resPrd)
				}
			}
			c.JSON(http.StatusOK, &resProducts)
			return
		}
		for _, product := range contr.Db {
			resPrd := productToResponseProduct(*product)
			resProducts = append(resProducts, resPrd)
		}
		c.JSON(http.StatusOK, &resProducts)
	}
}

func (contr *ControllerProducts) GetProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		idToSearch := c.Param("id")
		idParse, _ := strconv.Atoi(idToSearch)
		for _, product := range contr.Db {
			if product.Id == idParse {
				resProduct := productToResponseProduct(*product)
				c.JSON(http.StatusOK, &resProduct)
				return
			}
		}
		c.JSON(500, "No se encontró el producto")
		return
	}
}

func (contr *ControllerProducts) SaveProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqProduct RequestProduct
		err := c.ShouldBindJSON(&reqProduct)
		if err != nil {
			c.JSON(http.StatusBadRequest, "El JSON brindado no es válido")
			return
		}
		// si esta ok
		product := Product{
			Name:        reqProduct.Name,
			Quantity:    reqProduct.Quantity,
			CodeValue:   reqProduct.CodeValue,
			IsPublished: reqProduct.IsPublished,
			Expiration:  reqProduct.Expiration,
			Price:       reqProduct.Price,
		}

		if err := validateProduct(product); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, "Alguno de los campos ingresados no es válido. Intente nuevamente")
			return
		}
		product.Id = contr.LastId + 1
		Save(product, contr)
		res := productToResponseProduct(product)
		c.JSON(http.StatusCreated, res)
		contr.LastId++
	}
}
