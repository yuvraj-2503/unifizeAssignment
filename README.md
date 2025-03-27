# Unifize Backend Developer Assignment

## Overview

This project implements a **Discount Service** for an e-commerce platform, handling various types of discounts such as:

- **Brand-specific discounts** (e.g., "Min 40% off on PUMA")
- **Category-specific deals** (e.g., "Extra 10% off on T-shirts")
- **Bank card offers** (e.g., "10% instant discount on ICICI Bank cards")
- **Vouchers** (e.g., "SUPER69" for 69% off on any product)

## Features

- **Accurate discount calculations** ensuring proper stacking order
- **Extensible discount rules** using interfaces
- **Validation of discount codes** for eligibility
- **Detailed error messages** for better debugging
- **Well-structured Go modules** following best practices

## Tech Stack

- **Language**: Golang
- **Package Management**: Go Modules
- **Decimal Handling**: `shopspring/decimal`
- **Server Framework**: `gin`
- **Testing**: `testing` package (with table-driven tests)

## Project Structure

```
.
├── internal/
│   ├── common/
│   │   └── errors.go                    # custom errors structs
│   │   └── response.go                  # custom response structs and handlers 
│   ├── handlers/
│   │   └── handler.go                   # request handlers
│   │   └── load_handlers.go             # load_handlers / initialiser
│   ├── models/
│   │   └── models.go                    # Data models for products, cart, discounts
│   ├── service/
│   │   └── discount/
│   │   │   └── bank_discount.go         # Bank Offer Discount Logic
│   │   │   └── brand_discount.go        # Brand Offer Discount Logic
│   │   │   └── category_discount.go     # Category Offer Discount Logic
│   │   │   └── discount.go              # Discount Rule Interface
│   │   │   └── voucher_discount.go      # Voucher Offer Discount Logic
│   │   ├── api.go                       # Discount Service
│   │   ├── discount_service.go          # Discount Service Implementation
│   │   └── discount_service_test.go     # Unit tests
├── testdata/
│   └── fake_data.go                     # Sample test data
├── go.mod                               # Dependency management
├── go.sum                               # Package checksums
├── main.go                              # Entry point for the application
└── README.md                            # Documentation
```

## Installation & Setup

1. **Clone the repository:**
   ```sh
   git clone https://github.com/yuvraj-2503/unifizeAssignment
   cd unifizeAssignment
   ```
2. **Install dependencies:**
   ```sh
   go mod tidy
   ```
3. **Run Tests:**
   ```sh
   go test ./internal/service
   ```

## Usage

### Running the Service

```sh
go run main.go
```

## API Endpoints

### Calculate Discount
- `POST /api/v1/unifize/calculate/discount` - Calculate the final price after applying discounts.

### Sample JSON Request

```json
{
  "cartItems": [
    {
      "product": {
        "id": "P123",
        "brand": "PUMA",
        "brand_tier": "premium",
        "category": "T-shirt",
        "base_price": "1000.00",
        "current_price": "600.00"
      },
      "quantity": 2,
      "size": "M"
    },
    {
      "product": {
        "id": "P456",
        "brand": "Nike",
        "brand_tier": "regular",
        "category": "Shoes",
        "base_price": "5000.00",
        "current_price": "4000.00"
      },
      "quantity": 1,
      "size": "10"
    }
  ],
  "customerProfile": {
    "id": "C789",
    "tier": "PREMIUM"
  },
  "paymentInfo": {
    "method": "CARD",
    "bank_name": "ICICI",
    "card_type": "CREDIT",
    "voucher_code": "SUPER69"
  }
}
```

### Sample JSON Response

```json
{
  "original_price": "5200",
  "final_price": "4460",
  "applied_discounts": {
    "Bank Offer: ICICI": "260",
    "Brand Offer: PUMA": "480"
  },
  "message": "Discounts applied successfully"
}
```

### Validate Coupon
- `POST /api/v1/unifize/validate/coupon/:couponCode` - Validate if a discount coupon can be applied.

### Sample JSON Request

```json
{
  "cartItems": [
    {
      "product": {
        "id": "P123",
        "brand": "PUMA",
        "brand_tier": "premium",
        "category": "T-shirt",
        "base_price": "1000.00",
        "current_price": "600.00"
      },
      "quantity": 2,
      "size": "M"
    },
    {
      "product": {
        "id": "P456",
        "brand": "Nike",
        "brand_tier": "regular",
        "category": "Shoes",
        "base_price": "5000.00",
        "current_price": "4000.00"
      },
      "quantity": 1,
      "size": "10"
    }
  ],
  "customerProfile": {
    "id": "C789",
    "tier": "PREMIUM"
  }
}
```

### Sample JSON Response

```json
{
  "result": true
}
```

## Assumptions & Decisions

- Discounts are applied in the following order:
    1. **Brand & Category Discounts**
    2. **Voucher Discounts**
    3. **Bank Offers**
- A voucher discount is only valid if the **brand/category/customer tier** conditions are met.
- The system ensures **extensibility** by allowing new discount rules to be added dynamically.

## Contribution

Feel free to fork the repository and raise pull requests for improvements.

