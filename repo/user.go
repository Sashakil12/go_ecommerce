package repo

type User struct {
	Id             int    `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	IsShopperOwner bool   `json:"is_shop_owner"`
}

type userRepo struct {
	userList []*User
}

type UserRepo interface {
	Create(u User) *User
	Get(email string, password string) *User
	Update(id int, u User) *User
	Delete(id int) bool
	List() []*User
}

func (r *userRepo) Create(u User) *User {
	r.userList = append(r.userList, &u)
	return &u
}

func (r *userRepo) Get(email string, password string) *User {
	for _, user := range r.userList {
		if user.Email == email && user.Password == password {
			return user
		}
	}
	return nil
}

func (r *userRepo) Update(id int, u User) *User {
	for _, user := range r.userList {
		if user.Id == id {
			user.FirstName = u.FirstName
			user.LastName = u.LastName
			user.Email = u.Email
			user.Password = u.Password
			user.IsShopperOwner = u.IsShopperOwner
			return user
		}
	}
	return nil
}

func (r *userRepo) Delete(id int) bool {
	for i, user := range r.userList {
		if user.Id == id {
			r.userList = append(r.userList[:i], r.userList[i+1:]...)
			return true
		}
	}
	return false
}

func (r *userRepo) List() []*User {
	return r.userList
}

func NewUserRepo() UserRepo {
	return &userRepo{
		userList: []*User{},
	}
}

func generateInitialUsers(r *userRepo) {
	users := []*User{
		{
			Id:             1,
			FirstName:      "John",
			LastName:       "Doe",
			Email:          "user1@yopmail.com",
			Password:       "password123",
			IsShopperOwner: false,
		},
		{
			Id:             2,
			FirstName:      "Jane",
			LastName:       "Smith",
			Email:          "user2@yopmail.com",
			Password:       "password123",
			IsShopperOwner: true,
		},
		{
			Id:             3,
			FirstName:      "Alice",
			LastName:       "Johnson",
			Email:          "user3@yopmail.com",
			Password:       "password123",
			IsShopperOwner: false,
		},
		{
			Id:             4,
			FirstName:      "Bob",
			LastName:       "Brown",
			Email:          "user4@yopmail.com",
			Password:       "password123",
			IsShopperOwner: true,
		},
		{
			Id:             5,
			FirstName:      "Charlie",
			LastName:       "Davis",
			Email:          "user5@yopmail.com",
			Password:       "password123",
			IsShopperOwner: false,
		},
	}
	r.userList = append(r.userList, users...)
}
