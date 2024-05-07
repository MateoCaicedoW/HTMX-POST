package models

type User struct {
	Name     string `fako:"full_name"`
	Username string `fako:"user_name"`
	Email    string `fako:"email_address"`
	Phone    string `fako:"phone"`
	Active   bool
}

// OrderUserByName is a type that allows to sort a list of users by name.
type OrderUserByName []User

func (u OrderUserByName) Len() int {
	return len(u)
}

func (u OrderUserByName) Less(i, j int) bool {
	return u[i].Name < u[j].Name
}

func (u OrderUserByName) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

// OrderUserByStatus is a type that allows to sort a list of users by status.
type OrderUserByStatus []User

func (u OrderUserByStatus) Len() int {
	return len(u)
}

func (u OrderUserByStatus) Less(i, j int) bool {
	return u[i].Active
}

func (u OrderUserByStatus) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

// UserService is the service that allows to interact with the users.
type UserService interface {
	All(term string, perpage, page int, orderBy, order, status string) List
}
