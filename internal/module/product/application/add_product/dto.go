package addproduct

type AddProductInput struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Status      bool    `json:"status"`
}

type ProductOutput struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Status      bool    `json:"status"`
}

type AddProductOutput struct {
	Product ProductOutput `json:"product"`
}
