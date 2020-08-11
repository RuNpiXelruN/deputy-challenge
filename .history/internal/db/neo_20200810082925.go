package db

// NeoService type 
type NeoService interface{}

// NeoClient type 
type NeoClient struct {}

// NewNeoClient func 
func NewNeoClient() *NeoClient {
	return &NeoClient{}
}