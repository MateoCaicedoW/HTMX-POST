package models

import (
	"sort"

	"github.com/gobuffalo/validate"
)

type User struct {
	Name   string `fako:"full_name"`
	Email  string `fako:"email_address"`
	Phone  string `fako:"phone"`
	Active bool
}

type Users []User

func (u Users) Sort(orderBy, order string) {
	switch orderBy {
	case "name":
		if order == "desc" {
			orderByKey(u, sortByName, false)
			return
		}
		orderByKey(u, sortByName, true)
	case "status":
		if order == "desc" {
			orderByKey(u, sortByStatus, true)
			return
		}
		orderByKey(u, sortByStatus, false)
	case "email":
		if order == "desc" {
			orderByKey(u, sortByEmail, true)
			return
		}
		orderByKey(u, sortByEmail, false)
	case "phone":
		if order == "desc" {
			orderByKey(u, sortByPhone, true)
			return
		}
		orderByKey(u, sortByPhone, false)
	}

}

type userSortFunc func(u1, u2 *User) bool

func orderByKey(users Users, key userSortFunc, reverse bool) {
	sort.Slice(users, func(i, j int) bool {
		if reverse {
			return key(&users[j], &users[i])
		}
		return key(&users[i], &users[j])
	})
}

func sortByStatus(u1, u2 *User) bool {
	return u1.Active && !u2.Active
}

func sortByName(u1, u2 *User) bool {
	return u1.Name < u2.Name
}

func sortByEmail(u1, u2 *User) bool {
	return u1.Email < u2.Email
}

func sortByPhone(u1, u2 *User) bool {
	return u1.Phone < u2.Phone
}

// UserService is the service that allows to interact with the users.
type UserService interface {
	Validate(user User) *validate.Errors
	All(term string, perpage, page int, OrderBy, Order, status string) List
}
