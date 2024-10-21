package service

import (
	Telegram_Market "Telegram-Market"
	"Telegram-Market/api/pkg/repository"
)

type UserTelegram struct {
	repo repository.User
}

func NewUserTelegram(repo repository.User) *UserTelegram {
	return &UserTelegram{repo: repo}
}

func (t *UserTelegram) Create(userId int64) error {
	return t.repo.Create(userId)
}

func (t *UserTelegram) GetUserById(userId int64) (Telegram_Market.Users, error) {
	return t.repo.GetUserById(userId)
}

func (t *UserTelegram) GetUsers() ([]Telegram_Market.Users, error) {
	return t.repo.GetUsers()
}

func (t *UserTelegram) UpdateUser(userId int64, users Telegram_Market.Users) error {
	return t.repo.UpdateUser(userId, users)
}

func (t *UserTelegram) DeleteUser(userId int64) error {
	return t.repo.DeleteUser(userId)
}
