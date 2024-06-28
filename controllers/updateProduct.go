package controllers

import (
	"ProductSearchAppSql/models"
	"database/sql"
)

func UpdateProducts(tx *sql.Tx, p models.Product) error {
	_, err := tx.Exec("UPDATE products SET name = $1, price = $2, active = $3 WHERE barcode = $4", p.Name, p.Price, p.Active, p.Barcode)
	return err
}