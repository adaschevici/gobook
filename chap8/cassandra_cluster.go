package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

func main() {

	cass := gocql.NewCluster("127.0.0.1")
	cass.ProtoVersion = 3
	cass.Keyspace = "filemaster"
	cass.Consistency = gocql.LocalQuorum

	session, _ := cass.CreateSession()
	defer session.Close()

	var fileTime int

	if err := session.Query(`SELECT file_modified_time FROM filemaster WHERE filename = ? LIMIT 1 ALLOW FILTERING`, "test.txt").Consistency(gocql.One).Scan(&fileTime); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Last modified", fileTime)
}
