package models

import (
	"time"

	"github.com/google/uuid"
)

type PriceGroupSettings struct {
	ID           *int       `json:"id" db:"id" validate:"omitempty,int"`
	Name         *string    `json:"name,omitempty" db:"name" validate:"required"`
	Description  *string    `json:"description" db:"description" validate:"required,gte=10"`
	CurrencyType *string    `json:"currency_type" db:"currency_type" validate:"required"`
	CurrencyList *string    `json:"currency_list,omitempty" db:"currency_list" validate:"omitempty"`
	Active       *bool      `json:"active,omitempty" db:"active" validate:"omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	CreatedBy    *uuid.UUID `json:"created_by,omitempty" db:"created_by"`
	UpdatedBy    *uuid.UUID `json:"updated_by,omitempty" db:"updated_by"`
}
