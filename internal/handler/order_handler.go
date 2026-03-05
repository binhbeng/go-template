package handler

import (
	"github.com/binhbeng/goex/internal/dto"
	"github.com/binhbeng/goex/internal/service"
	"github.com/binhbeng/goex/internal/utils"
	"github.com/binhbeng/goex/internal/validation"
	"github.com/binhbeng/goex/pkg/kafka"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService  *service.OrderService
	kafkaProducer kafka.KafkaProducer
}

func NewOrderHandler(orderService *service.OrderService, kafkaProducer kafka.KafkaProducer) *OrderHandler {
	return &OrderHandler{
		orderService:  orderService,
		kafkaProducer: kafkaProducer,
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

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	ctx := c.Request.Context()
	userId := 1
	var data dto.CreateOrderInput
	if err := validation.ValidateBodyParams(c, &data); err != nil {
		return
	}

	h.kafkaProducer.Produce(ctx , "order", data)

	order, err := h.orderService.CreateOrder(ctx, userId, data)
	if err != nil {
		utils.HttpBadRequest(c, "create failed", err)
		return
	}

	utils.SuccessResponse(c, 200, "OK", order)
}
