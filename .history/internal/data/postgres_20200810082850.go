package db

// PGService type 
type PGService interface{}

// PGClient type 
type PGClient struct {}

// NewPGClient func 
func NewPGClient() *PGClient {
	return &PGClient{}
}