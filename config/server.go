// TODO: access clientId & client secret from env
// https://gobyexample.com/environment-variables
// ENV=DEBUG go run app.go

package config

import (
	"any_suite/utils"
	"os"
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
	clientID, _ := utils.Condition(isDebug(env), os.Getenv("INTECO_DEV_CLIENT_ID"), "").(string)
	clientSecret, _ := utils.Condition(isDebug(env), os.Getenv("INTECO_DEV_CLIENT_SECRET"), "").(string)
	return map[string]string{
		"server":       server,
		"redirectUrl":  redirectURL,
		"clientId":     clientID,
		"clientSecret": clientSecret,
	}
}
