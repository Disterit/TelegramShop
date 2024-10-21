package Telegram_Market

import "errors"

type Users struct {
	Id      int64   `json:"id"`
	UserId  int64   `json:"user_id"`
	Balance float32 `json:"balance"`
}

type Products struct {
	Id                  int64   `json:"id"`
	Description         string  `json:"description" binding:"required"`
	PhotoUrl            string  `json:"photo_url" binding:"required"`
	Price               float32 `json:"price" binding:"required"`
	Quantity            int64   `json:"quantity" binding:"required"`
	LocationCity        string  `json:"location_city" binding:"required"`
	LocationCoordinates string  `json:"location_coordinates" binding:"required"`
	HiddenPhotoUrl      string  `json:"hidden_photo_url"`
	HiddenDescription   string  `json:"hidden_description"`
	PaidFlag            bool    `json:"paid_flag" binding:"required"`
}

type Locations struct {
	Id            int64  `json:"id"`
	Country       string `json:"country"`
	City          string `json:"city"`
	City_district string `json:"city_district"`
}

type Wallet struct {
	ID        int64    `json:"id"`
	Name      string   `json:"name"`      // Например, "Monero"
	Addresses []string `json:"addresses"` // Массив адресов кошелька
}

type UpdateProducts struct {
	Description         *string  `json:"description"`
	PhotoUrl            *string  `json:"photo_url"`
	Price               *float32 `json:"price"`
	Quantity            *int64   `json:"quantity"`
	LocationCity        *string  `json:"location_city"`
	LocationCoordinates *string  `json:"location_coordinates"`
	HiddenPhotoUrl      *string  `json:"hidden_photo_url"`
	HiddenDescription   *string  `json:"hidden_description"`
	PaidFlag            *bool    `json:"paid_flag"`
}

func (i UpdateProducts) Validate() error {
	if i.Description == nil &&
		i.PhotoUrl == nil &&
		i.Price == nil &&
		i.Quantity == nil &&
		i.LocationCity == nil &&
		i.LocationCoordinates == nil &&
		i.HiddenPhotoUrl == nil &&
		i.HiddenDescription == nil &&
		i.PaidFlag == nil {
		return errors.New("either title or description must be provided")
	}

	return nil
}

type UpdateLocations struct {
	Id            *int64  `json:"id"`
	Country       *string `json:"country"`
	City          *string `json:"city"`
	City_district *string `json:"city_district"`
}

func (i UpdateLocations) Validate() error {
	if i.Id == nil &&
		i.Country == nil &&
		i.City == nil &&
		i.City_district == nil {
		return errors.New("either id or country or city must be provided")
	}
	return nil
}
