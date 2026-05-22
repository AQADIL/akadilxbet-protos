package wallet

type GetProfileRequest struct {
	UserId string
}

type GetProfileResponse struct {
	UserId   string
	Username string
	Email    string
	Balance  int64
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
