package services

// TODO: Let's try mysql

// import (
// 	"github.com/mongodb/mongo-go-driver/bson"
//   "github.com/mongodb/mongo-go-driver/mongo"
//   "github.com/mongodb/mongo-go-driver/mongo/options"
// )

// StorageManager stores data into `redis` or `mongo db`
type StorageManager struct{}

// NewStorageManager create a new instance of `StorageManager`
func NewStorageManager() (*StorageManager, error) {
	return &StorageManager{}, nil
}

func (p *StorageManager) connect() error {
	// client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")

	// if err != nil {
	// 		log.Fatal(err)
	// }

	// // Check the connection
	// err = client.Ping(context.TODO(), nil)

	// if err != nil {
	// 		log.Fatal(err)
	// }

	// fmt.Println("Connected to MongoDB!")
	return nil
}
