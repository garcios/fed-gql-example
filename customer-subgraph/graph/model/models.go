package model

type Customer struct {
	ID string `json:"id"`
}

func (Customer) IsEntity() {}

type Transaction struct {
	CustomerID string `json:"customerID"`
}

type Bet struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
	GameID string  `json:"gameID"`
}

type Game struct {
	ID           string `json:"id"`
	MarketTypeID string `json:"marketTypeID"`
}

func (Game) IsEntity() {}

