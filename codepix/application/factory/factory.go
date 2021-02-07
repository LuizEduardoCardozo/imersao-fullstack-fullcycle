package factory

import (
	usecase "github.com/LuizEduardoCardozo/imercao-fullcycle/codepix-go/application/usecase"
	repository "github.com/LuizEduardoCardozo/imercao-fullcycle/codepix-go/infra/repository"
	"github.com/jinzhu/gorm"
)

// TransactionUseCaseFactory creates the Transaction Use cases
func TransactionUseCaseFactory(db *gorm.DB) usecase.TransactionUseCase {

	pixRepository := repository.PixKeyRepositoryDB{DB: db}
	transactionRepository := repository.TransactionRepositoryDB{DB: db}

	transactionUseCase := usecase.TransactionUseCase{
		TransactionRepository: &transactionRepository,
		PixRepository:         pixRepository,
	}

	return transactionUseCase

}
