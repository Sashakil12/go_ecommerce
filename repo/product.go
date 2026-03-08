package repo

import "fmt"

type Product struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"image_url"`
}
type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(id int) (*Product, error)
	List() ([]*Product, error)
	Update(id int, p Product) (*Product, error)
	Delete(id int) (bool, error)
}

type productRepo struct {
	productList []*Product
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.Id = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}

func (r *productRepo) Get(id int) (*Product, error) {
	for i := range r.productList {
		if r.productList[i].Id == id {
			return r.productList[i], nil
		}
	}
	return nil, fmt.Errorf("product not found")
}
func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}

func (r *productRepo) Update(id int, p Product) (*Product, error) {
	for i := range r.productList {
		if r.productList[i].Id == id {
			// Update fields instead of replacing pointer
			r.productList[i].Title = p.Title
			r.productList[i].Description = p.Description
			r.productList[i].Price = p.Price
			r.productList[i].ImgUrl = p.ImgUrl
			return r.productList[i], nil
		}
	}
	return nil, fmt.Errorf("product not found")
}

func (r *productRepo) Delete(id int) (bool, error) {
	for i := range r.productList {
		if r.productList[i].Id == id {
			// Remove element efficiently
			copy(r.productList[i:], r.productList[i+1:])
			r.productList = r.productList[:len(r.productList)-1]
			// Optionally reassign IDs for remaining products
			for j := i; j < len(r.productList); j++ {
				r.productList[j].Id = j + 1
			}
			return true, nil
		}
	}
	return false, fmt.Errorf("product not found")
}

func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateInitialProducts(repo)
	return repo
}

func generateInitialProducts(r *productRepo) {

	prod1 := &Product{
		Id:          1,
		Title:       "Orange",
		Description: "demo orange",
		Price:       24.56,
		ImgUrl:      "https://images.unsplash.com/photo-1502741338009-cac277ee9bca?auto=format&fit=crop&w=400&q=80",
	}
	prod2 := &Product{
		Id:          2,
		Title:       "Banana",
		Description: "demo banana",
		Price:       18.99,
		ImgUrl:      "https://images.unsplash.com/photo-1574226516831-e1dff420e8f8?auto=format&fit=crop&w=400&q=80",
	}
	prod3 := &Product{
		Id:          3,
		Title:       "Mango",
		Description: "demo mango",
		Price:       30.00,
		ImgUrl:      "https://images.unsplash.com/photo-1519125323398-675f0ddb6308?auto=format&fit=crop&w=400&q=80",
	}
	prod4 := &Product{
		Id:          4,
		Title:       "Pineapple",
		Description: "demo pineapple",
		Price:       22.50,
		ImgUrl:      "https://images.unsplash.com/photo-1465101046530-73398c7fda65?auto=format&fit=crop&w=400&q=80",
	}
	prod5 := &Product{
		Id:          5,
		Title:       "Apple",
		Description: "demo apple",
		Price:       15.75,
		ImgUrl:      "https://images.unsplash.com/photo-1444065381814-865dc9f9e736?auto=format&fit=crop&w=400&q=80",
	}
	r.productList = append(r.productList, prod1)
	r.productList = append(r.productList, prod2)
	r.productList = append(r.productList, prod3)
	r.productList = append(r.productList, prod4)
	r.productList = append(r.productList, prod5)
}
