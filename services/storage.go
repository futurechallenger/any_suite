package services

// StorageManager stores data into `redis` or `mongo db`
type StorageManager struct{}

// NewStorageManager create a new instance of `StorageManager`
func NewStorageManager() (*StorageManager, error) {
	return &StorageManager{}, nil
}
