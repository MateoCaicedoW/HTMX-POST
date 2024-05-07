package models

// List is a generic list model that can be used to paginate
// a list of items.
type List struct {
	PerPage int
	Page    int
	Total   int

	Term string

	Items interface{}
}

// TotalPages returns the total number of pages based on the total
// number of items and the number of items per page.
func (l List) TotalPages() (pages int) {
	// Calculate the total number of pages based on the total number of items
	// and the number of items per page.
	pages = l.Total / l.PerPage
	// If the total number of items is not divisible by the number of items per page
	// we need to add an extra page.
	if (l.Total % l.PerPage) > 0.0 {
		pages++
	}

	// If there are no items we still need to have at least one page.
	if pages == 0 {
		pages = 1
	}

	return pages
}

// HasNext returns true if there is a next page.
func (l List) HasNext() bool {
	return l.Page < l.TotalPages()
}

// HasPrev returns true if there is a previous page.
func (l List) HasPrev() bool {
	return l.Page > 1
}

// NextPage returns the next page number.
func (l List) NextPage() (next int) {
	next = l.Page + 1
	if l.Page == l.TotalPages() {
		next = l.Page
	}

	return next
}

// PrevPage returns the previous page number.
func (l List) PrevPage() (prev int) {
	prev = l.Page - 1
	if prev < 0 {
		prev = 0
	}

	return prev
}
