package handlers

import (
	"github.com/gin-gonic/gin"
	"unifizeAssignment/internal/service"
	"unifizeAssignment/internal/test"
)

var handler *DiscountHandler

func LoadHandlers(router *gin.Engine) {
	manager := service.NewDiscountServiceImpl(test.GetDummyDiscount())
	handler = NewDiscountHandler(manager)
	loadRoutes(router)
}

func loadRoutes(router *gin.Engine) {
	group := router.Group("/api/v1/unifize")
	group.POST("/calculate/discount", handler.CalculateDiscount)
	group.POST("/validate/coupon/:couponCode", handler.ValidateCouponCode)
}
