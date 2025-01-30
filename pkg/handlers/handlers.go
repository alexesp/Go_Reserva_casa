package handlers

import (
	"net/http"

	"github.com/alexesp/Go_Web.git/pkg/config"
	"github.com/alexesp/Go_Web.git/pkg/models"
	"github.com/alexesp/Go_Web.git/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Es la pagina de inicio")
<<<<<<< HEAD
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
=======
>>>>>>> 78dbfccc79d9e2e36fc15478bb540ffb722be8bc
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Es la pagina sobre!")
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
<<<<<<< HEAD

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP
=======
>>>>>>> 78dbfccc79d9e2e36fc15478bb540ffb722be8bc
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
