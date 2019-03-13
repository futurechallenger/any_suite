package controllers

import (
	"fmt"
	"net/http"

	"int_ecosys/config"
	"int_ecosys/services"

	"github.com/labstack/echo/v4"
)

// HomeController handle user's upload
type HomeController struct {
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
	// codeRes, codeErr := requests.Get(codeServer, &requests.RequestOptions{
	// 	Params: map[string]string{
	// 		"response_type": "code",
	// 		"redirect_uri":  redirectURL,
	// 		"client_id":     clientID,
	// 		"display":       "popup",
	// 	}},
	// )

	queryParams := map[string]string{
		"response_type": "code",
		"redirect_uri":  redirectURL,
		"client_id":     clientID,
		"display":       "popup",
	}

	queryString := ""
	count := 0
	for key, val := range queryParams {
		if count == 0 {
			queryString = fmt.Sprintf("%s=%s", key, val)
		} else {
			queryString = fmt.Sprintf("%s&%s=%s", queryString, key, val)
		}

		count = count + 1
	}
	authURI = fmt.Sprintf("%s?%s", authURI, queryString)

	return c.Render(http.StatusOK, "auth.html", map[string]string{
		"AuthUri":     authURI,
		"RedirectUri": redirectURL,
		"TokenJson":   "i dont know yet",
	})
}

// AuthCallbackHandler handles `/auth/callback`
func (home *HomeController) AuthCallbackHandler(c echo.Context) error {
	code := c.QueryParam("code")
	fmt.Printf("===>Code is %s\n", code)

	// TODO: Now we have code, let's request token
	return c.String(http.StatusOK, code)
}
