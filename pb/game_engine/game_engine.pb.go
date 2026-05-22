package game_engine

type OutcomeRequest struct {
	GameType  string
	BetAmount int64
	PlayerId  string
}

type OutcomeResponse struct {
	PlayerWins bool
	Payout     int64
	Reason     string
}
