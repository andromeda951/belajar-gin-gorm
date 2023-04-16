package book

import "encoding/json"

type BookRequest struct {
	Title string      `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,number"` // json.Number can take number in string json type
}
