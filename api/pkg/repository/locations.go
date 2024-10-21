package repository

import (
	Telegram_Market "Telegram-Market"
	"database/sql"
	"fmt"
	"strings"
)

type LocationTelegramSql struct {
	db *sql.DB
}

func NewLocationTelegramSql(db *sql.DB) *LocationTelegramSql {
	return &LocationTelegramSql{db: db}
}

func (s *LocationTelegramSql) CreateLocation(location Telegram_Market.Locations) (int64, error) {
	query := fmt.Sprintf("INSERT INTO %s (country, city, city_district) VALUES ($1, $2, $3)", TableLocations)
	res, err := s.db.Exec(query,
		location.Country,
		location.City,
		location.City_district,
	)

	fmt.Println(err)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}

	return id, nil
}

func (s *LocationTelegramSql) GetLocationById(locationId int64) (Telegram_Market.Locations, error) {
	var location Telegram_Market.Locations

	query := fmt.Sprintf("SELECT id, country, city, city_district FROM %s WHERE id = $1", TableLocations)
	row := s.db.QueryRow(query, locationId)

	err := row.Scan(
		&location.Id,
		&location.Country,
		&location.City,
		&location.City_district,
	)
	if err != nil {
		return location, err
	}

	return location, nil
}

func (s *LocationTelegramSql) GetAllLocations() ([]Telegram_Market.Locations, error) {
	var locations []Telegram_Market.Locations
	query := fmt.Sprintf("SELECT id, country, city, city_district FROM %s", TableLocations)
	rows, err := s.db.Query(query)
	if err != nil {
		return locations, err
	}
	for rows.Next() {
		var location Telegram_Market.Locations
		err = rows.Scan(
			&location.Id,
			&location.Country,
			&location.City,
			&location.City_district,
		)

		if err != nil {
			return locations, err
		}
		locations = append(locations, location)
	}

	return locations, nil
}

func (s *LocationTelegramSql) UpdateLocations(locationId int64, location Telegram_Market.UpdateLocations) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if location.Country != nil {
		setValues = append(setValues, fmt.Sprintf("country=$%d", argId))
		args = append(args, *location.Country)
		argId++
	}

	if location.City != nil {
		setValues = append(setValues, fmt.Sprintf("city=$%d", argId))
		args = append(args, *location.City)
		argId++
	}

	if location.City_district != nil {
		setValues = append(setValues, fmt.Sprintf("city_district=$%d", argId))
		args = append(args, *location.City_district)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = ?", TableLocations, setQuery)
	args = append(args, locationId)
	_, err := s.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to update product from database: %w", err)
	}

	return nil
}

//id INTEGER PRIMARY KEY AUTOINCREMENT,
//country TEXT NOT NULL, -- страна
//city TEXT NOT NULL, -- город
//city_district TEXT NOT NULL -- район города
