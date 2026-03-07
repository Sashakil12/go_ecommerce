package database

type Product struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"image_url"`
}

var ProductList []Product

func init() {
	// create 5 random product with valid image url free image is ok

	prod1 := Product{
		Id:          1,
		Title:       "Orange",
		Description: "demo orange",
		Price:       24.56,
		ImgUrl:      "https://images.unsplash.com/photo-1502741338009-cac277ee9bca?auto=format&fit=crop&w=400&q=80",
	}
	prod2 := Product{
		Id:          2,
		Title:       "Banana",
		Description: "demo banana",
		Price:       18.99,
		ImgUrl:      "https://images.unsplash.com/photo-1574226516831-e1dff420e8f8?auto=format&fit=crop&w=400&q=80",
	}
	prod3 := Product{
		Id:          3,
		Title:       "Mango",
		Description: "demo mango",
		Price:       30.00,
		ImgUrl:      "https://images.unsplash.com/photo-1519125323398-675f0ddb6308?auto=format&fit=crop&w=400&q=80",
	}
	prod4 := Product{
		Id:          4,
		Title:       "Pineapple",
		Description: "demo pineapple",
		Price:       22.50,
		ImgUrl:      "https://images.unsplash.com/photo-1465101046530-73398c7fda65?auto=format&fit=crop&w=400&q=80",
	}
	prod5 := Product{
		Id:          5,
		Title:       "Apple",
		Description: "demo apple",
		Price:       15.75,
		ImgUrl:      "https://images.unsplash.com/photo-1444065381814-865dc9f9e736?auto=format&fit=crop&w=400&q=80",
	}
	ProductList = append(ProductList, prod1)
	ProductList = append(ProductList, prod2)
	ProductList = append(ProductList, prod3)
	ProductList = append(ProductList, prod4)
	ProductList = append(ProductList, prod5)
}
