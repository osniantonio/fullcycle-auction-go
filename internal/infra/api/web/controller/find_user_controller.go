package user_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/osniantonio/fullcycle-auction-go/configuration/rest_err"
	"github.com/osniantonio/fullcycle-auction-go/internal/usecase/user_usecase"
)

type userController struct {
	userUseCase user_usecase.UserUseCase
}

func NewUserController(userUseCase user_usecase.UserUseCase) *userController {
	return &userController{
		userUseCase: userUseCase,
	}
}

func (u *userController) FindUserById(c *gin.Context) {
	userId := c.Query("userId")

	if err := uuid.Validate(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid fields", rest_err.Causes{
			Field:   "userId",
			Message: "Invalid UUID value",
		})

		c.JSON(errRest.Code, errRest)
		return
	}

	userData, err := u.userUseCase.FindUserById(context.Background(), userId)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, userData)
}
