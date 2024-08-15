package bootstrap

import (
	"context"
	"path/filepath"

	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	Mongo *mongo.Client
	Env   *Env
}

func App() Application {
	projectRoot, _ := filepath.Abs(filepath.Join(".."))

	app := Application{}
	app.Env = NewEnv(projectRoot)

	app.Mongo = NewMongoDatabase(app.Env)

	return app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo, context.TODO())
}
