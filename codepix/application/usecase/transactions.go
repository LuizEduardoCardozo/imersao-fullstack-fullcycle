package usecase

import (
	"errors"

	"github.com/LuizEduardoCardozo/imercao-fullcycle/codepix-go/domain/model"
)

/*

type TransactionRepositoryInterface interface {
	Register(Transaction *Transacion) error
	Save(Transaction *Transacion) error
	Find(id string) (*Transacion, error)
}

*/

// TransactionUseCase struct that mades interface with the database
type TransactionUseCase struct {
	TransactionRepository model.TransactionRepositoryInterface
	PixRepository         model.PixKeyInterfaceRepository
}

// Register register a new transaction into database
func (t *TransactionUseCase) Register(accountID string, amount float64, pixKeyTo string, pixKeyKind string, description string) (*model.Transacion, error) {

	account, err := t.PixRepository.FindAccount(accountID)

	if err != nil {
		return nil, err
	}

	pixKey, err := t.PixRepository.FindKeyByKind(pixKeyTo, pixKeyKind)

	if err != nil {
		return nil, err
	}

	transaction, err := model.NewTransaction(account, amount, pixKey, description)

	if err != nil {
		return nil, err
	}

	t.TransactionRepository.Save(transaction)

	if transaction.ID == "" {
		return nil, errors.New("unable to process this transaction now")
	}

	return transaction, nil

}

// Confirm mark a transaction as "confirmed"
func (t *TransactionUseCase) Confirm(transactionID string) (*model.Transacion, error) {

	transaction, err := t.TransactionRepository.Find(transactionID)

	if err != nil {
		return nil, err
	}

	transaction.Status = "confirmed"

	err = t.TransactionRepository.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

// Completed mark a function as "completed"
func (t *TransactionUseCase) Completed(transactionID string) (*model.Transacion, error) {

	transaction, err := t.TransactionRepository.Find(transactionID)

	if err != nil {
		return nil, err
	}

	transaction.Status = "Completed"

	err = t.TransactionRepository.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, err

}

// Cancel mark a transaction as "canceled", and register the reason why
func (t *TransactionUseCase) Error(transactionID string, reason string) (*model.Transacion, error) {

	transaction, err := t.TransactionRepository.Find(transactionID)

	if err != nil {
		return nil, err
	}

	transaction.Status = "Error"
	transaction.CancelDescription = reason

	err = t.TransactionRepository.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil

}
