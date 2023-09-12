package handler

import (
	"bytes"
	"encoding/json"
	"github.com/bootcamp-go/Consignas-Go-Web.git/internal/domain"
	"github.com/bootcamp-go/Consignas-Go-Web.git/internal/product"
	"github.com/bootcamp-go/Consignas-Go-Web.git/pkg/store"
	"github.com/bootcamp-go/Consignas-Go-Web.git/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestProductHandler_GetAll(t *testing.T) {
	t.Run("Get all products correctly", func(t *testing.T) {
		r := createServer()
		httpReq, recorder := createRequestTest(http.MethodGet, "/products", "", "my-secret-token")

		//Act
		r.ServeHTTP(recorder, httpReq)
		//Assert)
		assert.Equal(t, http.StatusOK, recorder.Code)
		var objRes web.Response
		err := json.Unmarshal(recorder.Body.Bytes(), &objRes)
		assert.Nil(t, err)
	})

	t.Run("Get product 1", func(t *testing.T) {
		r := createServer()
		httpReq, recorder := createRequestTest(http.MethodGet, "/products/1", "", "my-secret-token")

		//Act
		r.ServeHTTP(recorder, httpReq)
		//Assert)
		assert.Equal(t, http.StatusOK, recorder.Code)
		var objRes web.Response
		err := json.Unmarshal(recorder.Body.Bytes(), &objRes)
		var prod domain.Product
		final, _ := json.Marshal(objRes.Data)
		err = json.Unmarshal(final, &prod)
		assert.Equal(t, prod, domain.Product{
			Id:          1,
			Name:        "Oil - Margarine",
			Quantity:    439,
			CodeValue:   "S82254D",
			IsPublished: true,
			Expiration:  "15/12/2021",
			Price:       71.42,
		})
		assert.Nil(t, err)
	})

}

func createServer() *gin.Engine {
	if err := godotenv.Load("/Users/scastelli/GolandProjects/bootcamp/02-go-web/05-API/cmd/server/.env"); err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	storage := store.NewStore("/Users/scastelli/GolandProjects/bootcamp/02-go-web/05-API/products.json")

	repo := product.NewRepository(storage)
	service := product.NewService(repo)
	productHandler := NewProductHandler(service)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	products := r.Group("/products")
	{
		products.GET("", productHandler.GetAll())
		products.GET(":id", productHandler.GetByID())
		products.GET("/search", productHandler.Search())
		products.POST("", productHandler.Post())
		products.DELETE(":id", productHandler.Delete())
		products.PATCH(":id", productHandler.Patch())
		products.PUT(":id", productHandler.Put())
	}
	return r
}

func createRequestTest(method string, url string, body string, token string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	if token != "" {
		req.Header.Add("TOKEN", token)
	}
	return req, httptest.NewRecorder()
}
func loadProducts(path string) ([]domain.Product, error) {
	var products []domain.Product
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(file), &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func writeProducts(path string, list []domain.Product) error {
	bytes, err := json.Marshal(list)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, bytes, 0644)
	if err != nil {
		return err
	}
	return err
}

func Test_Post_OK(t *testing.T) {
	var expectd = web.Response{Data: domain.Product{
		Id:          501,
		Name:        "Oil - Margarine",
		Quantity:    439,
		CodeValue:   "TEST4505010",
		IsPublished: true,
		Expiration:  "15/12/2021",
		Price:       50.50,
	}}

	product, _ := json.Marshal(expectd.Data)

	r := createServer()
	req, rr := createRequestTest(http.MethodPost, "/products", string(product), "my-super-secret-token")

	p, _ := loadProducts("/Users/scastelli/GolandProjects/bootcamp/02-go-web/05-API/products.json")

	r.ServeHTTP(rr, req)
	actual := map[string]domain.Product{}
	_ = json.Unmarshal(rr.Body.Bytes(), &actual)
	_ = writeProducts("/Users/scastelli/GolandProjects/bootcamp/02-go-web/05-API/products.json", p)

	assert.Equal(t, 201, rr.Code)
	assert.Equal(t, expectd.Data, actual["data"])

}

func Test_Delete_OK(t *testing.T) {

	r := createServer()
	req, rr := createRequestTest(http.MethodDelete, "/products/1", "", "my-super-secret-token")

	p, err := loadProducts("/Users/scastelli/GolandProjects/bootcamp/02-go-web/05-API/products.json")
	if err != nil {
		panic(err)
	}

	r.ServeHTTP(rr, req)

	err = writeProducts("/Users/scastelli/GolandProjects/bootcamp/02-go-web/05-API/products.json", p)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 204, rr.Code)
	assert.Nil(t, rr.Body.Bytes())
}

func Test_BadRequest(t *testing.T) {

	test := []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodDelete}

	r := createServer()
	for _, method := range test {
		req, rr := createRequestTest(method, "/products/not_number", "", "my-super-secret-token")
		r.ServeHTTP(rr, req)
		assert.Equal(t, 400, rr.Code)
	}

}

func Test_Unauthorized(t *testing.T) {

	test := []string{http.MethodPut, http.MethodPatch, http.MethodDelete}

	r := createServer()
	for _, method := range test {
		req, rr := createRequestTest(method, "/products/10", "{}", "not-my-token")
		r.ServeHTTP(rr, req)
		assert.Equal(t, 401, rr.Code)
	}
}

func Test_Post_Unauthorized(t *testing.T) {
	r := createServer()
	req, rr := createRequestTest(http.MethodPost, "/products", "{}", "not-my-token")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 401, rr.Code)
}
