package config

import (
	"int_ecosys/utils"
)

func isDebug(env string) func() bool {
	return func() bool {
		return env == "debug"
	}
}

// Config returns configurations
func Config() map[string]string {
	env := BuildEnv()
	server, _ := utils.Condition(isDebug(env),
		"https://platform.devtest.ringcentral.com",
		"https://platform.ringcentral.com").(string)
	redirectURL, _ := utils.Condition(isDebug(env),
		"http://localhost:1323/auth/callback",
		"").(string)
	clientID, _ := utils.Condition(isDebug(env), "ohnIbGTJTt-0CWB48kbNjQ", "").(string)
	return map[string]string{
		"server":      server,
		"redirectUrl": redirectURL,
		"clientId":    clientID}
}
