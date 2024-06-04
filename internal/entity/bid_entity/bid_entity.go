package bid_entity

import (
	"context"
	"time"

	"github.com/osniantonio/fullcycle-auction-go/internal/internal_error"
)

type Bid struct {
	Id        string
	UserId    string
	AuctionId string
	Amount    float64
	Timestamp time.Time
}

type BidEntityRepository interface {
	CreateBid(ctx context.Context, bidEntities []Bid) *internal_error.InternalError
	FindByAuctionId(ctx context.Context, auctionId string) ([]Bid, *internal_error.InternalError)
	FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*Bid, *internal_error.InternalError)
}
