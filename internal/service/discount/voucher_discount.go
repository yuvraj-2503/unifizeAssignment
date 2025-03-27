package discount

import (
	"fmt"
	"github.com/shopspring/decimal"
	"unifizeAssignment/internal/models"
)

/*
* Voucher Discount Logic
 */
type VoucherDiscount struct {
	Code               string
	Percent            decimal.Decimal
	BrandExclusions    []string
	CategoryExclusions []string
	CustomerTier       string
}

func (v VoucherDiscount) IsApplicable(cartItems []*models.CartItem, customer *models.CustomerProfile, paymentInfo *models.PaymentInfo) bool {
	return paymentInfo != nil && paymentInfo.VoucherCode != nil && v.Code == *paymentInfo.VoucherCode
}

func (v VoucherDiscount) Apply(cartItems []*models.CartItem, customer *models.CustomerProfile, paymentInfo *models.PaymentInfo) (decimal.Decimal, string) {
	var totalPrice decimal.Decimal
	for _, item := range cartItems {
		totalPrice = totalPrice.Add(item.Product.CurrentPrice.Mul(decimal.NewFromInt(int64(item.Quantity))))
	}

	discount := totalPrice.Mul(v.Percent).Div(decimal.NewFromInt(int64(100)))
	return discount, fmt.Sprintf("Voucher Code: %v", v.Code)
}
