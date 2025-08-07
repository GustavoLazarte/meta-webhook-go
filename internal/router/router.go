package router

import (
	"fmt"
	"net/http"
	"webhook/internal/controllers"

	"github.com/gorilla/mux"
)

type Router struct {
	Router *mux.Router
	http.Handler
	WebhookController controllers.Webhook
}

func NewRouter() *Router {
	return &Router{
		Router:            mux.NewRouter(),
		WebhookController: controllers.Webhook{},
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.Router.ServeHTTP(w, req) // Redirige las solicitudes a mux.Router
}

func (r *Router) SetupRouter() {
	r.Router.HandleFunc("/webhook", r.WebhookController.AllowAccess).Methods("GET")
	r.Router.HandleFunc("/webhook", r.WebhookController.HandleWebhookEvent).Methods("POST")
	r.Router.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Hola")
	}).Methods("GET")
}
