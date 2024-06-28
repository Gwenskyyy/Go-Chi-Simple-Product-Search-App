package controllers

import (
	"ProductSearchAppSql/models"
	"database/sql"
)

func SaveProducts(tx *sql.Tx, p models.Product) error {
	_, err := tx.Exec("INSERT INTO products (barcode, name, price, active) VALUES ($1, $2, $3, $4)", p.Barcode, p.Name, p.Price, p.Active)
	return err
}