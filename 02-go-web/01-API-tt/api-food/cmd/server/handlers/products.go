package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-bases/02-go-web/01-API-tt/api-food/internal/domain"
	"go-bases/02-go-web/01-API-tt/api-food/internal/products"
	"net/http"
	"strconv"
	"time"
)

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
type Controller interface {
	GetAll() gin.HandlerFunc
	GetById() gin.HandlerFunc
	Store() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
}

func NewProductController(service *products.ServiceController) *ProductController {
	return &ProductController{
		Service: service,
	}
}

type ProductController struct {
	Service products.Service
}

func (p *ProductController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		var response []ResponseProduct
		if qParam, ok := c.GetQuery("priceGt"); ok {
			priceToSearch, _ := strconv.Atoi(qParam)
			productsAll, err := p.Service.GetByPrice(float64(priceToSearch))
			if err != nil {
				//TODO Manejar error y devolver codigo segun errror
				return
			}
			for _, product := range productsAll {
				response = append(response, productToResponseProduct(product))
			}
			c.JSON(http.StatusOK, &response)
			return
		}
		productsAll, err := p.Service.GetAll()
		if err != nil {
			//TODO Manejar error y devolver codigo segun errror
			return
		}
		for _, product := range productsAll {
			pr := productToResponseProduct(product)
			response = append(response, pr)
		}
		c.JSON(http.StatusOK, &response)
		return
	}
}

func (p *ProductController) GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		idToSearch := c.Param("id")
		idParse, _ := strconv.Atoi(idToSearch)
		pr, err := p.Service.GetById(idParse)
		if err != nil {
			//TODO manejar errores
			return
		}
		resProduct := productToResponseProduct(*pr)
		c.JSON(http.StatusOK, &resProduct)
		return
	}
}

func (p *ProductController) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqProduct RequestProduct
		err := c.ShouldBindJSON(&reqProduct)
		if err != nil {
			c.JSON(http.StatusBadRequest, "El JSON brindado no es válido")
			return
		}
		// si esta ok
		err = ValidateProduct(reqProduct)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, "Alguno de los campos ingresados no es válido. Intente nuevamente")
			return
		}
		product := requestProductToProduct(reqProduct)
		storedProduct, err := p.Service.Store(&product)
		if err != nil {
			return
		}
		res := productToResponseProduct(*storedProduct)
		c.JSON(http.StatusCreated, res)
		return
	}
}

func (p *ProductController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqProduct RequestProduct
		err := c.ShouldBindJSON(&reqProduct)
		if err != nil {
			c.JSON(http.StatusBadRequest, "El JSON brindado no es válido")
			return
		}
		// si esta ok
		err = ValidateProduct(reqProduct)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, "Alguno de los campos ingresados no es válido. Intente nuevamente")
			return
		}
		idToSearch := c.Param("id")
		idParse, _ := strconv.Atoi(idToSearch)
		product := requestProductToProduct(reqProduct)
		updatedProduct, err := p.Service.Update(idParse, &product)
		if err != nil {
			//TODO manejar error
			return
		}
		res := productToResponseProduct(*updatedProduct)
		c.JSON(http.StatusOK, res)
		return
	}
}

func (p *ProductController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idToSearch := c.Param("id")
		idParse, _ := strconv.Atoi(idToSearch)
		deletedProduct, err := p.Service.Delete(idParse)
		if err != nil {
			//TODO manejar error
			return
		}
		res := productToResponseProduct(*deletedProduct)
		c.JSON(http.StatusOK, res)
		return
	}
}

func productToResponseProduct(product domain.Product) ResponseProduct {
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

func requestProductToProduct(reqPoduct RequestProduct) domain.Product {
	return domain.Product{
		Name:        reqPoduct.Name,
		Quantity:    reqPoduct.Quantity,
		CodeValue:   reqPoduct.CodeValue,
		IsPublished: reqPoduct.IsPublished,
		Expiration:  reqPoduct.Expiration,
		Price:       reqPoduct.Price,
	}
}

func ValidateProduct(product RequestProduct) (err error) {
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
	if err = validateDate(product); err != nil {
		return
	}

	return nil
}

func validateDate(p RequestProduct) (err error) {
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
