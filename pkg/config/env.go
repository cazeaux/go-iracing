package config

var IR_API_URL string
var IR_PATH_LEGACY_AUTH = "/auth"
var IR_OAUTH_URL = "https://oauth.iracing.com/oauth2/token"

func init() {
	IR_API_URL = "https://members-ng.iracing.com"
}