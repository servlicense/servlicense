package types

type ApiKeyScope string

// ApiKeyScope constants
const (
	ApiKeyScopeAdmin         ApiKeyScope = "admin"
	ApiKeyScopeListLicenses  ApiKeyScope = "list_licenses"
	ApiKeyScopeCreateLicense ApiKeyScope = "create_license"
	ApiKeyScopeRevokeLicense ApiKeyScope = "revoke_license"
	ApiKeyScopeManageApiKeys ApiKeyScope = "manage_api_keys"
)

// checks if the given scopes contain the required scope
func ContainsScope(scopes []string, required ApiKeyScope) bool {
	for _, scope := range scopes {
		if scope == string(required) {
			return true
		}
	}
	return false
}
