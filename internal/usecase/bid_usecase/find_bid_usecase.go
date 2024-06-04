package bid_usecase

import (
	"context"

	"github.com/osniantonio/fullcycle-auction-go/internal/internal_error"
)

func (bu *BidUseCase) FindByAuctionId(ctx context.Context, auctionId string) ([]BidOutputDTO, *internal_error.InternalError) {
	bidList, err := bu.BidRepository.FindByAuctionId(ctx, auctionId)
	if err != nil {
		return nil, err
	}

	var bidOutputDTOList []BidOutputDTO
	for _, bid := range bidList {
		bidOutputDTOList = append(bidOutputDTOList, BidOutputDTO{
			Id:        bid.Id,
			UserId:    bid.UserId,
			AuctionId: bid.AuctionId,
			Amount:    bid.Amount,
			Timestamp: bid.Timestamp,
		})
	}

	return bidOutputDTOList, nil
}

func (bu *BidUseCase) FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*BidOutputDTO, *internal_error.InternalError) {
	bidEntity, err := bu.BidRepository.FindWinningBidByAuctionId(ctx, auctionId)
	if err != nil {
		return nil, err
	}

	bidOutputDTO := &BidOutputDTO{
		Id:        bidEntity.Id,
		UserId:    bidEntity.UserId,
		AuctionId: bidEntity.AuctionId,
		Amount:    bidEntity.Amount,
		Timestamp: bidEntity.Timestamp,
	}

	return bidOutputDTO, nil
}
