package home

import (
	"blogpost/internal/models"
	"net/http"

	"github.com/leapkit/core/form"
	"github.com/leapkit/core/render"
)

func New(w http.ResponseWriter, r *http.Request) {
	rx := render.FromCtx(r.Context())
	user := models.User{}

	rx.Set("user", user)
	if err := rx.RenderWithLayout("home/new.html", "modal.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	userService := r.Context().Value("userService").(models.UserService)
	if err := form.Decode(r, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	errors := userService.Validate(user)
	if errors.HasAny() {
		rx := render.FromCtx(r.Context())
		rx.Set("user", user)
		rx.Set("errors", errors.Errors)
		if err := rx.RenderWithLayout("home/new.html", "modal.html"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	users = append(users, user)

	w.Header().Set("HX-Redirect", "/")
}
