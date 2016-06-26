package main

import (
	"fmt"
	"log"

	"github.com/couchbase/gocb"
)

func main() {

	conn, err := gocb.Connect("couchbase://localhost")
	if err != nil {
		fmt.Println("Error:", err)
	}
	bucket, err := conn.OpenBucket("beer-sample", "")
	if err != nil {
		fmt.Println("ERRROR OPENING BUCKET:", err)
	}
	value := "test value"
	cas, error2 := bucket.Insert("document_name", &value, 0)
	if error2 != nil {
		log.Fatalf("Error setting data: %v", error2)
	}
	fmt.Printf("Inserted document CAS is `%08x`\n", cas)
	// var beer interface{}
	// scas, err := bucket.Get("williams_brothers_brewing_company-kelpie_seaweed_ale", &beer)
	// fmt.Printf("Got value `%+v` with CAS `%08x`\n", beer, scas)

	// if err != nil {
	// 	log.Fatalf("error getting data: %v", err)
	// }

	// for _, pn := range conn.Info.Pools {
	// 	fmt.Printf("Found pool:  %s -> %s\n", pn.Name, pn.URI)
	// }
	// pool, err := conn.GetPool("default")
	// if err != nil {
	// 	log.Fatalf("Error getting pool:  %v", err)
	// }
	//   	spew.Dump(pool)

	// var someValue interface{}
	// value := "test value"
	// cas, _ := bucket.Insert("document_name", &value, 0)
	// fmt.Printf("Inserted document CAS is `%08x`\n", cas)

	// // var someValue interface{}
	// // scas, _ := bucket.Get("document_name", &someValue)
	// // fmt.Printf("Got value `%+v` with CAS `%08x`\n", value, scas)
	// scas, err2 := bucket.Get("document_name", &someValue)
	// fmt.Printf("Got value `%+v` with CAS `%08x`\n", value, scas)

	// if err2 != nil {
	// 	log.Fatalf("Error getting bucket:  %v", err2)
	// }

	// var data interface{}
	// err = bucket.Get("12", &data)
	// if err != nil {
	// 	log.Printf("Error is: %v", err)
	// }
	// bucket.Set("someKey", 0, []string{"an", "example", "list"})
}
