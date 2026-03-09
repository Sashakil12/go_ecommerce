package repo

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type User struct {
	Id             int    `json:"id" db:"id"`
	FirstName      string `json:"first_name" db:"first_name"`
	LastName       string `json:"last_name" db:"last_name"`
	Email          string `json:"email" db:"email"`
	Password       string `json:"password" db:"password"`
	IsShopperOwner bool   `json:"is_shop_owner" db:"is_shop_owner"`
}

type userRepo struct {
	db *sqlx.DB
}

type UserRepo interface {
	Create(u User) (*User, error)
	Get(email string, password string) (*User, error)
	Update(id int, u User) (*User, error)
	Delete(id int) (bool, error)
	List() ([]*User, error)
}

func (r *userRepo) Create(u User) (*User, error) {
	query := `
        INSERT INTO users 
		(first_name, last_name, email, password, is_shop_owner)
        VALUES 
		(:first_name, :last_name, :email, :password, :is_shop_owner)
        RETURNING id`

	// NamedQuery handles the mapping; we use .Next() to get the returned ID
	rows, err := r.db.NamedQuery(query, u)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // ALWAYS close your rows!

	var userId int
	if rows.Next() {
		if err := rows.Scan(&userId); err != nil {
			return nil, err
		}
	}
	u.Id = userId
	return &u, nil
}

func (r *userRepo) Get(email string, password string) (*User, error) {
	var user User

	// sqlx.Get handles the mapping of columns to struct fields automatically
	query := `SELECT id, first_name, last_name, email, password, is_shop_owner
	 FROM users 
	 WHERE email = $1 AND password = $2 LIMIT 1`

	err := r.db.Get(&user, query, email, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) Update(id int, u User) (*User, error) {
	// Ensure the ID from the URL/Param matches the struct ID for the query
	u.Id = id

	query := `
        UPDATE users 
        SET first_name = :first_name, 
            last_name = :last_name, 
            email = :email, 
            password = :password, 
            is_shop_owner = :is_shop_owner
        WHERE id = :id`

	result, err := r.db.NamedExec(query, u)
	if err != nil {
		return nil, err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, nil // Or return a custom "Not Found" error
	}

	return &u, nil
}
func (r *userRepo) Delete(id int) (bool, error) {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	// Returns true if a row was actually deleted, false if the ID didn't exist
	return rowsAffected > 0, nil
}

func (r *userRepo) List() ([]*User, error) {
	// Create an empty slice to hold the results
	var users []*User

	// Select maps all rows returned by the query into the users slice
	query := `SELECT id, first_name, last_name, email, password, is_shop_owner FROM users`

	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}
