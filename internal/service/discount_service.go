package service

import (
	"context"
	"fmt"
	"github.com/shopspring/decimal"
	"slices"
	"unifizeAssignment/internal/common"
	"unifizeAssignment/internal/models"
	"unifizeAssignment/internal/service/discount"
)

type DiscountServiceImpl struct {
	discountRules []discount.DiscountRule
}

func NewDiscountServiceImpl(discountRules []discount.DiscountRule) DiscountServiceImpl {
	return DiscountServiceImpl{discountRules: discountRules}
}

func (d DiscountServiceImpl) CalculateCartDiscounts(ctx context.Context, cartItems []*models.CartItem, customer *models.CustomerProfile,
	paymentInfo *models.PaymentInfo) (*models.DiscountedPrice, error) {
	if len(cartItems) == 0 {
		return nil, &common.NotFoundError{Message: "cart is empty"}
	}

	var originalPrice, finalPrice, totalDiscount decimal.Decimal
	appliedDiscounts := make(map[string]decimal.Decimal)

	for _, item := range cartItems {
		originalPrice = originalPrice.Add(item.Product.CurrentPrice.Mul(decimal.NewFromInt(int64(item.Quantity))))
	}

	finalPrice = originalPrice

	for _, rule := range d.discountRules {
		if rule.IsApplicable(cartItems, customer, paymentInfo) {
			discount, message := rule.Apply(cartItems, customer, paymentInfo)
			if !discount.IsZero() {
				appliedDiscounts[message] = discount
				totalDiscount = totalDiscount.Add(discount)
			}
		}
	}

	finalPrice = finalPrice.Sub(totalDiscount)

	return &models.DiscountedPrice{
		OriginalPrice:    originalPrice,
		FinalPrice:       finalPrice,
		AppliedDiscounts: appliedDiscounts,
		Message:          "Discounts applied successfully",
	}, nil

}

func (d DiscountServiceImpl) ValidateDiscountCode(ctx context.Context, code string,
	cartItems []*models.CartItem,
	customer *models.CustomerProfile) (bool, error) {
	if code == "" {
		return false, &common.ValidationError{Message: "code is empty"}
	}

	for _, rule := range d.discountRules {
		if voucher, ok := rule.(discount.VoucherDiscount); ok && voucher.Code == code {
			if voucher.BrandExclusions != nil {
				for _, item := range cartItems {
					if slices.Contains(voucher.BrandExclusions, item.Product.Brand) {
						return false, &common.ValidationError{Message: fmt.Sprintf("discount code '%s' is not applicable for brand: %s", code, item.Product.Brand)}
					}
				}
			}

			if voucher.CategoryExclusions != nil {
				for _, item := range cartItems {
					if slices.Contains(voucher.CategoryExclusions, item.Product.Category) {
						return false, &common.ValidationError{Message: fmt.Sprintf("discount code '%s' is not applicable for caategory: %s", code, item.Product.Category)}
					}
				}
			}

			if voucher.CustomerTier != "" && voucher.CustomerTier != customer.Tier {
				return false, &common.ValidationError{Message: fmt.Sprintf("discount code '%s' requires customer tier: %s", code, voucher.CustomerTier)}
			}

			return true, nil
		}
	}

	return false, &common.ValidationError{Message: fmt.Sprintf("discount code '%s' is not valid", code)}
}
