// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphModel

import (
	"time"
)

type AffiliateData struct {
	Items       []*AffiliateItem `json:"items,omitempty"`
	LowestPrice *int             `json:"lowestPrice,omitempty"`
}

type AffiliateItem struct {
	Name     string  `json:"name"`
	Price    *int    `json:"price,omitempty"`
	URL      string  `json:"URL"`
	ImageURL *string `json:"imageURL,omitempty"`
}

type AuthPayload struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}

type BoardInput struct {
	LiquorID string `json:"liquorID"`
	Text     string `json:"text"`
	Rate     *int   `json:"rate,omitempty"`
}

type BoardPost struct {
	ID           string    `json:"id"`
	UserID       *string   `json:"userId,omitempty"`
	UserName     *string   `json:"userName,omitempty"`
	CategoryID   int       `json:"categoryId"`
	CategoryName string    `json:"categoryName"`
	LiquorID     string    `json:"liquorId"`
	LiquorName   string    `json:"liquorName"`
	Text         string    `json:"text"`
	Rate         *int      `json:"rate,omitempty"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type BookMarkListUser struct {
	UserID    string    `json:"userId"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type Category struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Parent      *int        `json:"parent,omitempty"`
	Description *string     `json:"description,omitempty"`
	ImageURL    *string     `json:"imageUrl,omitempty"`
	ImageBase64 *string     `json:"imageBase64,omitempty"`
	VersionNo   *int        `json:"versionNo,omitempty"`
	UpdatedAt   *time.Time  `json:"updatedAt,omitempty"`
	Children    []*Category `json:"children,omitempty"`
}

type CategoryHistory struct {
	Now       *Category   `json:"now"`
	Histories []*Category `json:"histories,omitempty"`
}

type CategoryTrail struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Liquor struct {
	ID            string           `json:"id"`
	CategoryID    int              `json:"categoryId"`
	CategoryName  string           `json:"categoryName"`
	CategoryTrail []*CategoryTrail `json:"categoryTrail,omitempty"`
	Name          string           `json:"name"`
	Description   *string          `json:"description,omitempty"`
	ImageURL      *string          `json:"imageUrl,omitempty"`
	ImageBase64   *string          `json:"imageBase64,omitempty"`
	UpdatedAt     time.Time        `json:"updatedAt"`
	Rate5Users    []string         `json:"rate5Users"`
	Rate4Users    []string         `json:"rate4Users"`
	Rate3Users    []string         `json:"rate3Users"`
	Rate2Users    []string         `json:"rate2Users"`
	Rate1Users    []string         `json:"rate1Users"`
	VersionNo     int              `json:"versionNo"`
}

type LiquorHistory struct {
	Now       *Liquor   `json:"now"`
	Histories []*Liquor `json:"histories,omitempty"`
}

type ListFromCategory struct {
	CategoryName        string    `json:"categoryName"`
	CategoryDescription *string   `json:"categoryDescription,omitempty"`
	Liquors             []*Liquor `json:"liquors"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Mutation struct {
}

type Query struct {
}

type RegisterInput struct {
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Password    *string `json:"password,omitempty"`
	Profile     *string `json:"profile,omitempty"`
	ImageBase64 *string `json:"imageBase64,omitempty"`
}

type User struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Profile     *string `json:"profile,omitempty"`
	ImageBase64 *string `json:"imageBase64,omitempty"`
}

type UserEvaluateList struct {
	RecentComments []*UserLiquor `json:"recentComments,omitempty"`
	Rate5Liquors   []*UserLiquor `json:"rate5Liquors,omitempty"`
	Rate4Liquors   []*UserLiquor `json:"rate4Liquors,omitempty"`
	Rate3Liquors   []*UserLiquor `json:"rate3Liquors,omitempty"`
	Rate2Liquors   []*UserLiquor `json:"rate2Liquors,omitempty"`
	Rate1Liquors   []*UserLiquor `json:"rate1Liquors,omitempty"`
	NoRateLiquors  []*UserLiquor `json:"noRateLiquors,omitempty"`
}

type UserLiquor struct {
	ID           string    `json:"id"`
	LiquorID     string    `json:"liquorId"`
	Name         string    `json:"name"`
	CategoryID   int       `json:"categoryId"`
	CategoryName string    `json:"categoryName"`
	ImageBase64  *string   `json:"imageBase64,omitempty"`
	Comment      *string   `json:"comment,omitempty"`
	Rate         *int      `json:"rate,omitempty"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type UserPageData struct {
	EvaluateList *UserEvaluateList `json:"evaluateList"`
	User         *User             `json:"user"`
}
