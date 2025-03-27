package test

import (
	"github.com/shopspring/decimal"
	"unifizeAssignment/internal/models"
	"unifizeAssignment/internal/service/discount"
)

func GetDummyCart() ([]*models.CartItem, *models.PaymentInfo) {
	pumaTShirt := &models.Product{
		ID:           "1",
		Brand:        "PUMA",
		BrandTier:    models.BrandTierRegular,
		Category:     "T-shirts",
		BasePrice:    decimal.NewFromInt(1000),
		CurrentPrice: decimal.NewFromInt(1000),
	}

	cartItems := []*models.CartItem{
		{Product: *pumaTShirt, Quantity: 1, Size: "M"},
	}

	paymentInfo := &models.PaymentInfo{
		Method:      "CARD",
		BankName:    ptr("ICICI"),
		CardType:    ptr("CREDIT"),
		VoucherCode: ptr("SUPER12"),
	}

	return cartItems, paymentInfo
}

func GetDummyDiscount() []discount.DiscountRule {
	return []discount.DiscountRule{
		discount.BrandDiscount{
			Brand:   "PUMA",
			Percent: decimal.NewFromInt(40),
		},
		discount.BankDiscount{
			BankName: "ICICI",
			Percent:  decimal.NewFromInt(5),
		},
		discount.CategoryDiscount{
			Category: "T-shirts",
			Percent:  decimal.NewFromInt(5),
		},
		discount.VoucherDiscount{
			Code:               "SUPER12",
			Percent:            decimal.NewFromInt(12),
			BrandExclusions:    []string{},
			CategoryExclusions: []string{},
			CustomerTier:       "PREMIUM",
		},
	}
}

func ptr(s string) *string {
	return &s
}
