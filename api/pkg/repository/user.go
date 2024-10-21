package repository

import (
	Telegram_Market "Telegram-Market"
	"database/sql"
	"fmt"
)

type UserTelegramSql struct {
	db *sql.DB
}

func NewUserTelegramSql(db *sql.DB) *UserTelegramSql {
	return &UserTelegramSql{db: db}
}

func (s *UserTelegramSql) Create(userId int64) error {

	query := fmt.Sprintf("INSERT INTO %s (user_id) VALUES (?)", TableUsers)
	_, err := s.db.Exec(query, userId)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserTelegramSql) GetUserById(userId int64) (Telegram_Market.Users, error) {
	var user Telegram_Market.Users

	query := fmt.Sprintf("SELECT id, user_id, balance FROM %s WHERE id = ?", TableUsers)
	row := s.db.QueryRow(query, userId)

	err := row.Scan(&user.Id, &user.UserId, &user.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user with ID %d not found", userId)
		}
		return user, fmt.Errorf("failed to scan user: %w", err)
	}

	return user, nil
}

func (s *UserTelegramSql) GetUsers() ([]Telegram_Market.Users, error) {
	var users []Telegram_Market.Users

	query := fmt.Sprintf("SELECT * FROM %s", TableUsers)
	rows, err := s.db.Query(query)
	if err != nil {
		return users, err
	}
	for rows.Next() {
		var user Telegram_Market.Users
		err = rows.Scan(&user.Id, &user.UserId, &user.Balance)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (s *UserTelegramSql) DeleteUser(userId int64) error {

	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = ?", TableUsers)
	_, err := s.db.Exec(query, userId)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserTelegramSql) UpdateUser(userId int64, users Telegram_Market.Users) error {
	query := fmt.Sprintf("UPDATE %s SET balance = balance+? WHERE user_id = ?", TableUsers)
	_, err := s.db.Exec(query, users.Balance, userId)

	if err != nil {
		return err
	}

	return nil
}
