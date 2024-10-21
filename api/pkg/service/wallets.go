package service

import (
	Telegram_Market "Telegram-Market"
	"Telegram-Market/api/pkg/repository"
)

type WalletTelegramJson struct {
	repo repository.Wallets
}

func NewWalletTelegramJson(repo repository.Wallets) *WalletTelegramJson {
	return &WalletTelegramJson{repo: repo}
}

func (j *WalletTelegramJson) CreateWallet(wallet Telegram_Market.Wallet) (int64, error) {
	return j.repo.CreateWallet(wallet)
}

func (j *WalletTelegramJson) GetAllWallets() ([]Telegram_Market.Wallet, error) {
	return j.repo.GetAllWallets()
}

func (j *WalletTelegramJson) GetWalletById(WalletID int64) (*Telegram_Market.Wallet, error) {
	return j.repo.GetWalletById(WalletID)
}

func (j *WalletTelegramJson) UpdateWallet(walletID int64, Wallet Telegram_Market.Wallet) error {
	return j.repo.UpdateWallet(walletID, Wallet)
}

func (j *WalletTelegramJson) DeleteWallet(WalletID int64) error {
	return j.repo.DeleteWallet(WalletID)
}
