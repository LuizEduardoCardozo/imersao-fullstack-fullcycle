package repository

import (
	"errors"

	"github.com/LuizEduardoCardozo/imercao-fullcycle/codepix-go/domain/model"
	"github.com/jinzhu/gorm"
)

/* PixKeyInterfaceRepository methods

type PixKeyInterfaceRepository interface {
	AddBank(bank *Bank) (*Bank, error)
	addAccount(account *Account) (*Account, error)
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	findAccound(id string) (*Account, error)
}

*/

// PixKeyRepositoryDB database
type PixKeyRepositoryDB struct {
	DB *gorm.DB
}

// AddBank add a new bank to database
func (r PixKeyRepositoryDB) AddBank(bank *model.Bank) error {

	err := r.DB.Create(bank).Error
	return err

}

// AddAccount add a new account to database
func (r PixKeyRepositoryDB) AddAccount(account *model.Account) error {

	err := r.DB.Create(account).Error
	return err

}

// RegisterKey registar a new pix key to database
func (r PixKeyRepositoryDB) RegisterKey(pixKey *model.PixKey) error {

	err := r.DB.Create(pixKey).Error
	return err

}

// FindKeyByKind finds a registred key, using key, based on kind
func (r PixKeyRepositoryDB) FindKeyByKind(key string, kind string) (*model.PixKey, error) {

	var pixKey model.PixKey
	r.DB.Preload("Account.Bank").First(&pixKey, "key = ? and kind = ?", key, kind)

	if pixKey.ID == "" {
		return nil, errors.New("no key was found")
	}

	return &pixKey, nil

}

// FindAccount finds an account registred on database, based on id
func (r PixKeyRepositoryDB) FindAccount(id string) (*model.Account, error) {

	var account model.Account
	r.DB.Preload("Bank").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, errors.New("no accounts was found")
	}

	return &account, nil

}

// FindBank finds a bank using their id
func (r PixKeyRepositoryDB) FindBank(id string) (*model.Bank, error) {

	var bank model.Bank
	r.DB.First(&bank, "id = ?", id)

	if bank.ID == "" {
		return nil, errors.New("no banks found")
	}

	return &bank, nil

}
