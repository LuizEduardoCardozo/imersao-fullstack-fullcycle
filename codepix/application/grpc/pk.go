package grpc

import (
	"context"

	"github.com/LuizEduardoCardozo/imercao-fullcycle/codepix-go/application/grpc/pb"
	"github.com/LuizEduardoCardozo/imercao-fullcycle/codepix-go/application/usecase"
)

// PixGrpcService usecases for Pix Services
type PixGrpcService struct {
	PixUseCase usecase.PixUseCase
	pb.UnimplementedPixServiceServer
}

// RegisterPixKey register a new pix key
func (p *PixGrpcService) RegisterPixKey(ctx context.Context, in *pb.PixKeyRegistration) (*pb.PixKeyCreatedResult, error) {

	key, err := p.PixUseCase.RegisterKey(in.Key, in.Kind, "", in.Accont)

	if err != nil {
		return &pb.PixKeyCreatedResult{
			Status: "not created",
			Error:  err.Error(),
		}, err
	}

	return &pb.PixKeyCreatedResult{
		Status: "Created",
		Id:     key.ID,
	}, nil

}

// Find a registred Pix Key
func (p *PixGrpcService) Find(ctx context.Context, in *pb.PixKey) (*pb.PixKeyInfo, error) {

	key, err := p.PixUseCase.FindKey(in.Key, in.Key)

	if err != nil {
		return &pb.PixKeyInfo{}, err
	}

	return &pb.PixKeyInfo{
		Id:   key.ID,
		Key:  key.Key,
		Kind: key.Kind,
		Account: &pb.Account{
			AccountId:     key.AccountID,
			AccountNumber: key.Account.Number,
			BankId:        key.Account.BankID,
			OwnerName:     key.Account.OwnerName,
			CreatedAt:     key.Account.CreatedAt.String(),
			BankName:      key.Account.Bank.Name,
		},
		CreatedAt: key.CreatedAt.String(),
	}, err

}

// NewPixGrpcService returns PixGrpcService
func NewPixGrpcService(usecase usecase.PixUseCase) *PixGrpcService {
	return &PixGrpcService{
		PixUseCase: usecase,
	}
}
