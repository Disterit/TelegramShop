package repository

import (
	Telegram_Market "Telegram-Market"
	"database/sql"
	"fmt"
	"strings"
)

type ProductTelegramSql struct {
	db *sql.DB
}

func NewProductTelegramSql(db *sql.DB) *ProductTelegramSql {
	return &ProductTelegramSql{db: db}
}

func (s *ProductTelegramSql) CreateProduct(product Telegram_Market.Products) (int64, error) {
	query := fmt.Sprintf(`INSERT INTO %s (description, photo_url, price, quantity, location_city, location_coordinates, paid_flag) 
									VALUES (?, ?, ?, ?, ?, ?, ?)`, TableProducts)

	res, err := s.db.Exec(query,
		product.Description,
		product.PhotoUrl,
		product.Price,
		product.Quantity,
		product.LocationCity,
		product.LocationCoordinates,
		product.PaidFlag,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to insert product: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}

	return id, nil
}

func (s *ProductTelegramSql) GetProductById(productId int64) (Telegram_Market.Products, error) {
	var product Telegram_Market.Products

	query := fmt.Sprintf(`SELECT 
        id, description, photo_url, price, quantity, location_city, location_coordinates, hidden_photo_url, hidden_description, paid_flag 
        FROM %s WHERE id = ?`, TableProducts)

	row := s.db.QueryRow(query, productId)

	// Считываем значения из строки в структуру `product`
	err := row.Scan(
		&product.Id,
		&product.Description,
		&product.PhotoUrl,
		&product.Price,
		&product.Quantity,
		&product.LocationCity,
		&product.LocationCoordinates,
		&product.HiddenPhotoUrl,
		&product.HiddenDescription,
		&product.PaidFlag,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return product, fmt.Errorf("product with ID %d not found", productId)
		}
		return product, fmt.Errorf("failed to scan product: %w", err)
	}

	return product, nil
}

func (s *ProductTelegramSql) GetAllProducts() ([]Telegram_Market.Products, error) {
	query := fmt.Sprintf("SELECT * FROM %s", TableProducts)
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get products from database: %w", err)
	}

	var products []Telegram_Market.Products
	for rows.Next() {
		var product Telegram_Market.Products
		err = rows.Scan(
			&product.Id,
			&product.Description,
			&product.PhotoUrl,
			&product.Price,
			&product.Quantity,
			&product.LocationCity,
			&product.LocationCoordinates,
			&product.HiddenPhotoUrl,
			&product.HiddenDescription,
			&product.PaidFlag)

		if err != nil {
			return nil, fmt.Errorf("failed to get products from database: %w", err)
		}
		products = append(products, product)
	}

	return products, nil
}

func (s *ProductTelegramSql) DeleteProduct(productId int64) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = ?`, TableProducts)
	_, err := s.db.Exec(query, productId)
	if err != nil {
		return fmt.Errorf("failed to delete product from database: %w", err)
	}

	return nil
}

func (s *ProductTelegramSql) UpdateProduct(productId int64, product Telegram_Market.UpdateProducts) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if product.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *product.Description)
		argId++
	}

	if product.PhotoUrl != nil {
		setValues = append(setValues, fmt.Sprintf("photo_url=$%d", argId))
		args = append(args, *product.PhotoUrl)
		argId++
	}

	if product.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, *product.Price)
		argId++
	}

	if product.Quantity != nil {
		setValues = append(setValues, fmt.Sprintf("quantity=$%d", argId))
		args = append(args, *product.Quantity)
		argId++
	}

	if product.LocationCity != nil {
		setValues = append(setValues, fmt.Sprintf("location_city=$%d", argId))
		args = append(args, *product.LocationCity)
		argId++
	}

	if product.LocationCoordinates != nil {
		setValues = append(setValues, fmt.Sprintf("location_coordinates=$%d", argId))
		args = append(args, *product.LocationCoordinates)
		argId++
	}

	if product.HiddenPhotoUrl != nil {
		setValues = append(setValues, fmt.Sprintf("hidden_photo_url=$%d", argId))
		args = append(args, *product.HiddenPhotoUrl)
		argId++
	}

	if product.HiddenDescription != nil {
		setValues = append(setValues, fmt.Sprintf("hidden_description=$%d", argId))
		args = append(args, *product.HiddenDescription)
		argId++
	}

	if product.PaidFlag != nil {
		setValues = append(setValues, fmt.Sprintf("paid_flag=$%d", argId))
		args = append(args, *product.PaidFlag)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = ?", TableProducts, setQuery)
	args = append(args, productId)
	_, err := s.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to update product from database: %w", err)
	}

	return nil
}

func (s *LocationTelegramSql) DeleteLocation(locationId int64) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = ?`, TableLocations)
	_, err := s.db.Exec(query, locationId)
	if err != nil {
		return fmt.Errorf("failed to delete location from database: %w", err)
	}
	return nil
}
