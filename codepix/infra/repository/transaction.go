package repository

import (
	"errors"

	"github.com/LuizEduardoCardozo/imercao-fullcycle/codepix-go/domain/model"
	"github.com/jinzhu/gorm"
)

/* TransactionRepositoryInterface methods for deal with transactions

type TransactionRepositoryInterface interface {
	Register(Transaction *Transacion) error
	Save(Transaction *Transacion) error
	Find(id string) (*Transacion, error)
}

*/

// TransactionRepositoryDB database
type TransactionRepositoryDB struct {
	DB *gorm.DB
}

// Register a new transaction to database
func (r TransactionRepositoryDB) Register(transaction *model.Transacion) error {

	err := r.DB.Create(transaction).Error
	return err

}

// Save a model to database, updates a information
func (r TransactionRepositoryDB) Save(transaction *model.Transacion) error {

	err := r.DB.Save(transaction).Error
	return err

}

// Find a transaction registred on database, based on their id
func (r TransactionRepositoryDB) Find(id string) (*model.Transacion, error) {

	var transaction model.Transacion
	r.DB.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, errors.New("no transaction was found")
	}

	return &transaction, nil

}
