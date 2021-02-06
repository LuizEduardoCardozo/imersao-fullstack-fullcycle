package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	TransacionPending    string = "pending"
	TransactionCompleted string = "completed"
	TransactionError     string = "error"
	TransacionConfirmed  string = "confirmed"
)

// Transacion model
type Transacion struct {
	Base              `valid:"required"`
	AccountFrom       *Account `valid:"-"`
	Amount            float64  `json:"amount" vald:"notnull"`
	PixKeyTo          *PixKey  `valid:"-"`
	Status            string   `json:"status" valid:"notnull"`
	Description       string   `json:"description" valid:"notnull"`
	CancelDescription string   `json:"cancel_description" valid:"-"`
}

// TransactionRepositoryInterface methods for deal with transactions
type TransactionRepositoryInterface interface {
	Register(Transaction *Transacion) error
	Save(Transaction *Transacion) error
	Find(id string) (*Transacion, error)
}

// Transactions - A List of transections
type Transactions struct {
	Transactions []*Transacion
}

func (transaction *Transacion) isValid() error {

	_, err := govalidator.ValidateStruct(transaction)

	tAmount := transaction.Amount
	tStatus := transaction.Status
	tID := transaction.ID

	if tAmount <= 0 {
		return errors.New("Tranfer amount incompatible")
	}

	if tStatus != TransacionConfirmed && tStatus != TransacionPending && tStatus != TransactionCompleted && tStatus != TransactionError {
		return errors.New("Invalid type of status")
	}

	if tID == transaction.PixKeyTo.AccountID {
		return errors.New("You cannot send money to yourself")
	}

	if err != nil {
		return nil
	}

	return err

}

// NewTransaction - Creates a new transaction
func NewTransaction(accountFrom *Account, amount float64, pixKeyTo *PixKey, description string) (*Transacion, error) {

	transaction := Transacion{
		AccountFrom: accountFrom,
		Amount:      amount,
		PixKeyTo:    pixKeyTo,
		Description: description,
		Status:      TransacionPending,
	}

	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()

	err := transaction.isValid()

	if err != nil {
		return nil, err
	}

	return &transaction, nil

}

// Completed change the transaction status for "completed"
func (transaction *Transacion) Completed() error {

	transaction.Status = TransactionCompleted
	transaction.UpdatedAt = time.Now()

	err := transaction.isValid()
	return err

}

// Confirm change the transaction status for "confirmed"
func (transaction *Transacion) Confirm() error {

	transaction.Status = TransacionConfirmed
	transaction.UpdatedAt = time.Now()

	err := transaction.isValid()
	return err

}

// Cancel change the transaction status for "canceled"
func (transaction *Transacion) Cancel(description string) error {

	transaction.Status = TransactionError
	transaction.UpdatedAt = time.Now()
	transaction.Description = description

	err := transaction.isValid()
	return err

}
