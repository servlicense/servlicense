package types

type ApiKeyScope string

// checks if the current scope is contained in scopes
func (k ApiKeyScope) InScopes(scopes []string) bool {
	asStr := string(k)
	for _, scope := range scopes {
		if scope == asStr {
			return true
		}
	}
	return false

}

// ApiKeyScope constants
const (
	ApiKeyScopeAdmin         ApiKeyScope = "admin"
	ApiKeyScopeListLicenses  ApiKeyScope = "list_licenses"
	ApiKeyScopeCreateLicense ApiKeyScope = "create_license"
	ApiKeyScopeRevokeLicense ApiKeyScope = "revoke_license"
	ApiKeyScopeManageApiKeys ApiKeyScope = "manage_api_keys"
)
