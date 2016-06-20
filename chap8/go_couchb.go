package main

import (
	"fmt"

	"github.com/couchbase/gocb"
)

// bucket reference - reuse as bucket reference in the application
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
	// value := "test value"
	// cas, _ := bucket.Insert("document_name", &value, 0)
	// fmt.Printf("Inserted document CAS is `%08x`\n", cas)

	// var someValue interface{}
	// scas, _ := bucket.Get("document_name", &someValue)
	// fmt.Printf("Got value `%+v` with CAS `%08x`\n", value, scas)

	vq := gocb.NewViewQuery("dev_beer", "by_name") //.Skip(6).Limit(3)
	rows, err := bucket.ExecuteViewQuery(vq)

	if err := rows.Close(); err != nil {
		fmt.Printf("View query error: %s\n", err)
	}
	var row interface{}
	for rows.Next(&row) {
		fmt.Printf("Got row `%+v`\n")
	}
	fmt.Println("Example Successful - Exiting")
}
