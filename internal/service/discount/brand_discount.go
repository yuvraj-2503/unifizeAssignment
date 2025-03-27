package discount

import (
	"fmt"
	"github.com/shopspring/decimal"
	"unifizeAssignment/internal/models"
)

/*
* Brand Discount Logic
 */
type BrandDiscount struct {
	Brand   string
	Percent decimal.Decimal
}

func (b BrandDiscount) IsApplicable(cartItems []*models.CartItem, customer *models.CustomerProfile, paymentInfo *models.PaymentInfo) bool {
	for _, item := range cartItems {
		if item.Product.Brand == b.Brand {
			return true
		}
	}
	return false
}

func (b BrandDiscount) Apply(cartItems []*models.CartItem, customer *models.CustomerProfile, paymentInfo *models.PaymentInfo) (decimal.Decimal, string) {
	var totalDiscount decimal.Decimal
	for _, item := range cartItems {
		if item.Product.Brand == b.Brand {
			price := item.Product.CurrentPrice
			disc := b.Percent.Div(decimal.NewFromInt(int64(100)))
			qty := decimal.NewFromInt(int64(item.Quantity))
			discount := price.Mul(disc).Mul(qty)
			totalDiscount = totalDiscount.Add(discount)
		}
	}
	return totalDiscount, fmt.Sprintf("Brand Offer: %v", b.Brand)
}
