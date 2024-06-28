package routes

import (
	"ProductSearchAppSql/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/upload", controllers.UploadProducts)
	r.Get("/search/{barcode}", controllers.SearchProduct)
	r.Handle("/", http.FileServer(http.Dir("./pages")))

	return r
}