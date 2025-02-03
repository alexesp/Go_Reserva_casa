package handlers

import (
	"net/http"

	"github.com/alexesp/Go_Reserva_casa/pkg/config"
	"github.com/alexesp/Go_Reserva_casa/pkg/models"
	"github.com/alexesp/Go_Reserva_casa/pkg/render"
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

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Casas(w http.ResponseWriter, r *http.Request) {
//fmt.Fprintf(w, "Es la pagina de inicio")

	// remoteIP := r.RemoteAddr
	// m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "casas.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Habitaciones(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "habitaciones.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Precios(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "precios.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Contacto(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contacto.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Search(w http.ResponseWriter, r *http.Request) {	
	render.RenderTemplate(w, "search-abailability.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {	
	render.RenderTemplate(w, "generals.page.tmpl", &models.TemplateData{})
}