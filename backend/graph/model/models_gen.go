// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type Category struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Parent   *int        `json:"parent,omitempty"`
	Children []*Category `json:"children,omitempty"`
}

type CreateLiquorRequest struct {
	Name        string          `json:"name"`
	CategoryID  int             `json:"category_id"`
	Description *string         `json:"description,omitempty"`
	Image       *graphql.Upload `json:"image,omitempty"`
}

type Liquor struct {
	ID          string    `json:"id"`
	CategoryID  int       `json:"category_id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	ImageURL    *string   `json:"imageUrl,omitempty"`
	ImageBase64 *string   `json:"imageBase64,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Mutation struct {
}

type Query struct {
}
