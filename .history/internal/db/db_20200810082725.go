package db

// Service type 
type Service interface {}

// Client type
type Client struct {
	pg PGService
	neo NeoService
}

func NewClient(pg PGService, neo NeoService) *Client {
	return &Client{
		pg: PGService,
		neo: NeoService,
	}
}