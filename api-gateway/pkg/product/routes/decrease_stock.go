package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/allrested/product/api-gateway/pkg/product/pb"
	"net/http"
	"strconv"
)

func DecreaseStock(ctx *gin.Context, c pb.ProductServiceClient) {
	orderId, err := strconv.ParseInt(ctx.Param("order_id"), 10, 32)

	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	res, er := c.DecreaseStock(context.Background(), &pb.DecreaseStockRequest{
		OrderId: orderId,
	})

	if er != nil {
		_ = ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
