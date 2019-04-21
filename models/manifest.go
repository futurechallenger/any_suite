package models

// Manifest stands for the app appmanifest.json
type Manifest struct {
	Triggers map[string]interface{} `json:"triggers"`
}
