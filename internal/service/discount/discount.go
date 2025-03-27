package discount

import (
	"github.com/shopspring/decimal"
	"unifizeAssignment/internal/models"
)

/*
* Interface for a type of discount, every discount rule must implement this
 */
type DiscountRule interface {

	// Checks if the current discount offer/rule is applicable
	IsApplicable(cartItems []*models.CartItem, customer *models.CustomerProfile, paymentInfo *models.PaymentInfo) bool

	// Calculates discount for the current rule/offer
	Apply(cartItems []*models.CartItem,
		customer *models.CustomerProfile,
		paymentInfo *models.PaymentInfo) (decimal.Decimal, string)
}
