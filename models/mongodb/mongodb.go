package mongodb
import (
	"context"
	"fiber/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// MongoInstance contains the Mongo client and database objects
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mongodb MongoInstance

// Source: https://github.com/mongodb/mongo-go-driver
func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(helpers.LoadConfiguration("mongo_db_uri").(string)))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(helpers.LoadConfiguration("mongo_db").(string))

	if err != nil {
		panic(err)
	}

	mongodb = MongoInstance{
		Client: client,
		Db:     db,
	}
}
