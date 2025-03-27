package discount

import (
	"fmt"
	"github.com/shopspring/decimal"
	"unifizeAssignment/internal/models"
)

/*
* Category Discount Logic
 */
type CategoryDiscount struct {
	Category string
	Percent  decimal.Decimal
}

func (c CategoryDiscount) IsApplicable(cartItems []*models.CartItem, customer *models.CustomerProfile, paymentInfo *models.PaymentInfo) bool {
	for _, item := range cartItems {
		if item.Product.Category == c.Category {
			return true
		}
	}
	return false
}

func (c CategoryDiscount) Apply(cartItems []*models.CartItem, customer *models.CustomerProfile, paymentInfo *models.PaymentInfo) (decimal.Decimal, string) {
	var totalDiscount decimal.Decimal
	for _, item := range cartItems {
		if item.Product.Category == c.Category {
			price := item.Product.CurrentPrice
			disc := c.Percent.Div(decimal.NewFromInt(int64(100)))
			qty := decimal.NewFromInt(int64(item.Quantity))
			discount := price.Mul(disc).Mul(qty)
			totalDiscount = totalDiscount.Add(discount)
		}
	}
	return totalDiscount, fmt.Sprintf("Category Offer: %v", c.Category)
}
