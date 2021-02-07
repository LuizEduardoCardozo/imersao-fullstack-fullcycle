package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

// PixKeyInterfaceRepository methods
type PixKeyInterfaceRepository interface {
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) (*Bank, error)
	AddAccount(account *Account) (*Account, error)
	FindAccount(id string) (*Account, error)
}

// PixKey Model
type PixKey struct {
	Base      `valid:"required"`
	Kind      string   `json:"kind" valid:"notnull"`
	Key       string   `json:"key" valid:"notnull"`
	AccountID string   `gorm:"column:account_id;type:uuid;not null" valid:"-"`
	Account   *Account `valid:"-"`
	Status    string   `json:"status" valid:"notnull"`
}

func (pixKey *PixKey) isValid() error {

	_, err := govalidator.ValidateStruct(pixKey)

	kKind := pixKey.Kind
	kStatus := pixKey.Status

	if kKind != "email" && kKind != "cpf" {
		return errors.New("Invald type of key")
	}

	if kStatus != "inactive" && kStatus != "active" {
		return errors.New("Invalid status of key")
	}

	if err != nil {
		return err
	}

	return nil

}

// NewPixKey creates a new pix key
func NewPixKey(kind string, account *Account, key string) (*PixKey, error) {

	pixKey := PixKey{
		Kind:    kind,
		Key:     key,
		Account: account,
		Status:  "active",
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()

	err := pixKey.isValid()

	if err != nil {
		return nil, err
	}

	return &pixKey, nil
}
