package product

type product struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// FIX: Add the database logic

var productList = []product{
	product{ID: 1, Title: "Apples", Content: "Apples: Content"},
	product{ID: 2, Title: "Bananas", Content: "Bananas: Content"},
	product{ID: 3, Title: "Pears", Content: "Pears: Pears"},
	product{ID: 4, Title: "Oranges", Content: "Oranges: Oranges"},
}

// Return a list of all the products
func getAllProducts() []product {
	return productList
}
