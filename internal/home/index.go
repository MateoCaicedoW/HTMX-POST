package home

import (
	"blogpost/internal/models"
	"net/http"
	"strconv"

	"github.com/leapkit/core/render"
)

func Index(w http.ResponseWriter, r *http.Request) {
	rw := render.FromCtx(r.Context())
	userService := r.Context().Value("userService").(models.UserService)
	term := r.URL.Query().Get("term")

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}

	orderBy := r.URL.Query().Get("orderBy")
	if orderBy == "" {
		orderBy = "name"
	}

	order := r.URL.Query().Get("order")
	if order == "" {
		order = "asc"
	}

	status := r.URL.Query().Get("status")
	if status == "" {
		status = "all"
	}

	users := userService.All(term, 5, page, orderBy, order, status)

	rw.Set("list", users)
	rw.Set("term", term)
	rw.Set("page", page)
	rw.Set("orderBy", orderBy)
	rw.Set("order", order)
	rw.Set("status", status)

	if r.Header.Get("HX-Request") == "true" {
		w.Header().Set("HX-Push-Url", r.URL.String())
		err := rw.RenderClean("home/table.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = rw.Render("home/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
