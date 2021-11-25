package gophermart

import (
	user2 "gophermart/internal/app/handler/user" //todo change to 'user'
	"gophermart/internal/app/service/balance"
	"gophermart/internal/app/service/orders"
	"gophermart/internal/app/store"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
	store  store.Store
}

func newServer(store store.Store, accuralSystemAddress string) *server {
	s := &server{
		router: mux.NewRouter(),
		store:  store,
	}

	s.configureRouter(accuralSystemAddress)

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter(accuralSystemAddress string) {
	s.router.HandleFunc("/api/user/register", user2.RegisterHandler(s.store)).Methods(http.MethodPost)
	s.router.HandleFunc("/api/user/login", user2.LoginHandler(s.store)).Methods(http.MethodPost)

	private := s.router.NewRoute().Subrouter()
	private.Use(user2.AuthMiddleware)
	private.HandleFunc("/api/user/orders", orders.OrdersPostHandler(s.store, accuralSystemAddress)).Methods(http.MethodPost)
	private.HandleFunc("/api/user/orders", orders.OrdersGetHandler(s.store)).Methods(http.MethodGet)

	private.HandleFunc("/api/user/balance", balance.BalanceHandler(s.store)).Methods(http.MethodGet)
	private.HandleFunc("/api/user/balance/withdraw", balance.WithdrawPostHandler(s.store)).Methods(http.MethodPost)
	private.HandleFunc("/api/user/balance/withdrawals", balance.WithdrawGetHandler(s.store)).Methods(http.MethodGet)
}
