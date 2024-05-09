package models

// List is a generic list model that can be used to paginate
// a list of items.
type List struct {
	PerPage int
	Page    int
	Total   int
	Pages   []paginationValue

	Term  string
	Items interface{}
}

type paginationValue struct {
	Value  interface{}
	IsPage bool
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

func (l List) pagesRange() []int {
	if l.Pages != nil {
		pages := make([]int, 0)
		for _, page := range l.Pages {
			if page.IsPage {
				pages = append(pages, page.Value.(int))
			}
		}
		return pages
	}

	pages := make([]int, l.TotalPages())
	for i := range pages {
		pages[i] = i + 1
	}

	return pages
}

func (l List) NumericPagination() []paginationValue {
	pages := l.pagesRange()
	pagination := make([]paginationValue, 0)
	append := func(value interface{}, isPage bool) {
		pagination = append(pagination, paginationValue{
			Value:  value,
			IsPage: isPage,
		})
	}

	if len(pages) <= 5 {
		for _, page := range pages {
			append(page, true)
		}
		return pagination
	}

	if l.Page <= 3 {
		for i := 0; i < 5; i++ {
			append(pages[i], true)
		}
		append("...", false)
	} else if l.Page >= len(pages)-2 {
		append(pages[0], true)
		append("...", false)
		for i := len(pages) - 5; i < len(pages); i++ {
			append(pages[i], true)
		}
	} else {
		append(pages[0], true)
		append("...", false)
		for i := l.Page - 2; i <= l.Page+2; i++ {
			append(pages[i], true)
		}

		append("...", false)
		append(pages[len(pages)-1], true)
	}
	return pagination
}
