package controllers

import (
	"ProductSearchAppSql/config"
	"ProductSearchAppSql/models"
	"bufio"
	"net/http"
	"strconv"
	"strings"
	"database/sql"
)

func UploadProducts(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		header := scanner.Text()
		if header == "" {
			http.Error(w, "Missing header line", http.StatusBadRequest)
			return
		}
	}

	var newProducts []models.Product
	var updatedProducts []models.Product

	for scanner.Scan() {
		line := scanner.Text()
		record := strings.Split(line, "|")

		if len(record) < 4 {
			http.Error(w, "Invalid record format", http.StatusBadRequest)
			return
		}

		barcode, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		price, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		active, err := strconv.ParseBool(record[3])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		newProduct := models.Product{
			Barcode: barcode,
			Name:    record[1],
			Price:   price,
			Active:  active,
		}

		existingProduct := models.Product{}
		err = config.DB.QueryRow("SELECT barcode, name, price, active FROM products WHERE barcode = $1", barcode).Scan(&existingProduct.Barcode, &existingProduct.Name, &existingProduct.Price, &existingProduct.Active)

		if err != nil && err != sql.ErrNoRows {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err == sql.ErrNoRows {
			newProducts = append(newProducts, newProduct)
		} else {
			existingProduct.Name = newProduct.Name
			existingProduct.Price = newProduct.Price
			existingProduct.Active = newProduct.Active
			updatedProducts = append(updatedProducts, existingProduct)
		}
	}

	tx, err := config.DB.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}()

	for _, p := range newProducts {
		if err := SaveProducts(tx, p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	for _, p := range updatedProducts {
		if err := UpdateProducts(tx, p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
