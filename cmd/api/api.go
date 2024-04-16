package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/shammianand/ecom/service/cart"
	"github.com/shammianand/ecom/service/httphelp"
	"github.com/shammianand/ecom/service/order"
	"github.com/shammianand/ecom/service/product"
	"github.com/shammianand/ecom/service/user"
	"github.com/shammianand/ecom/utils"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {

	router := http.NewServeMux()
	subrouter := httphelp.Subrouter(router, "/api/v1")

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)

	orderStore := order.NewStore(s.db)

	cartHandler := cart.NewHandler(orderStore, productStore, userStore)
	cartHandler.RegisterRoutes(subrouter)

	log.Println("Listening on: ", s.addr)
	return http.ListenAndServe(s.addr, utils.Logging(router))
}
