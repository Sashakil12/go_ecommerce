package repo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	Id          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImgUrl      string  `json:"image_url" db:"image_url"`
}
type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(id int) (*Product, error)
	List() ([]*Product, error)
	Update(id int, p Product) (*Product, error)
	Delete(id int) (bool, error)
}

type productRepo struct {
	db *sqlx.DB
}

func (r *productRepo) Create(p Product) (*Product, error) {
	query := `
		INSERT INTO products (title, description, price, image_url)
		VALUES (:title, :description, :price, :image_url)
		RETURNING id`

	// Prepare statement and execute
	rows, err := r.db.NamedQuery(query, p)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&p.Id)
	}
	return &p, nil
}

func (r *productRepo) Get(id int) (*Product, error) {
	var p Product
	query := `SELECT * FROM products WHERE id = $1`
	err := r.db.Get(&p, query, id)
	if err != nil {
		return nil, fmt.Errorf("product not found: %w", err)
	}
	return &p, nil
}

func (r *productRepo) List() ([]*Product, error) {
	var products []*Product
	query := `SELECT * FROM products ORDER BY id ASC`
	err := r.db.Select(&products, query)
	return products, err
}

func (r *productRepo) Update(id int, p Product) (*Product, error) {
	p.Id = id
	query := `
		UPDATE products 
		SET title = :title, description = :description, price = :price, image_url = :image_url
		WHERE id = :id`

	result, err := r.db.NamedExec(query, p)
	if err != nil {
		return nil, err
	}

	count, _ := result.RowsAffected()
	if count == 0 {
		return nil, fmt.Errorf("product not found")
	}
	return &p, nil
}

func (r *productRepo) Delete(id int) (bool, error) {
	query := `DELETE FROM products WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return false, err
	}

	count, _ := result.RowsAffected()
	return count > 0, nil
}
func NewProductRepo(db *sqlx.DB) ProductRepo {
	repo := &productRepo{
		db: db,
	}
	return repo
}
