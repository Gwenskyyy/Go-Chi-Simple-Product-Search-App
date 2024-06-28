package controllers

import (
	"ProductSearchAppSql/models"
	"ProductSearchAppSql/config"
	"log"
)

func LoadProducts() ([]models.Product, error){
	var products []models.Product

	rows, err := config.DB.Query("SELECT barcode, name, price, active FROM products")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

	for rows.Next() {
        var p models.Product
        if err := rows.Scan(&p.Barcode, &p.Name, &p.Price, &p.Active); err != nil {
            return nil, err
        }
        products = append(products, p)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }

    log.Printf("Loaded %d products from database", len(products))
    return products, nil
} 