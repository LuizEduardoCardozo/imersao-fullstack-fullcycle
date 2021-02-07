package usecase

import (
	"errors"

	"github.com/LuizEduardoCardozo/imercao-fullcycle/codepix-go/domain/model"
)

// PixUseCase struct that mades interface with the database
type PixUseCase struct {
	PixKeyRepository model.PixKeyInterfaceRepository
}

// RegisterKey - register a new key
func (p *PixUseCase) RegisterKey(key string, kind, string, accountID string) (*model.PixKey, error) {

	account, err := p.PixKeyRepository.FindAccount(accountID)

	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, account, key)

	if err != nil {
		return nil, err
	}

	p.PixKeyRepository.RegisterKey(pixKey)

	if pixKey.ID == "" {
		return nil, errors.New("unable to create new key at the moment")
	}

	return pixKey, nil

}

// FindKey - find a key
func (p *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := p.PixKeyRepository.FindKeyByKind(key, kind)

	if err != nil {
		return nil, err
	}

	return pixKey, nil

}
