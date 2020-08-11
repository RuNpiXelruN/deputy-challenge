package db

// Client type
type Client struct {
	pg PGService
	neo NeoService
}

// NewClient func 
func NewClient(pg PGService, neo NeoService) *Client {
	return &Client{
		pg: pg,
		neo: neo,
	}
}

// SetRoles func 
func (c *Client) SetRoles() error {
	return nil
}