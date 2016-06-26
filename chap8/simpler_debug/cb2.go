package main

import (
	"fmt"
	"log"

	"github.com/couchbase/gocb"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	myCluster, _ := gocb.Connect("couchbase://127.0.0.1")
	myBucket, _ := myCluster.OpenBucket("beer-sample", "")

	var beer map[string]interface{}
	cas, err := myBucket.Get("abbaye_notre_dame_du_st_remy-rochefort_8", &beer)
	if err != nil {
		log.Printf("Error is : %v", err)
	}
	spew.Dump(beer)
	fmt.Printf("Got document CAS is %v\n", cas)

	myBucket, _ = myCluster.OpenBucket("default", "")
	myDoc := "Hello World!"
	cas2, _ := myBucket.Insert("document_name", &myDoc, 0)

	fmt.Printf("Got document CAS is %v\n", cas2)
	// cas, _ := myBucket.Get("aass_brewery-juleol", &beer)

	// beer["comment"] = "Random beer from Nooooomyway"

	// myBucket.Replace("aass_brewery-juleol", &beer, cas, 0)

}
