package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/allrested/product/api-gateway/pkg/product/pb"
	"net/http"
	"strconv"
)

func FineOne(ctx *gin.Context, c pb.ProductServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)

	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	res, er := c.FindOne(context.Background(), &pb.FindOneRequest{
		Id: id,
	})

	if er != nil {
		_ = ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
