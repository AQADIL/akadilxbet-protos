package wallet

import (
	context "context"
)

type GetProfileRequest struct {
	UserId string
}

type GetProfileResponse struct {
	UserId    string
	Username  string
	Email     string
	Balance   int64
	CreatedAt int64
}

type GetBalanceRequest struct {
	UserId string
}

type GetBalanceResponse struct {
	UserId  string
	Balance int64
}

type DepositRequest struct {
	UserId string
	Amount int64
}

type DeductRequest struct {
	UserId string
	Amount int64
}

type BalanceResponse struct {
	UserId     string
	NewBalance int64
}

// WalletServiceServer is the server API for WalletService service.
type WalletServiceServer interface {
	GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error)
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	Deposit(context.Context, *DepositRequest) (*BalanceResponse, error)
	Deduct(context.Context, *DeductRequest) (*BalanceResponse, error)
}

// UnimplementedWalletServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWalletServiceServer struct{}

func (UnimplementedWalletServiceServer) GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error) {
	return nil, nil
}

func (UnimplementedWalletServiceServer) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, nil
}

func (UnimplementedWalletServiceServer) Deposit(context.Context, *DepositRequest) (*BalanceResponse, error) {
	return nil, nil
}

func (UnimplementedWalletServiceServer) Deduct(context.Context, *DeductRequest) (*BalanceResponse, error) {
	return nil, nil
}
