package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"unifizeAssignment/internal/common"
	"unifizeAssignment/internal/models"
	"unifizeAssignment/internal/service"
)

type DiscountHandler struct {
	service service.DiscountService
}

func NewDiscountHandler(service service.DiscountService) *DiscountHandler {
	return &DiscountHandler{service: service}
}

func (d *DiscountHandler) CalculateDiscount(ctx *gin.Context) {
	reqCtx := ctx.Request.Context()
	var request CalculateDiscountRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		common.BadRequest(ctx, "Invalid request payload")
		return
	}

	response, err := d.service.CalculateCartDiscounts(reqCtx, request.CartItems, request.CustomerProfile,
		request.PaymentInfo)
	if err != nil {
		handleError(ctx, err)
		return
	}
	common.Success(ctx, response)
}

func (d *DiscountHandler) ValidateCouponCode(ctx *gin.Context) {
	reqCtx := ctx.Request.Context()
	couponCode := ctx.Param("couponCode")

	var request ValidateCodeRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		common.BadRequest(ctx, "Invalid request payload")
		return
	}

	response, err := d.service.ValidateDiscountCode(reqCtx, couponCode, request.CartItems, request.CustomerProfile)
	if err != nil {
		handleError(ctx, err)
		return
	}
	common.Success(ctx, gin.H{"result": response})
}

func handleError(ctx *gin.Context, err error) {
	var notFoundError *common.NotFoundError
	if errors.As(err, &notFoundError) {
		common.NotFound(ctx, err.Error())
	} else {
		common.BadRequest(ctx, err.Error())
	}
}

type CalculateDiscountRequest struct {
	CartItems       []*models.CartItem      `json:"cartItems"`
	CustomerProfile *models.CustomerProfile `json:"customerProfile"`
	PaymentInfo     *models.PaymentInfo     `json:"paymentInfo"`
}

type ValidateCodeRequest struct {
	CartItems       []*models.CartItem      `json:"cartItems"`
	CustomerProfile *models.CustomerProfile `json:"customerProfile"`
}
