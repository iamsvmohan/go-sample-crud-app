package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/zlp-ecommerce/customer-service/docs"
	"github.com/zlp-ecommerce/customer-service/driver"
	ph "github.com/zlp-ecommerce/customer-service/handler/http"
)

func main() {
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connection, err := driver.ConnectSQL(dbHost, dbPort, "root", dbPass, dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	Handlers := ph.NewPostHandler(connection)
	r.Route("/", func(rt chi.Router) {
		rt.Mount("/customers", resgisterRoutes(Handlers))
	})
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	fmt.Println("Server listen at localhost:5601")
	fmt.Println("Server listen at: http://localhost:5601")
	fmt.Println("Swagger listen at: http://localhost:5601/swagger/index.html")
	http.ListenAndServe(":5601", r)
}

// A completely separate router for customer routes
func resgisterRoutes(Handlers *ph.Customer) http.Handler {
	r := chi.NewRouter()
	r.Get("/", Handlers.Fetch)
	r.Get("/{id:[0-9]+}", Handlers.GetByID)
	r.Post("/", Handlers.Create)
	r.Put("/{id:[0-9]+}", Handlers.Update)
	r.Delete("/{id:[0-9]+}", Handlers.Delete)

	return r
}
