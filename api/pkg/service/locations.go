package service

import (
	Telegram_Market "Telegram-Market"
	"Telegram-Market/api/pkg/repository"
)

type LocationTelegram struct {
	repo repository.Locations
}

func NewLocationTelegram(repo repository.Locations) *LocationTelegram {
	return &LocationTelegram{repo: repo}
}

func (t *LocationTelegram) CreateLocation(location Telegram_Market.Locations) (int64, error) {
	return t.repo.CreateLocation(location)
}

func (t *LocationTelegram) GetLocationById(locationId int64) (Telegram_Market.Locations, error) {
	return t.repo.GetLocationById(locationId)
}

func (t *LocationTelegram) GetAllLocations() ([]Telegram_Market.Locations, error) {
	return t.repo.GetAllLocations()
}

func (t *LocationTelegram) UpdateLocations(locationId int64, location Telegram_Market.UpdateLocations) error {
	return t.repo.UpdateLocations(locationId, location)
}

func (t *LocationTelegram) DeleteLocation(locationId int64) error {
	return t.repo.DeleteLocation(locationId)
}
