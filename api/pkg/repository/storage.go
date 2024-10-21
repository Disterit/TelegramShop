package repository

import (
	Telegram_Market "Telegram-Market"
	"database/sql"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
)

type User interface {
	CreateUser(userId int64) error
	GetUserById(userId int64) (Telegram_Market.Users, error)
	GetUsers() ([]Telegram_Market.Users, error)
	UpdateUser(userId int64, users Telegram_Market.Users) error
	DeleteUser(userId int64) error
}

type Products interface {
	CreateProduct(product Telegram_Market.Products) (int64, error)
	GetProductById(productId int64) (Telegram_Market.Products, error)
	GetAllProducts() ([]Telegram_Market.Products, error)
	UpdateProduct(productId int64, product Telegram_Market.UpdateProducts) error
	DeleteProduct(productId int64) error
}

type Locations interface {
	CreateLocation(location Telegram_Market.Locations) (int64, error)
	GetLocationById(locationId int64) (Telegram_Market.Locations, error)
	GetAllLocations() ([]Telegram_Market.Locations, error)
	UpdateLocations(locationId int64, location Telegram_Market.UpdateLocations) error
	DeleteLocation(locationId int64) error
}

type Wallets interface {
	CreateWallet(wallet Telegram_Market.Wallet) (int64, error)
	GetAllWallets() ([]Telegram_Market.Wallet, error)
	GetWalletById(WalletID int64) (*Telegram_Market.Wallet, error)
	UpdateWallet(walletID int64, Wallet Telegram_Market.Wallet) error
	DeleteWallet(WalletID int64) error
}

type Repository struct {
	User
	Products
	Locations
	Wallets
}

func NewStorage(db *sql.DB, writer *kafka.Writer) *Repository {
	return &Repository{
		User:      NewUserTelegramSql(db, writer),
		Products:  NewProductTelegramSql(db, writer),
		Locations: NewLocationTelegramSql(db, writer),
		Wallets:   NewWalletsTelegramJson(viper.GetString("jsonpath"), writer),
	}
}
