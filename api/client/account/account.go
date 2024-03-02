package account

type ClientAccount interface {
	Id() (int, error)
	IsAdmin() (bool, error)
	Username() (string, error)
	Email() (string, error)
	FirstName() (string, error)
	LastName() (string, error)
	Language() (string, error)
	ChangePassword(string, string) (bool, error)
	ListAPIKeys() ([]APIKey, error)
	CreateAPIKey(string, []string) (APIKey, error)
	DeleteAPIKey(string) (bool, error)
}

type APIKey struct {
	Identifier  string   `json:"identifier"`
	Description string   `json:"description"`
	AllowedIps  []string `json:"allowed_ips"`
	LastUsedAt  string   `json:"last_used_at"`
	CreatedAt   string   `json:"created_at"`
}
