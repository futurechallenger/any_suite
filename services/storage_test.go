package services

func TestMongoConnection(t *testing.T) {
	s, err := NewStorageManager();
	if err != nil {
		t.Errorf("Connect to mongodb error: %v", err)
	}
}