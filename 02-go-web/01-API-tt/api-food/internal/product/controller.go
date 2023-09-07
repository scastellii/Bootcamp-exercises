package product

func NewControllerProducts(db []*Product, lastId int) *ControllerProducts {
	return &ControllerProducts{
		db:     db,
		lastId: lastId,
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
	db     []*Product
	lastId int
}
