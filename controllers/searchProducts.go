package controllers

import (
	"ProductSearchAppSql/config"
	"ProductSearchAppSql/models"
	"encoding/json"
	"net/http"
	"strconv"

	"database/sql"
	"github.com/go-chi/chi/v5"
)

func SearchProduct(w http.ResponseWriter, r *http.Request) {
	barcode := chi.URLParam(r, "barcode")
	parsedBarcode, err := strconv.ParseInt(barcode, 10, 64)
	if err != nil {
		http.Error(w, "Invalid barcode", http.StatusBadRequest)
		return
	}
 
	var product models.Product
	row := config.DB.QueryRow("SELECT barcode, name, price, active FROM products WHERE barcode = $1 AND active = $2", parsedBarcode, true)
	switch err := row.Scan(&product.Barcode, &product.Name, &product.Price, &product.Active); err {
	case nil:
		json.NewEncoder(w).Encode(product)
	case sql.ErrNoRows:
		http.Error(w, "Product not found", http.StatusNotFound)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
