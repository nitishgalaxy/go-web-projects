package handlers

import (
	"net/http"

	"github.com/nitishgalaxy/go-course/pkg/config"
	"github.com/nitishgalaxy/go-course/pkg/models"
	"github.com/nitishgalaxy/go-course/pkg/render"
)

// Repo is the Repossitory used by the handlers
var Repo *Repository

// Repository is the Repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handler
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	// Store remote IP in the session.
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// Perform some Business Logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello"

	// m.App.Session is accessble here
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// Send data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
