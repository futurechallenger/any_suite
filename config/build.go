package config

import "fmt"

// BuildEnv current running environment, default is `debug`
var buildEnv string

// SetBuildEnv set the build env variable
func SetBuildEnv(env *string) string {
	if env == nil {
		buildEnv = "debug"
	}
	buildEnv = *env

	fmt.Printf("===>Build Env: %s\n", buildEnv)

	return buildEnv
}

// BuildEnv returns `build env`, default is `debug`, it can also be `dev`
func BuildEnv() string {
	return buildEnv
}
