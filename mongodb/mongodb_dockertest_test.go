package mongodb

import (
	"context"
	"errors"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	go_dbconnectify "github.com/SyaibanAhmadRamadhan/go-dbconnectify"
)

func TestMongoDockerTest(t *testing.T) {
	// dockerTest := ginfra.InitDockerTest()
	// defer dockerTest.CleanUp()

	// mongoDockerTestConf := MongoDockerTestConf{}

	ctx := context.Background()
	opts := options.Client().ApplyURI("mongodb://localhost:20001/?replicaSet=dbrs&directConnection=true")
	mongoClient, err := OpenConnMongoClient(ctx, opts)
	if err != nil {
		panic(err)
	}
	// dockerTest.NewContainer(mongoDockerTestConf.ImageVersion(dockerTest, ""), func(res *dockertest.Resource) error {
	// 	conn, err := mongoDockerTestConf.ConnectClient(res)
	// 	gcommon.PanicIfError(err)
	// 	mongoClient = conn
	//
	// 	return nil
	// })

	mongodbConn := mongoClient.Database("exampleDatabase")
	err = mongodbConn.CreateCollection(context.Background(), "exampleCollection")
	if err != nil {
		t.Errorf("Failed to create collection: %v", err)
	}
	collection := mongodbConn.Collection("exampleCollection")

	tx := NewTxMongodb(mongoClient)
	err = tx.DoTransaction(context.Background(), &go_dbconnectify.TxOption{
		Type:   go_dbconnectify.TxTypeMongoDB,
		Option: &options.SessionOptions{},
	}, func(c context.Context) (bool, error) {
		document := bson.M{"name": "example", "value": "entry"}
		insertResult, err := collection.InsertOne(c, document)
		if err != nil {
			t.Errorf("Failed to insert document: %v", err)
		}

		if insertResult.InsertedID == nil {
			t.Error("No ID found for inserted document")
		}
		filter := bson.M{"name": "example"}
		var result bson.M
		err = collection.FindOne(c, filter).Decode(&result)
		if err != nil {
			t.Errorf("Failed to find document: %v", err)
		}

		if result["name"] != "example" || result["value"] != "entry" {
			t.Errorf("Document does not match expected values: %v", result)
		}

		return false, errors.New("asd")
	})
	if err != nil {
		panic(err)
	}

	t.Run("Sample Integration Test for MongoDB", func(t *testing.T) {
		document := bson.M{"name": "example", "value": "entry"}
		insertResult, err := collection.InsertOne(context.Background(), document)
		if err != nil {
			t.Errorf("Failed to insert document: %v", err)
		}

		if insertResult.InsertedID == nil {
			t.Error("No ID found for inserted document")
		}

		// Retrieve and validate
		filter := bson.M{"name": "example"}
		var result bson.M
		err = collection.FindOne(context.Background(), filter).Decode(&result)
		if err != nil {
			t.Errorf("Failed to find document: %v", err)
		}

		if result["name"] != "example" || result["value"] != "entry" {
			t.Errorf("Document does not match expected values: %v", result)
		}
	})
}

func setReplica(client *mongo.Client) {
	// replicaCfg := map[string]any{
	// 	"_id": "rs0",
	// 	"members": []any{
	// 		map[string]any{
	// 			"_id":  0,
	// 			"host": "localhost:27017",
	// 		},
	// 	},
	// 	"settings": map[string]any{
	// 		"chainingAllowed": true,
	// 	},
	// 	"heartbeatTimeoutSecs": 10,
	// }
	// adminDB := client.Database("admin")
	// cmd := bson.D{{"replSetInitiate", replicaCfg}}
	// result := bson.M{}
	// err := adminDB.RunCommand(context.Background(), cmd).Decode(&result)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// time.Sleep(10 * time.Second)
}

// query := bson.D{
// 	{
// 		Key: "$and", Value: bson.A{
// 		bson.D{{Key: "field1", Value: "value1"}},
// 		bson.D{{Key: "field2", Value: "value2"}},
// 	},
// 	},
// 	{
// 		Key: "$or", Value: bson.A{
// 		bson.D{{Key: "field3", Value: "value3"}},
// 		bson.D{{Key: "field4", Value: "value4"}},
// 	},
// 	},
// }
