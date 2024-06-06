package product

type Products struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

type CreateProductsInput struct {
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

type UpdateProductsInput struct {
	Name  string `json:"name"`
	Price uint   `json:"price"`
}
