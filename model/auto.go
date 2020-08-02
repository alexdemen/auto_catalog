package model

const (
	TransitStatus    = "transit"
	InStockStatus    = "in stock"
	SoldStatus       = "sold"
	OutOfSalesStatus = "out of sales"
)

type Car struct {
	ID        int    `jsonapi:"primary,cars"`
	BrandName string `jsonapi:"attr,brand"`
	Model     string `jsonapi:"attr,model"`
	Price     int    `jsonapi:"attr,price"`
	Status    string `jsonapi:"attr,status"`
	Mileage   int    `jsonapi:"attr,mileage"`
}

func (c Car) Validate() (result string) {
	if c.BrandName == "" {
		result += "Brand name is empty. "
	}
	if c.Model == "" {
		result += "Model is empty. "
	}

	if c.Price < 0 {
		result += "Value of Price less 0. "
	}
	if c.Mileage < 0 {
		result += "Value of Mileage less 0. "
	}
	if c.Status != TransitStatus && c.Status != SoldStatus && c.Status != InStockStatus && c.Status != OutOfSalesStatus {
		result += "Value of Status is not valid."
	}

	return
}
