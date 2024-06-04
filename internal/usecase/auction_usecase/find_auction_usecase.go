package auction_usecase

import (
	"context"

	"github.com/osniantonio/fullcycle-auction-go/internal/entity/auction_entity"
	"github.com/osniantonio/fullcycle-auction-go/internal/internal_error"
)

func (au *AuctionUseCase) FindAuctionById(ctx context.Context, id string) (*AuctionOutputDTO, *internal_error.InternalError) {
	auctionEntity, err := au.auctionRepositoryInterface.FindAuctionById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &AuctionOutputDTO{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   ProductCondition(auctionEntity.Condition),
		Status:      AuctionStatus(auctionEntity.Status),
		Timestamp:   auctionEntity.Timestamp,
	}, nil
}

func (au *AuctionUseCase) FindAuctions(
	ctx context.Context,
	status auction_entity.AuctionStatus,
	category, productName string) ([]AuctionOutputDTO, *internal_error.InternalError) {
	auctionEntity, err := au.auctionRepositoryInterface.FindAuctions(ctx, auction_entity.AuctionStatus(status), category, productName)
	if err != nil {
		return nil, err
	}

	var auctionOutputs []AuctionOutputDTO

	for _, value := range auctionEntity {
		auctionOutputs = append(auctionOutputs, AuctionOutputDTO{
			Id:          value.Id,
			ProductName: value.ProductName,
			Category:    value.Category,
			Description: value.Description,
			Condition:   ProductCondition(value.Condition),
			Status:      AuctionStatus(value.Status),
			Timestamp:   value.Timestamp,
		})
	}

	return auctionOutputs, nil
}