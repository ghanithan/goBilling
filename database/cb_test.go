package database

import (
	"fmt"
	"testing"

	"github.com/couchbase/gocb/v2"
	"github.com/ghanithan/goBilling/config"
)

func TestInitCouchbaseDb(t *testing.T) {
	config, err := config.GetConfig("../settings/sample.yaml")
	if err != nil {
		t.Fatalf("Error in reading config: %s", err)
	}

	db := InitCouchbaseDb(config)

	cluster := db.GetDb()

	// col := bucket.Scope("_default").Collection("customers")

	// // Get the document back
	// getResult, err := col.Get("custid:C1", nil)
	// if err != nil {
	// 	t.Errorf("Error fetching: %s", err)
	// }

	qOpts := gocb.QueryOptions{}

	// create query
	queryStr := "select * from default:vizha._default.customers"

	fmt.Printf("query: %v\n", queryStr)

	rows, err := cluster.Query(queryStr, &qOpts)
	if err != nil {
		panic(err)
	}

	fmt.Printf("rows: %v\n", rows)

	for rows.Next() {
		var intfc interface{}
		err = rows.Row(&intfc)
		if err != nil {
			panic(err)
		}
		fmt.Printf("interface result: %v\n", intfc)
	}

	// type UserType struct {
	// 	Id      string `json:custid`
	// 	Name    string `json: name`
	// 	Address struct {
	// 		Street  string `json: street`
	// 		City    string `json: city`
	// 		Zipcode string `json: zipcode`
	// 	} `json: address`
	// }

	// user := &UserType{}
	// err = getResult.Content(&user)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("User: %v\n", user)

}
