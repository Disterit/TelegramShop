package repository

import (
	Telegram_Market "Telegram-Market"
	"Telegram-Market/api/pkg/repository/kafkaRepository"
	"database/sql"
	"fmt"
	"github.com/segmentio/kafka-go"
)

type UserTelegramSql struct {
	db     *sql.DB
	writer *kafka.Writer
}

func NewUserTelegramSql(db *sql.DB, writer *kafka.Writer) *UserTelegramSql {
	return &UserTelegramSql{db: db, writer: writer}
}

func (s *UserTelegramSql) CreateUser(userId int64) error {
	const op = "api.pkg.repository.user.CreateUser()"

	query := fmt.Sprintf("INSERT INTO %s (user_id) VALUES (?)", TableUsers)
	_, err := s.db.Exec(query, userId)
	if err != nil {
		kafkaRepository.KafkaResponse(s.writer, err.Error(), op)
		return err
	}

	return nil
}

func (s *UserTelegramSql) GetUserById(userId int64) (Telegram_Market.Users, error) {
	const op = "api.pkg.repository.user.GetUserById()"
	var user Telegram_Market.Users

	query := fmt.Sprintf("SELECT id, user_id, balance FROM %s WHERE user_id = ?", TableUsers)
	row := s.db.QueryRow(query, userId)

	err := row.Scan(&user.Id, &user.UserId, &user.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			kafkaRepository.KafkaResponse(s.writer, err.Error(), op)
			return user, fmt.Errorf("user with ID %d not found", userId)
		}
		kafkaRepository.KafkaResponse(s.writer, err.Error(), op)
		return user, fmt.Errorf("failed to scan user: %w", err)
	}

	return user, nil
}

func (s *UserTelegramSql) GetUsers() ([]Telegram_Market.Users, error) {
	const op = "api.pkg.repository.user.GetUsers()"
	var users []Telegram_Market.Users

	query := fmt.Sprintf("SELECT * FROM %s", TableUsers)
	rows, err := s.db.Query(query)
	if err != nil {
		kafkaRepository.KafkaResponse(s.writer, err.Error(), op)
		return users, err
	}

	for rows.Next() {
		var user Telegram_Market.Users
		err = rows.Scan(&user.Id, &user.UserId, &user.Balance)
		if err != nil {
			kafkaRepository.KafkaResponse(s.writer, err.Error(), op)
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *UserTelegramSql) DeleteUser(userId int64) error {
	const op = "api.pkg.repository.user.DeleteUser()"

	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = ?", TableUsers)

	_, err := s.db.Exec(query, userId)
	if err != nil {
		kafkaRepository.KafkaResponse(s.writer, err.Error(), op)
		return err
	}

	return nil
}

func (s *UserTelegramSql) UpdateUser(userId int64, users Telegram_Market.Users) error {
	const op = "api.pkg.repository.user.UpdateUser()"
	query := fmt.Sprintf("UPDATE %s SET balance = balance+? WHERE user_id = ?", TableUsers)
	_, err := s.db.Exec(query, users.Balance, userId)

	if err != nil {
		kafkaRepository.KafkaResponse(s.writer, err.Error(), op)
		return err
	}

	return nil
}
