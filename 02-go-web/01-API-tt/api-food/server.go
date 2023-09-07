package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

//TODO crear estructura de respuesta
//TODO crear constructor
//TODO crear struct con config
//TODO Crear estructura controller repository y service y refactorizar el codigo

func Init() {
	router := gin.Default()
	router.GET("/ping", ping)
	pr := router.Group("/products")
	pr.GET("", getProducts)
	pr.GET("/:id", getProduct)
	pr.POST("", SaveProduct())
	router.Run()

}

func ping(c *gin.Context) {
	c.String(200, "pong")

}

func getProducts(c *gin.Context) {
	if qParam, ok := c.GetQuery("priceGt"); ok {
		priceToSearch, _ := strconv.Atoi(qParam)
		var filteredProducts []Product
		for _, product := range products {
			if product.Price > float64(priceToSearch) {
				filteredProducts = append(filteredProducts, product)
			}
		}
		c.JSON(http.StatusOK, &filteredProducts)
		return
	}
	c.JSON(http.StatusOK, products)
}

func getProduct(c *gin.Context) {
	idToSearch := c.Param("id")
	idParse, _ := strconv.Atoi(idToSearch)
	for _, product := range products {
		if product.Id == idParse {
			c.JSON(http.StatusOK, product)
			return
		}
	}
	c.JSON(500, "No se encontró el producto")
	return
}

func SaveProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var product Product
		err := c.ShouldBindJSON(&product)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Se ingresaron campos de forma incorrecta")
			return
		}
		// si esta ok
		if err := validateProduct(product); err != nil {
			c.JSON(http.StatusBadRequest, "Se ingresaron campos de forma incorrecta. Intente nuevamente")
			return
		}
		product.Id = LastId + 1
		products = append(products, product)
		c.JSON(http.StatusCreated, "Se creo el producto correctamente")
		LastId++
	}
}

type EmptyError struct {
	Field string
}

func (e *EmptyError) Error() string {
	return fmt.Sprintf("El %s del producto no debe estar vacio", e.Field)
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
	for _, product := range products {
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
	fmt.Println(err)
	if err != nil {
		return
	}
	return nil
}

/*
Ejercicio 1: Añadir un producto
En esta ocasión vamos a añadir un producto al slice cargado en memoria. Dentro de la ruta /products añadimos el método POST,
al cual vamos a enviar en el cuerpo de la request el nuevo producto. El mismo tiene ciertas restricciones, conozcámoslas:
No es necesario pasar el Id, al momento de añadirlo se debe inferir del estado de la lista de productos,
verificando que no se repitan ya que debe ser un campo único.
Ningún dato puede estar vacío, exceptuando is_published (vacío indica un valor false).
El campo code_value debe ser único para cada producto.
Los tipos de datos deben coincidir con los definidos en el planteo del problema.
La fecha de vencimiento debe tener el formato: XX/XX/XXXX, además debemos verificar que día, mes y año sean valores válidos.
Recordá: si una consulta está mal formulada por parte del cliente, el status code cae en los 4XX.
Ejercicio 2: Traer el producto
Realiza una consulta a un método GET con el id del producto recién añadido, tené en cuenta que la lista de producto se encuentra cargada en la memoria,
si terminás la ejecución del programa este producto no estará en la próxima ejecución.
*/
