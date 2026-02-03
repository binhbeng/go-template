package handler

import (
	"github.com/binhbeng/goex/internal/dto"
	"github.com/binhbeng/goex/internal/service"
	"github.com/binhbeng/goex/internal/utils"
	"github.com/binhbeng/goex/internal/validation"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService *service.OrderService
}

func NewOrderHandler(orderService *service.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

func (h *OrderHandler) GetListOrder(c *gin.Context) {
	ctx := c.Request.Context()
	var req dto.QueryOrdersInput
	if err := validation.ValidateQueryParams(c, &req); err != nil {
		return
	}

	orders, total, err := h.orderService.GetListOrders(ctx, req)

	if err != nil {
		utils.HttpBadRequest(c, "get failed", err)
		return
	}
	
	paginationResp := utils.NewPagiantionResponse(orders, req.Page, req.Limit, total)
	utils.SuccessResponse(c, 200, "OK", paginationResp)
}
