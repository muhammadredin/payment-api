package service

import (
	"PaymentAPI/entity"
	"fmt"
	"github.com/stretchr/testify/mock"
)

type WalletServiceMock struct {
	mock.Mock
}

func (w *WalletServiceMock) CreateWallet(customerId string) error {
	args := w.Called(customerId)

	return args.Error(0)
}

func (w *WalletServiceMock) GetWallet(customerId string) (entity.Wallet, error) {
	args := w.Called(customerId)

	wallet, ok := args.Get(0).(entity.Wallet)
	if !ok {
		return entity.Wallet{}, fmt.Errorf("invalid type for Wallet")
	}
	return wallet, args.Error(1)
}

func (w *WalletServiceMock) UpdateWallet(customerId string, balance float64) error {
	args := w.Called(customerId, balance)
	return args.Error(0)
}