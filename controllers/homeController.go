package controllers

import (
	"fmt"
	"net/http"

	"int_ecosys/config"
	"int_ecosys/services"

	"github.com/labstack/echo/v4"

	requests "github.com/levigross/grequests"
)

// HomeController handle user's upload
type HomeController struct {
}

// HomeHandler will handle user's upload
func (home *HomeController) HomeHandler(c echo.Context) error {
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

// AuthHandler handles auth
func (home *HomeController) AuthHandler(c echo.Context) error {
	server := config.Config()["server"]
	redirectURL := config.Config()["redirectUrl"]
	clientID := config.Config()["clientId"]

	codeServer := fmt.Sprintf("%s/restapi/oauth/authorize", server)
	codeRes, codeErr := requests.Get(codeServer, &requests.RequestOptions{
		Params: map[string]string{
			"response_type": "code",
			"redirect_uri":  redirectURL,
			"client_id":     clientID,
			"display":       "popup",
		}})

	if codeErr != nil {
		fmt.Printf("Unable to make such request:  %v\n", codeRes)
		return codeErr
	}
	if !codeRes.Ok {
		fmt.Printf("Request failed with: %d\n", codeRes.StatusCode)
	}

	fmt.Printf("RES HTML: %s\n", codeRes.String())
	// server = fmt.Sprintf("%s/restapi/oauth/token", server)
	// res, err := requests.Post(server, &requests.RequestOptions{
	// 	Data: map[string]string{
	// 		"grant_type": "authorization_code",
	// 	}})
	// if err != nil {
	// 	// log.Fatalln("Unable to make such request: ", err)
	// 	fmt.Printf("Unable to make such request:  %v\n", res)
	// 	return err
	// }

	// if !res.Ok {
	// 	fmt.Printf("Request failed with: %d\n", res.StatusCode)
	// }

	return c.HTML(http.StatusOK, codeRes.String())
}
