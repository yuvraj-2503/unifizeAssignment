package service

import (
	"context"
	"fmt"
	"github.com/shopspring/decimal"
	"reflect"
	"testing"
	"unifizeAssignment/internal/models"
	"unifizeAssignment/internal/service/discount"
	"unifizeAssignment/internal/test"
)

var dummyDiscountRules []discount.DiscountRule
var dummyCart []*models.CartItem
var dummyPaymentInfo *models.PaymentInfo
var ctx context.Context

func init() {
	dummyDiscountRules = test.GetDummyDiscount()
	dummyCart, dummyPaymentInfo = test.GetDummyCart()
	ctx = context.Background()
}

func TestDiscountServiceImpl_CalculateCartDiscounts(t *testing.T) {
	type fields struct {
		discountRules []discount.DiscountRule
	}
	type args struct {
		ctx         context.Context
		cartItems   []*models.CartItem
		customer    *models.CustomerProfile
		paymentInfo *models.PaymentInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.DiscountedPrice
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Success Test",
			fields: fields{
				discountRules: dummyDiscountRules,
			},
			args: args{
				ctx:       ctx,
				cartItems: dummyCart,
				customer: &models.CustomerProfile{
					ID:   "C001",
					Tier: "PREMIUM",
				},
				paymentInfo: dummyPaymentInfo,
			},
			want: &models.DiscountedPrice{
				OriginalPrice: decimal.NewFromInt(int64(1000)),
				FinalPrice:    decimal.NewFromInt(int64(380)),
				AppliedDiscounts: map[string]decimal.Decimal{
					"Bank Offer: ICICI":        decimal.NewFromInt(50),
					"Brand Offer: PUMA":        decimal.NewFromInt(400),
					"Category Offer: T-shirts": decimal.NewFromInt(50),
					"Voucher Code: SUPER12":    decimal.NewFromInt(120),
				},
				Message: "Discounts applied successfully",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DiscountServiceImpl{
				discountRules: tt.fields.discountRules,
			}
			got, err := d.CalculateCartDiscounts(tt.args.ctx, tt.args.cartItems, tt.args.customer, tt.args.paymentInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateCartDiscounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				fmt.Println("Discounted Price:  " + got.FinalPrice.String())
				t.Errorf("CalculateCartDiscounts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiscountServiceImpl_ValidateDiscountCode(t *testing.T) {
	type fields struct {
		discountRules []discount.DiscountRule
	}
	type args struct {
		ctx       context.Context
		code      string
		cartItems []*models.CartItem
		customer  *models.CustomerProfile
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Success Test",
			fields: fields{
				discountRules: dummyDiscountRules,
			},
			args: args{
				ctx:       ctx,
				code:      "SUPER12",
				cartItems: dummyCart,
				customer: &models.CustomerProfile{
					ID:   "C001",
					Tier: "PREMIUM",
				},
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DiscountServiceImpl{
				discountRules: tt.fields.discountRules,
			}
			got, err := d.ValidateDiscountCode(tt.args.ctx, tt.args.code, tt.args.cartItems, tt.args.customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateDiscountCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidateDiscountCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDiscountServiceImpl(t *testing.T) {
	type args struct {
		discountRules []discount.DiscountRule
	}
	tests := []struct {
		name string
		args args
		want DiscountServiceImpl
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDiscountServiceImpl(tt.args.discountRules); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDiscountServiceImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}
