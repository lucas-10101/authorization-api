package utils

type RequestContextKey string

const (
	HeaderAuthorization                    string = "Authorization"
	HeaderContentType                      string = "Content-Type"
	HeaderAccept                           string = "Accept"
	HeaderCacheControl                     string = "Cache-Control"
	HeaderUserAgent                        string = "User-Agent"
	HeaderXRequestedWith                   string = "X-Requested-With"
	HeaderOrigin                           string = "Origin"
	HeaderAccessControlAllowOrigin         string = "Access-Control-Allow-Origin"
	HeaderAccessControlAllowMethods        string = "Access-Control-Allow-Methods"
	HeaderAccessControlAllowHeaders        string = "Access-Control-Allow-Headers"
	HeaderAccessControlExposeHeaders       string = "Access-Control-Expose-Headers"
	HeaderAccessControlAllowCredentials    string = "Access-Control-Allow-Credentials"
	HeaderAccessControlMaxAge              string = "Access-Control-Max-Age"
	HeaderAccessControlRequestMethod       string = "Access-Control-Request-Method"
	HeaderAccessControlRequestHeaders      string = "Access-Control-Request-Headers"
	HeaderAccessControlRequestCredentials  string = "Access-Control-Request-Credentials"
	HeaderAccessControlAllowPrivateNetwork string = "Access-Control-Allow-Private-Network"
	HeaderAccessControlAllowPublicKeyPins  string = "Access-Control-Allow-Public-Key-Pins"
	HeaderXApiKey                          string = "X-API-Key"
)

const (
	RequestContextPrincipalKey RequestContextKey = "principal"
	RequestContextUserIdKey    RequestContextKey = "userid"
	RequestContextTenantIdKey  RequestContextKey = "tenantid"
)
