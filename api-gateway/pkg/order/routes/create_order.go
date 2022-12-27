package routes

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/allrested/product/api-gateway/pkg/order/pb"
	"net/http"
)

type CreateOrderRequestBody struct {
	ProductId int64 `json:"productId"`
	Quantity  int64 `json:"quantity"`
}

func CreateOrder(ctx *gin.Context, c pb.OrderServiceClient) {
	b := CreateOrderRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, ok := ctx.Get("userId")
	if !ok {
		_ = ctx.AbortWithError(http.StatusNotFound, errors.New("user not found"))
		return
	}

	res, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		ProductId: b.ProductId,
		Quantity:  b.Quantity,
		UserId:    userId.(int64),
	})

	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
