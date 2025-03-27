package service

import (
	"context"
	"unifizeAssignment/internal/models"
)

// Interface for Handling Discount Calculations
type DiscountService interface {
	// Calculate Cart Discounts
	CalculateCartDiscounts(ctx context.Context, cartItems []*models.CartItem,
		customer *models.CustomerProfile, paymentInfo *models.PaymentInfo) (*models.DiscountedPrice, error)

	// Validate Discount Code
	ValidateDiscountCode(ctx context.Context, code string, cartItems []*models.CartItem,
		customer *models.CustomerProfile) (bool, error)
}
