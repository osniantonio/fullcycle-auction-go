package auction_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/osniantonio/fullcycle-auction-go/configuration/rest_err"
	"github.com/osniantonio/fullcycle-auction-go/internal/infra/api/web/validation"
	"github.com/osniantonio/fullcycle-auction-go/internal/usecase/auction_usecase"
)

type auctionController struct {
	auctionUseCase auction_usecase.AuctionUseCase
}

func NewAuctionController(auctionUseCase auction_usecase.AuctionUseCase) *auctionController {
	return &auctionController{
		auctionUseCase: auctionUseCase,
	}
}

func (u *auctionController) CreateAuciont(c *gin.Context) {
	var auctionInputDTO auction_usecase.AuctionInputDTO

	if err := c.ShouldBindJSON(&auctionInputDTO); err != nil {
		restErr := validation.ValidateErr(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	err := u.auctionUseCase.CreateAuction(context.Background(), auctionInputDTO)
	if err != nil {
		restErr := rest_err.ConvertError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	c.Status(http.StatusCreated)
}
