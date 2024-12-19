package models

type License struct {
	License    string `json:"license"`
	Active     bool   `json:"active"`
	ValidUntil string `json:"valid_until"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type Apikey struct {
	Id        int      `json:"id"`
	ApiKey    string   `json:"api_key"`
	Name      string   `json:"name"`
	Scopes    []string `json:"scopes"`
	CreatedAt string   `json:"created_at"`
}
