package controllers

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"int_ecosys/config"
	"int_ecosys/models"
	"int_ecosys/services"
	"int_ecosys/utils"

	"github.com/labstack/echo/v4"

	requests "github.com/levigross/grequests"
)

// HomeController handle user's upload
type HomeController struct {
	tokenInfo models.AuthInfo
}

// ContainerHandler will handle user's upload
func (home *HomeController) ContainerHandler(c echo.Context) error {
	fmt.Println("Hello Uploader")

	container := &services.Container{}
	container.CheckInstalled()

	return c.String(http.StatusOK, "Hello World!")
}

// HelloHandler handle path `/hello`
func (home *HomeController) HelloHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "hello.html", map[string]interface{}{
		"name": "Jack",
	})
}

// HomeHandler handle path `/home`
func (home *HomeController) HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"Title":     "Home",
		"NeedLogin": true,
		"Data":      "Formal functions here",
	})
}

// AuthHandler handles auth
func (home *HomeController) AuthHandler(c echo.Context) error {
	server := config.Config()["server"]
	redirectURL := config.Config()["redirectUrl"]
	clientID := config.Config()["clientId"]

	authURI := fmt.Sprintf("%s/restapi/oauth/authorize", server)
	queryParams := map[string]string{
		"response_type": "code",
		"redirect_uri":  redirectURL,
		"client_id":     clientID,
		"display":       "popup",
	}

	queryString := utils.URLEncode(queryParams)
	authURI = fmt.Sprintf("%s?%s", authURI, queryString)

	retAuthRUI := authURI
	retRedirectURL := redirectURL
	retTokenJSON := "i dont know"
	if home.tokenInfo.AccessToken != "" {
		retAuthRUI = ""
		retRedirectURL = ""
		retTokenJSON = home.tokenInfo.AccessToken
	}
	retMap := map[string]string{
		"AuthUri":     retAuthRUI,
		"RedirectUri": retRedirectURL,
		"TokenJson":   retTokenJSON,
	}

	return c.Render(http.StatusOK, "auth.html", retMap)
}

// AuthCallbackHandler handles `/auth/callback`
func (home *HomeController) AuthCallbackHandler(c echo.Context) error {
	code := c.QueryParam("code")
	server := config.Config()["server"]
	redirectURL := config.Config()["redirectUrl"]

	authURI := fmt.Sprintf("%s/restapi/oauth/token", server)
	apiKey := fmt.Sprintf("%s:%s", config.Config()["clientId"], config.Config()["clientSecret"])
	apiKey = base64.StdEncoding.EncodeToString([]byte(apiKey))
	authHeader := fmt.Sprintf("Basic %s", apiKey)

	res, err := requests.Post(authURI, &requests.RequestOptions{
		Data: map[string]string{
			"grant_type":   "authorization_code",
			"code":         code,
			"redirect_uri": redirectURL,
		},
		Headers: map[string]string{
			"Authorization": authHeader,
			"Accept":        "application/json",
			"Content-Type":  "application/x-www-form-urlencoded",
		},
	})

	if err != nil || res.Ok != true {
		fmt.Printf("ERROR for request token %v, res: %s\n", err, res.String())
	}

	// TODO: Now we have code, let's request token
	resString := res.String()
	authInfo := &models.AuthInfo{}
	res.JSON(authInfo)
	home.tokenInfo = *authInfo

	fmt.Printf("res ===> %s\n", resString)

	return c.String(http.StatusOK, resString)

	// return c.Render(http.StatusOK, "home.html", map[string]string{
	// 	"Title":     "Home Page",
	// 	"NeedLogin": "No",
	// 	"Data":      res.String(),
	// })

	// return c.Render(http.StatusOK, "auth.html", map[string]string{
	// 	"AuthUri":     "",
	// 	"RedirectUri": "",
	// 	"TokenJson":   resString,
	// })
}
