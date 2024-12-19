package types

type ApiKeyScope string

// ApiKeyScope constants
const (
	ApiKeyScopeAdmin         ApiKeyScope = "admin"
	ApiKeyScopeCreateLicense ApiKeyScope = "create_license"
	ApiKeyScopeRevokeLicense ApiKeyScope = "revoke_license"
	ApiKeyScopeManageApiKeys ApiKeyScope = "manage_api_keys"
)
