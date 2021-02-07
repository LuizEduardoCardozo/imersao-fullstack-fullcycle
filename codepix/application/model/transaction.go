package model

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Transaction model, used for json parsing
type Transaction struct {
	ID           string  `json:"id" validate:"required,uuid4"`
	AccountID    string  `json:"accountId" validate:"required,uuid4"`
	Amount       float64 `json:"amount" validate:"required,numeric"`
	PixKeyTo     string  `json:"pixKeyTo" validate:"required"`
	PixKeyKindTo string  `json:"pixKeyKindTo" validate:"required"`
	Description  string  `json:"description" validate:"required"`
	Status       string  `json:"status" validate:"-"`
	Error        string  `json:"error"`
}

func (t *Transaction) isValid() error {

	v := validator.New()
	err := v.Struct(t)

	if err != nil {
		fmt.Printf("error during Transaction Validation: %s", err.Error())
		return err
	}

	return nil

}

// ParseJSON converts the go struct from a Json Format
func (t *Transaction) ParseJSON(data []byte) error {

	err := json.Unmarshal(data, t)

	if err != nil {
		return err
	}

	err = t.isValid()

	if err != nil {
		return err
	}

	return nil

}

// ToJSON Converts a struct to a JSON
func (t *Transaction) ToJSON() ([]byte, error) {

	err := t.isValid()

	if err != nil {
		return nil, err
	}

	result, err := json.Marshal(t)

	if err != nil {
		return nil, err
	}

	return result, nil

}

// NewTransaction returns a new empty transaction struct
func NewTransaction() *Transaction {
	return &Transaction{}
}
