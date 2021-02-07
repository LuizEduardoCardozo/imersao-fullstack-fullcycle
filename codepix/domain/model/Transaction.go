package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	transacionPending    string = "pending"
	transactionCompleted string = "completed"
	transactionError     string = "error"
	transacionConfirmed  string = "confirmed"
)

// Transacion model
type Transacion struct {
	Base              `valid:"required"`
	AccountFrom       *Account `valid:"-"`
	AccountFromID     string   `gorm:"column:account_from_id;type:uui;not null" valid:"notnull"`
	Amount            float64  `json:"amount" gorm:"type:float" vald:"notnull"`
	PixKeyTo          *PixKey  `valid:"-"`
	PixKeyToID        string   `gorm:"column:pix_key_id_to;type:uuid;not null" valid:"notnull"`
	Status            string   `json:"status" gorm:"type:varchar(20)" valid:"notnull"`
	Description       string   `json:"description" gorm:"type:varchar(255)" valid:"notnull"`
	CancelDescription string   `json:"cancel_description" gorm:"type:varchar(255)" valid:"-"`
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

	if tStatus != transacionConfirmed && tStatus != transacionPending && tStatus != transactionCompleted && tStatus != transactionError {
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
		Status:      transacionPending,
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

	transaction.Status = transactionCompleted
	transaction.UpdatedAt = time.Now()

	err := transaction.isValid()
	return err

}

// Confirm change the transaction status for "confirmed"
func (transaction *Transacion) Confirm() error {

	transaction.Status = transacionConfirmed
	transaction.UpdatedAt = time.Now()

	err := transaction.isValid()
	return err

}

// Cancel change the transaction status for "canceled"
func (transaction *Transacion) Cancel(description string) error {

	transaction.Status = transactionError
	transaction.UpdatedAt = time.Now()
	transaction.Description = description

	err := transaction.isValid()
	return err

}
