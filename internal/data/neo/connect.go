package neo

import (
	"fmt"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

// Connect func
func Connect() bolt.Conn {
	driver := bolt.NewDriver()

	// NOTE: would usually be in env vars.
	neoConnString := fmt.Sprintf("bolt://%s:%s@%s:7687", "neo4j", "test", "localhost")

	conn, err := driver.OpenNeo(neoConnString)
	if err != nil {
		fmt.Printf("Error getting neo connection: %+v\n", err)
	}

	return conn
}
