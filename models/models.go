package models

type License struct {
	License    string `json:"license"`
	Active     bool   `json:"active"`
	ValidUntil string `json:"valid_until"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
