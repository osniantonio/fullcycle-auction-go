package auction_usecase

import (
	"context"
	"time"

	"github.com/osniantonio/fullcycle-auction-go/internal/entity/auction_entity"
	"github.com/osniantonio/fullcycle-auction-go/internal/entity/bid_entity"
	"github.com/osniantonio/fullcycle-auction-go/internal/internal_error"
	"github.com/osniantonio/fullcycle-auction-go/internal/usecase/bid_usecase"
)

type AuctionInputDTO struct {
	ProductName string           `json:"product_name" binding:"required,min=1"`
	Category    string           `json:"category" binding:"required,min=2"`
	Description string           `json:"description" binding:"required,min=10,max=200"`
	Condition   ProductCondition `json:"condition"`
}

type AuctionOutputDTO struct {
	Id          string           `json:"id"`
	ProductName string           `json:"product_name"`
	Category    string           `json:"category"`
	Description string           `json:"description"`
	Condition   ProductCondition `json:"condition"`
	Status      AuctionStatus    `json:"status"`
	Timestamp   time.Time        `json:"timestamp"`
}

type WinningInfoOutputDTO struct {
	Auction AuctionOutputDTO          `json:"auction"`
	Bid     *bid_usecase.BidOutputDTO `json:"bid,omitempty"`
}

type ProductCondition int64
type AuctionStatus int64

type AuctionUseCase struct {
	auctionRepositoryInterface auction_entity.AuctionRepositoryInterface
	bidRepositoryInterface     bid_entity.BidEntityRepository
}

func (au *AuctionUseCase) CreateAuction(ctx context.Context, auctionInputDTO AuctionInputDTO) *internal_error.InternalError {
	auction, err := auction_entity.CreateAuction(auctionInputDTO.ProductName, auctionInputDTO.Category, auctionInputDTO.Description, auction_entity.ProductCondition(auctionInputDTO.Condition))

	if err != nil {
		return err
	}

	if err := au.auctionRepositoryInterface.CreateAuction(ctx, auction); err != nil {
		return err
	}

	return nil
}
