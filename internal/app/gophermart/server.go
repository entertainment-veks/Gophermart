package gophermart

import (
	"gophermart/internal/app/service/orders"
	"gophermart/internal/app/service/user"
	"gophermart/internal/app/store"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/api/user/register", user.RegisterHandler(s.store)).Methods(http.MethodPost)
	s.router.HandleFunc("/api/user/login", user.LoginHandler(s.store)).Methods(http.MethodPost)

	private := s.router.NewRoute().Subrouter()
	private.Use(user.AuthMiddleware)
	private.HandleFunc("/api/user/orders", orders.OrdersPostHandler(s.store)).Methods(http.MethodPost)
	private.HandleFunc("/api/user/orders", orders.OrdersGetHandler(s.store)).Methods(http.MethodGet)

}
