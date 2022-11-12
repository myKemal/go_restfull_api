package app

import (
	"github.com/myKemal/go_restfull_api/application/common"
	"github.com/myKemal/go_restfull_api/application/config"

	"github.com/myKemal/go_restfull_api/application/server"

	db "github.com/myKemal/go_restfull_api/application/dbClient"
	handler "github.com/myKemal/go_restfull_api/application/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
)

type App struct {
	applicationServer *server.ApplicationServer
	mongoClient       db.MongoClient
	inMemoryClient    db.InMemoryClient
	mongoHandler      handler.MongoHandler
	inMemoryHandler   handler.InMemoryHandler
}

// initDependencies Initializes necessary dependencies of the application.
func (a *App) initDependencies() {
	a.applicationServer = server.NewApplicationServer()
	a.mongoClient = db.NewMongoClient()
	a.inMemoryClient = db.NewInMemoryClient()
	a.mongoHandler = handler.NewMongoHandler(a.mongoClient)
	a.inMemoryHandler = handler.NewInMemoryHandler(a.inMemoryClient)
}

// initDependencies Initializes route endpoints of the application.
func (a *App) initRoutes() {
	a.applicationServer.HandleFunctions("/api/v1/mongo", server.HandlerFunctions{
		Post: a.mongoHandler.GetRecords,
	})
	a.applicationServer.HandleFunctions("/api/v1/in-memory", server.HandlerFunctions{
		Get:  a.inMemoryHandler.GetRecords,
		Post: a.inMemoryHandler.Create,
	})

	a.applicationServer.HandleFunc("/docs/", httpSwagger.Handler(
		httpSwagger.URL("https://go-restfull-api.herokuapp.com/static/swagger.json"),
	))

}

// Run Runs the application.
func (a *App) Run() {

	port := config.GetPort()
	common.Logger.Infof("Application running at port:%s", port)
	err := a.applicationServer.Run(port)
	if err != nil {
		panic(err)
	}
}

// Stop Stops the application.
func (a *App) Stop() {
	a.mongoClient.Disconnect()
}

func NewApp() *App {
	app := &App{}
	app.initDependencies()
	app.initRoutes()
	return app
}
