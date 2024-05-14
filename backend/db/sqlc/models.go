// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type Gender string

const (
	GenderMale         Gender = "male"
	GenderFemale       Gender = "female"
	GenderNotSpecified Gender = "not-specified"
	GenderNotDisclosed Gender = "not-disclosed"
)

func (e *Gender) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Gender(s)
	case string:
		*e = Gender(s)
	default:
		return fmt.Errorf("unsupported scan type for Gender: %T", src)
	}
	return nil
}

type NullGender struct {
	Gender Gender `json:"gender"`
	Valid  bool   `json:"valid"` // Valid is true if Gender is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullGender) Scan(value interface{}) error {
	if value == nil {
		ns.Gender, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Gender.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullGender) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Gender), nil
}

type LanguageCode string

const (
	LanguageCodeChn LanguageCode = "chn"
	LanguageCodeJp  LanguageCode = "jp"
)

func (e *LanguageCode) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = LanguageCode(s)
	case string:
		*e = LanguageCode(s)
	default:
		return fmt.Errorf("unsupported scan type for LanguageCode: %T", src)
	}
	return nil
}

type NullLanguageCode struct {
	LanguageCode LanguageCode `json:"language_code"`
	Valid        bool         `json:"valid"` // Valid is true if LanguageCode is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullLanguageCode) Scan(value interface{}) error {
	if value == nil {
		ns.LanguageCode, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.LanguageCode.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullLanguageCode) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.LanguageCode), nil
}

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusShipped    OrderStatus = "shipped"
	OrderStatusDelivered  OrderStatus = "delivered"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

func (e *OrderStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = OrderStatus(s)
	case string:
		*e = OrderStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for OrderStatus: %T", src)
	}
	return nil
}

type NullOrderStatus struct {
	OrderStatus OrderStatus `json:"order_status"`
	Valid       bool        `json:"valid"` // Valid is true if OrderStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullOrderStatus) Scan(value interface{}) error {
	if value == nil {
		ns.OrderStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.OrderStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullOrderStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.OrderStatus), nil
}

type ProductStatus string

const (
	ProductStatusInStock      ProductStatus = "in-stock"
	ProductStatusOutOfStock   ProductStatus = "out-of-stock"
	ProductStatusDiscontinued ProductStatus = "discontinued"
)

func (e *ProductStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ProductStatus(s)
	case string:
		*e = ProductStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for ProductStatus: %T", src)
	}
	return nil
}

type NullProductStatus struct {
	ProductStatus ProductStatus `json:"product_status"`
	Valid         bool          `json:"valid"` // Valid is true if ProductStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullProductStatus) Scan(value interface{}) error {
	if value == nil {
		ns.ProductStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ProductStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullProductStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ProductStatus), nil
}

type Order struct {
	Pk              int64              `json:"pk"`
	ID              pgtype.UUID        `json:"id"`
	UserPk          int64              `json:"user_pk"`
	Status          OrderStatus        `json:"status"`
	IsPaid          bool               `json:"is_paid"`
	TotalPrice      int32              `json:"total_price"`
	ShippingAddress string             `json:"shipping_address"`
	ShippingDate    pgtype.Timestamptz `json:"shipping_date"`
	DeliveredDate   pgtype.Timestamptz `json:"delivered_date"`
	CreatedAt       pgtype.Timestamptz `json:"created_at"`
	UpdatedAt       pgtype.Timestamptz `json:"updated_at"`
}

type OrderDetail struct {
	Pk        int64              `json:"pk"`
	OrderPk   int64              `json:"order_pk"`
	ProductPk int64              `json:"product_pk"`
	Quantity  int32              `json:"quantity"`
	Price     int32              `json:"price"`
	Discount  int32              `json:"discount"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

type Product struct {
	Pk        int64              `json:"pk"`
	ID        pgtype.UUID        `json:"id"`
	Price     int32              `json:"price"`
	Discount  int32              `json:"discount"`
	ImageURLs []string           `json:"imageURLs"`
	Status    ProductStatus      `json:"status"`
	Quantity  int32              `json:"quantity"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

type ProductTranslation struct {
	Pk          int64              `json:"pk"`
	ProductPk   int64              `json:"product_pk"`
	Language    LanguageCode       `json:"language"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Category    string             `json:"category"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
}

type User struct {
	Pk        int64              `json:"pk"`
	ID        pgtype.UUID        `json:"id"`
	LineID    pgtype.Text        `json:"line_id"`
	BirthYear pgtype.Int4        `json:"birth_year"`
	Gender    Gender             `json:"gender"`
	Phone     pgtype.Text        `json:"phone"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
	FirstName string             `json:"first_name"`
	LastName  string             `json:"last_name"`
	Language  LanguageCode       `json:"language"`
	Address   string             `json:"address"`
	Longitude pgtype.Float8      `json:"longitude"`
	Latitude  pgtype.Float8      `json:"latitude"`
	LastLogin pgtype.Timestamptz `json:"last_login"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}
