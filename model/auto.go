package model

const (
	TRANSIT_STATUS      = "transit"
	IN_STOCK_STATUS     = "in stock"
	SOLD_STATUS         = "sold"
	OUT_OF_SALES_STATUS = "out of sales"
)

type Car struct {
	ID        int64  `jsonapi:"primary,cars"`
	BrandName string `jsonapi:"attr,brand"`
	Model     string `jsonapi:"attr,model"`
	Price     int    `jsonapi:"attr,price"`
	Status    string `jsonapi:"attr,status"`
	Mileage   int    `jsonapi:"attr,mileage"`
}
