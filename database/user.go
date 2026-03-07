package database

type User struct {
	Id             int    `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	IsShopperOwner bool   `json:"is_shop_owner"`
}

var users []User

func (u User) Store() *User {
	users = append(users, u)
	return &users[len(users)-1]
}

func (u User) Find(email string, password string) *User {
	for idx, user := range users {
		if user.Email == email && user.Password == password {
			return &users[idx]
		}
	}
	return nil
}
