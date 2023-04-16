package addproduct

type AddProductInput struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Status      bool    `json:"status"`
}

type ProductOutput struct {
	ID          string
	Name        string
	Description string
	Price       float64
	Quantity    int
	Status      bool
}

type AddProductOutput struct {
	Product ProductOutput
}
