package home

import (
	"blogpost/internal/models"
	"strconv"
	"strings"

	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/wawandco/fako"
)

var _ models.UserService = (*service)(nil)

type service struct {
}

func NewService() *service {
	return &service{}
}

var users []models.User

// All returns a list of users based on the term, perPage, page, orderBy and order.
func (s *service) All(term string, perPage, page int, orderBy, order, status string) models.List {

	// Filter users by term
	var items models.Users
	for _, user := range users {
		term = strings.ToLower(term)
		name := strings.ToLower(user.Name)
		email := strings.ToLower(user.Email)
		if strings.Contains(name, term) || strings.Contains(email, term) {
			items = append(items, user)
		}
	}

	// Filter users by status
	if status != "all" {
		boolValue, _ := strconv.ParseBool(status)

		filtered := make([]models.User, 0, len(items))
		for _, user := range items {
			if user.Active == boolValue {
				filtered = append(filtered, user)
			}
		}

		items = filtered
	}

	items.Sort(orderBy, order)

	// Paginate users
	total := len(users)
	if term != "" || status != "all" {
		total = len(items)
	}

	// Calculate the start and end indexes for the items.
	start := (page - 1) * perPage
	end := start + perPage
	if end > len(items) {
		end = len(items)
	}

	// Get the items for the current page.
	items = items[start:end]

	return models.List{
		PerPage: perPage,
		Page:    page,
		Total:   total,
		Term:    term,
		Items:   items,
	}
}

// Validate validates the user.
func (s *service) Validate(user models.User) *validate.Errors {
	errs := validate.Validate(
		&validators.EmailIsPresent{Field: user.Email, Name: "Email"},
		&validators.StringIsPresent{Field: user.Name, Name: "Name"},
		&validators.StringIsPresent{Field: user.Phone, Name: "Phone"},
	)

	return errs
}

func init() {
	// Generate fake data
	users = make([]models.User, 10_000)
	for i := range users {
		fako.Fill(&users[i])

		if i%2 == 0 {
			users[i].Active = true
		}

	}
}
