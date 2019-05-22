package services

// TODO: Let's try mysql

// import (
// 	"github.com/mongodb/mongo-go-driver/bson"
//   "github.com/mongodb/mongo-go-driver/mongo"
//   "github.com/mongodb/mongo-go-driver/mongo/options"
// )
import (
	"any_suite/data"
)

// StorageManager stores data into `redis` or `mongo db`
type StorageManager struct{}

// NewStorageManager create a new instance of `StorageManager`
func NewStorageManager() (*StorageManager, error) {
	return &StorageManager{}, nil
}

// Run runs storage method
func (p *StorageManager) Run() error {
	db := data.NewAppDB()
	db.Conn()

	return nil
}
