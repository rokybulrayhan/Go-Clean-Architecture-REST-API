package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PriceGroupSettings struct {
	bun.BaseModel `bun:"public.price_group_setting"`
	ID            int        `json:"id" db:"id" validate:"omitempty,int" bun:"type:id,pk"`
	Name          *string    `json:"name,omitempty" db:"name" validate:"required"`
	Description   *string    `json:"description" db:"description" validate:"required,gte=10"`
	CurrencyType  *string    `json:"currency_type" db:"currency_type" validate:"required"`
	CurrencyList  *string    `json:"currency_list,omitempty" db:"currency_list" validate:"omitempty"`
	IsActive      *bool      `json:"is_active,omitempty" db:"is_active" validate:"omitempty"`
	CreatedAt     *string    `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt     *string    `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt     *string    `json:"deleted_at,omitempty" db:"deleted_at"`
	CreatedBy     *uuid.UUID `json:"created_by,omitempty" db:"created_by"`
	UpdatedBy     *uuid.UUID `json:"updated_by,omitempty" db:"updated_by"`
}

type PriceGroupSettingsList struct {
	TotalCount int                   `json:"total_count"`
	TotalPages int                   `json:"total_pages"`
	Page       int                   `json:"page"`
	Size       int                   `json:"size"`
	HasMore    bool                  `json:"has_more"`
	PriceGroup []*PriceGroupSettings `json:"price_group"`
}
