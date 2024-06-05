package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/osniantonio/fullcycle-auction-go/configuration/database/mongodb"
	"github.com/osniantonio/fullcycle-auction-go/internal/infra/api/web/controller/auction_controller"
	"github.com/osniantonio/fullcycle-auction-go/internal/infra/api/web/controller/bid_controller"
	"github.com/osniantonio/fullcycle-auction-go/internal/infra/api/web/controller/user_controller"
	"github.com/osniantonio/fullcycle-auction-go/internal/infra/database/auction"
	"github.com/osniantonio/fullcycle-auction-go/internal/infra/database/bid"
	"github.com/osniantonio/fullcycle-auction-go/internal/infra/database/user"
	"github.com/osniantonio/fullcycle-auction-go/internal/usecase/auction_usecase"
	"github.com/osniantonio/fullcycle-auction-go/internal/usecase/bid_usecase"
	"github.com/osniantonio/fullcycle-auction-go/internal/usecase/user_usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

// rodar o mongo db via docker
// docker container run -d -p 27017:27017 --name auctionsDB mongo

func main() {
	ctx := context.Background()

	if err := godotenv.Load("./.env"); err != nil {
		fmt.Println(err)
		log.Fatal("Error trying to load env variables")
		return
	}

	databaseConnection, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// inicializando as rotas
	router := gin.Default()

	userController, bidController, auctionController := initDependencies(databaseConnection)

	router.GET("/auctions", auctionController.FindAuctions)
	router.GET("/auctions/:auctionId", auctionController.FindAuctionById)
	router.POST("/auctions", auctionController.CreateAucion)
	router.GET("/auction/winner/:auctionId", auctionController.FindWinningBidByAuctionById)
	router.POST("/bid", bidController.CreateBid)
	router.GET("/bid/:auctionId", bidController.FindBidByAuctionId)
	router.GET("/user/:userId", userController.FindUserById)

	router.Run(":8080")
}

func initDependencies(database *mongo.Database) (
	userController *user_controller.UserController,
	bidController *bid_controller.BidController,
	auctionController *auction_controller.AuctionController) {

	auctionRepository := auction.NewAuctionRepository(database)
	bidRepository := bid.NewBidRepository(database, auctionRepository)
	userRepository := user.NewUserRepository(database)

	userController = user_controller.NewUserController(user_usecase.NewUserUseCase(userRepository))
	auctionController = auction_controller.NewAuctionController(auction_usecase.NewAuctionUseCase(auctionRepository, bidRepository))
	bidController = bid_controller.NewBidController(bid_usecase.NewBidUseCase(bidRepository))

	return
}
