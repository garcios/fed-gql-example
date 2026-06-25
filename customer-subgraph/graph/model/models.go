package model

type Customer struct {
	ID string `json:"id"`
}

func (Customer) IsEntity() {}
