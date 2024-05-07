package home

import (
	"blogpost/internal/models"
	"sort"
	"strconv"
	"strings"

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
	items := []models.User{}
	for _, user := range users {
		term = strings.ToLower(term)
		name := strings.ToLower(user.Name)
		if strings.Contains(name, term) {
			items = append(items, user)
		}
	}

	// Filter users by status
	if status != "all" {
		boolValue, _ := strconv.ParseBool(status)

		filtered := []models.User{}
		for _, user := range items {
			if user.Active == boolValue {
				filtered = append(filtered, user)
			}
		}

		items = filtered
	}

	// Sort users
	if orderBy == "name" {
		if order == "desc" {
			sort.Sort(sort.Reverse(models.OrderUserByName(items)))
		}

		if order == "asc" {
			sort.Sort(models.OrderUserByName(items))
		}
	}

	if orderBy == "status" {
		if order == "desc" {
			sort.Sort(sort.Reverse(models.OrderUserByStatus(items)))
		}

		if order == "asc" {
			sort.Sort(models.OrderUserByStatus(items))
		}
	}

	// Paginate users
	total := len(users)
	if term != "" || status != "all" {
		total = len(items)
	}

	// If the page is greater than the total number of pages
	// we need to reset the page to 1.
	if page > len(items)/perPage {
		page = 1
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

func init() {
	// Generate fake data
	users = make([]models.User, 20)
	for i := range users {
		fako.Fill(&users[i])

		if i%2 == 0 {
			users[i].Active = true
		}

	}
}
