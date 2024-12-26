package models

type Application struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type License struct {
	License    string `json:"license"`
	AppID      int    `json:"app_id"`
	Active     bool   `json:"active"`
	ValidUntil string `json:"valid_until"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type ApiKey struct {
	Id        string   `json:"id"`
	ApiKey    string   `json:"api_key"`
	Name      string   `json:"name"`
	Scopes    []string `json:"scopes"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}
