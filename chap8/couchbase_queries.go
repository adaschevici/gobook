package main

import (
	"fmt"

	"github.com/couchbase/gocb"
)

var bucket *gocb.Bucket

func main() {
	// Connect to Cluster
	cluster, err := gocb.Connect("couchbase://127.0.0.1")
	if err != nil {
		fmt.Println("ERRROR CONNECTING TO CLUSTER:", err)
	}
	// Open Bucket
	bucket, err = cluster.OpenBucket("beer-sample", "")
	if err != nil {
		fmt.Println("ERRROR OPENING BUCKET:", err)
	}
	myQuery := gocb.NewN1qlQuery("SELECT * FROM `beer-sample` LIMIT 4")
	rows, err := bucket.ExecuteN1qlQuery(myQuery, nil)
	if err != nil {
		fmt.Printf("N1QL query error: %s\n", err)
	}

	var row interface{}
	for rows.Next(&row) {
		fmt.Printf("Row: %+v\n", row)
	}
	if err := rows.Close(); err != nil {
		fmt.Printf("N1QL query error: %s\n", err)
	}
}
