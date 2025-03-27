package discount

import (
	"fmt"
	"github.com/shopspring/decimal"
	"unifizeAssignment/internal/models"
)

/*
* Bank Discount Logic
 */
type BankDiscount struct {
	BankName string
	Percent  decimal.Decimal
}

func (b BankDiscount) IsApplicable(cartItems []*models.CartItem, customer *models.CustomerProfile, paymentInfo *models.PaymentInfo) bool {
	return paymentInfo != nil && paymentInfo.BankName != nil && *paymentInfo.BankName == b.BankName
}

func (b BankDiscount) Apply(cartItems []*models.CartItem, customer *models.CustomerProfile, paymentInfo *models.PaymentInfo) (decimal.Decimal, string) {
	if paymentInfo == nil || paymentInfo.BankName == nil || *paymentInfo.BankName != b.BankName {
		return decimal.Zero, ""
	}

	var totalPrice decimal.Decimal
	for _, item := range cartItems {
		totalPrice = totalPrice.Add(item.Product.CurrentPrice.Mul(decimal.NewFromInt(int64(item.Quantity))))
	}

	discount := totalPrice.Mul(b.Percent).Div(decimal.NewFromInt(100))
	return discount, fmt.Sprintf("Bank Offer: %v", b.BankName)
}
