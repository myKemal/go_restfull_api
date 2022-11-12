package app

import (
	"github.com/myKemal/go_restfull_api/application/common"

	"github.com/myKemal/go_restfull_api/application/server"

	db "github.com/myKemal/go_restfull_api/application/dbClient"
	handler "github.com/myKemal/go_restfull_api/application/handlers"
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

}

// Run Runs the application.
func (a *App) Run() {
	common.Logger.Infof("Application running at port : 8080 ")
	err := a.applicationServer.Run("8080")
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
