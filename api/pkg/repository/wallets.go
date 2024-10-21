package repository

import (
	Telegram_Market "Telegram-Market"
	"encoding/json"
	"fmt"
	"os"
)

type WalletsTelegramJson struct {
	filePath string
}

func NewWalletsTelegramJson(filePath string) *WalletsTelegramJson {
	return &WalletsTelegramJson{filePath: filePath}
}

func (w *WalletsTelegramJson) generateID() (int64, error) {
	wallets, err := w.GetAllWallets()
	if err != nil {
		return 0, err
	}

	if len(wallets) == 0 {
		return 1, nil
	}

	maxID := int64(0)
	for _, wallet := range wallets {
		if wallet.ID > maxID {
			maxID = wallet.ID
		}
	}

	return maxID + 1, nil
}

func (w *WalletsTelegramJson) CreateWallet(wallet Telegram_Market.Wallet) (int64, error) {
	MutexWalletWrite.Lock()
	defer MutexWalletWrite.Unlock()

	id, err := w.generateID()
	if err != nil {
		return 0, err
	}
	wallet.ID = id

	wallets, err := w.GetAllWallets()
	if err != nil {
		return 0, err
	}

	wallets = append(wallets, wallet)

	file, err := os.Create(w.filePath)
	if err != nil {
		return 0, fmt.Errorf("error to open JSON file: %v", err)
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(wallets, "", "    ")
	if err != nil {
		return 0, fmt.Errorf("error to coding JSON: %v", err)
	}

	if _, err := file.Write(jsonData); err != nil {
		return 0, fmt.Errorf("error to wtiring JSON file: %v", err)
	}

	return id, nil
}

func (w *WalletsTelegramJson) GetAllWallets() ([]Telegram_Market.Wallet, error) {
	MutexWalletRead.Lock()
	defer MutexWalletRead.Unlock()
	file, err := os.Open(w.filePath)
	if err != nil {
		return nil, fmt.Errorf("error to open JSON file: %v", err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("error to get info: %v", err)
	}

	if stat.Size() == 0 {
		return []Telegram_Market.Wallet{}, nil
	}

	var wallets []Telegram_Market.Wallet
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&wallets); err != nil {
		return nil, fmt.Errorf("error to reading JSON file: %v", err)
	}

	return wallets, nil
}

func (w *WalletsTelegramJson) GetWalletById(id int64) (*Telegram_Market.Wallet, error) {
	wallets, err := w.GetAllWallets()
	if err != nil {
		return nil, err
	}

	for _, wallet := range wallets {
		if wallet.ID == id {
			return &wallet, nil
		}
	}

	return nil, fmt.Errorf("wallet with ID %d not found", id)
}

func (w *WalletsTelegramJson) DeleteWallet(WalletID int64) error {
	MutexWalletDelete.Lock()
	defer MutexWalletDelete.Unlock()

	wallets, err := w.GetAllWallets()
	if err != nil {
		return err
	}

	index := -1
	for i, wallet := range wallets {
		if wallet.ID == WalletID {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("wallet with ID %d not found", WalletID)
	}

	wallets = append(wallets[:index], wallets[index+1:]...)

	file, err := os.Create(w.filePath)
	if err != nil {
		return fmt.Errorf("error to open JSON file: %v", err)
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(wallets, "", "    ")
	if err != nil {
		return fmt.Errorf("error to coding JSON: %v", err)
	}

	if _, err := file.Write(jsonData); err != nil {
		return fmt.Errorf("error to wtiring JSON file: %v", err)
	}

	return nil
}

func (w *WalletsTelegramJson) UpdateWallet(walletID int64, Wallet Telegram_Market.Wallet) error {
	wallets, err := w.GetAllWallets()
	if err != nil {
		return err
	}

	index := -1
	for i, wallet := range wallets {
		if wallet.ID == walletID {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("wallet with ID %d not found", walletID)
	}

	if Wallet.Name != "" {
		wallets[index].Name = Wallet.Name
	}
	if len(Wallet.Addresses) > 0 {
		wallets[index].Addresses = Wallet.Addresses
	}

	file, err := os.Create(w.filePath)
	if err != nil {
		return fmt.Errorf("error to open JSON file: %v", err)
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(wallets, "", "    ")
	if err != nil {
		return fmt.Errorf("error to coding JSON: %v", err)
	}

	if _, err := file.Write(jsonData); err != nil {
		return fmt.Errorf("error to wtiring JSON file: %v", err)
	}

	return nil
}