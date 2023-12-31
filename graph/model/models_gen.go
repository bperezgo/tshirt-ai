// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type CustomizedProduct struct {
	ID      string   `json:"id"`
	Product *Product `json:"product"`
	Design  string   `json:"design"`
}

type Description struct {
	Short string `json:"short"`
	Long  string `json:"long"`
	HTML  string `json:"html"`
}

type NewCustomizedProduct struct {
	ProductID string   `json:"productId"`
	Text      string   `json:"text"`
	UserID    string   `json:"userId"`
	Images    []string `json:"images"`
}

type Product struct {
	ID          string       `json:"id"`
	Title       string       `json:"title"`
	Type        ProductType  `json:"type"`
	Size        ProductSize  `json:"size"`
	Images      []string     `json:"images"`
	Description *Description `json:"description"`
	Price       float64      `json:"price"`
	Stock       int          `json:"stock"`
	Reviews     []*Review    `json:"reviews"`
}

type Review struct {
	ID      string   `json:"id"`
	Product *Product `json:"product"`
	Body    string   `json:"body"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ProductSize string

const (
	ProductSizeXs ProductSize = "XS"
	ProductSizeS  ProductSize = "S"
	ProductSizeM  ProductSize = "M"
	ProductSizeL  ProductSize = "L"
	ProductSizeXl ProductSize = "XL"
)

var AllProductSize = []ProductSize{
	ProductSizeXs,
	ProductSizeS,
	ProductSizeM,
	ProductSizeL,
	ProductSizeXl,
}

func (e ProductSize) IsValid() bool {
	switch e {
	case ProductSizeXs, ProductSizeS, ProductSizeM, ProductSizeL, ProductSizeXl:
		return true
	}
	return false
}

func (e ProductSize) String() string {
	return string(e)
}

func (e *ProductSize) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProductSize(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProductSize", str)
	}
	return nil
}

func (e ProductSize) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ProductType string

const (
	ProductTypeTshirt ProductType = "TSHIRT"
	ProductTypeHoodie ProductType = "HOODIE"
)

var AllProductType = []ProductType{
	ProductTypeTshirt,
	ProductTypeHoodie,
}

func (e ProductType) IsValid() bool {
	switch e {
	case ProductTypeTshirt, ProductTypeHoodie:
		return true
	}
	return false
}

func (e ProductType) String() string {
	return string(e)
}

func (e *ProductType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProductType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProductType", str)
	}
	return nil
}

func (e ProductType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
