// @title Go Restfull api
// @version 1.0
// @description This is a case.
// @termsOfService http://swagger.io/terms/

// @host go-restfull-api.herokuapp.com
// @BasePath /api/v1

package main

import app "github.com/myKemal/go_restfull_api/application"

func main() {
	application := app.NewApp()
	defer application.Stop()
	application.Run()
}
