# Deputy Tech Challenge
#### by Justin Davidson  
  
### Task
 
Given set of linked hierarchical _'roles'_, and a set of _'users'_, each with a specific role were provided,

- Set the roles in a memory store  
- Set the users in a memory store  
- For a given userID, return all the users that are subordinates of the given user 
  

### Solution
This task has been written in Go and for convenience has been made into a CLI app. 
Two different implementations have been completed. 
One backed by a Postgres database, and one backed by the graph database Neo4j.  
I assumed recursion was what was hoping to be seen in the solution hence the Postgres version, however I wanted to also include a Neo4j solution to show how simple it is to query a Neo4j database.  

### To run

git clone this repo. In a folder of your choice run
```
git clone https://github.com/RuNpiXelruN/deputy-challenge.git deputy-justin
```
..and then `cd` to the new folder from the command line
```
cd deputy-justin
```
If you have Go installed locally you can run `go mod tidy` which will download the dependencies required.  

Postgres and Neo4j are run in Docker containers. To start them run,
```
make up
```  
**It may take up to 30 seconds for the containers to be running.
```
go run cmd/main.go
```  
to run using the Go runtime.  

If you don't have Go installed, you can still run through the built binary provided within this repo **(which is a lot quicker)**  

Three binaries have been provided. One each for OSX, Linux, and Windows environments.  
To run the mac binary and view the CLI options,
```
./deputyJD
```
To run the windows binary,
```
./deputyJD_win
```
To run the linux binary,
```
./deputyJD_linux
```

To seed the databases, run (using the mac binary as an example below)
```
./deputyJD seed
```
To fetch the postgres subordinates (remember to run `./deputyJD` if you want to see the available commands)
```
./deputyJD pgGetSub --userID <someID>
```
To fetch the postgres subordinates
```
./deputyJD neoGetSub --userID <someID>
```

A `Makefile` has been provided for convenience.  
To see available make commands run,

```go
make help
```
Some other useful make commands are,
```go
make up          // starts docker containers for both neo4j and postgres
make down        // stops neo4j and postgres containers
make test        // runs unit tests
make build       // builds mac binary
make buildWin    // builds windows binary
make buildLinux  // builds linux binary
```

## Connecting to the databases
### Postgres
Once the container is running (`make up`), you can connect to the Postgres db with the following creds,
```
dbname: depchallenge 
user: postgres
pass: password
host: localhost (or 127.0.0.1)
port: 5433
```
** PLEASE NOTE THE PORT OF __5433__ and not the usual 5432 **  
### Neo4j  
Neo4j has a __browser dashboard__ for viewing your database (among others, it also has the `cypher-shell` tool for the command line).  
Once the container is running (`make up`), in your browser navigate to `localhost:7474`. If nothing loads initially the container may still be starting up. Once it loads you may be presented with a Neo4j login screen.  
Login with,
```
username: neo4j
password: test
```
From there you will land on your Neo4j dashboard. Have a play around with the data (once you run the `seed` command from the cli `./deputyJD seed`).  
You can execute cypher queries in the editor at the top. A couple of basic `cypher` queries for you to play with are, (`cmd + enter` to execute commands, `cmd + up arrow` to scroll through command history, `esc` to toggle editor full screen)
```go
match (n) return n  // returns all nodes and relationships. click on the graph tab to see the nodes interconnected and my architecture.
```
..and the query which returns subordinates. An example of userID=3 has been used here, but try others.
```go
MATCH (:User {id: 3})-[:HAS_ROLE_OF]->(:Role)<-[*]-(r)-[*]-(us :User)
UNWIND [us] AS u
WITH DISTINCT u
RETURN u AS users
```
```
match (n) detach delete n // deletes all nodes and relationships.
```
run `./deputyJD seed` to populate it once more.

## Thanks!
Thanks for this challenge guys! If there's anything else you need or wanted to ask do hesitate to ask :)  
  
Justin Davidson