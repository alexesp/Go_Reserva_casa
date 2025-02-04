package handlers

import (
	"fmt"
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

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Casas(w http.ResponseWriter, r *http.Request) {
//fmt.Fprintf(w, "Es la pagina de inicio")

	// remoteIP := r.RemoteAddr
	// m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r,"casas.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Habitaciones(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "habitaciones.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Precios(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r,"precios.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Contacto(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r,"contacto.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r,"make-reservation.page.tmpl", &models.TemplateData{})
}
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("start data is %s and end date is %s", start, end)))
	render.RenderTemplate(w, r,"make-reservation.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Search(w http.ResponseWriter, r *http.Request) {	
	render.RenderTemplate(w, r,"search-abailability.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {	
	render.RenderTemplate(w, r,"generals.page.tmpl", &models.TemplateData{})
}
func (m *Repository) PostGenerals(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("start data is %s and end date is %s", start, end)))
	render.RenderTemplate(w, r,"generals.page.tmpl", &models.TemplateData{})
}