package database

import (
	"fmt"
	"testing"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/ghanithan/goBilling/config"
	"github.com/ghanithan/goBilling/instrumentation"
)

type userType struct {
	Id      string `json:custid`
	Name    string `json: name`
	Address struct {
		Street  string `json: street`
		City    string `json: city`
		Zipcode string `json: zipcode`
	} `json: address`
}

func TestInitCouchbaseDb(t *testing.T) {
	logger := instrumentation.InitInstruments()
	defer logger.TimeTheFunction(time.Now(), "TestInitCouchbaseDb")

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	config, err := config.GetConfig("../settings/default.yaml")
	if err != nil {
		t.Fatalf("Error in reading config: %s", err)
	}

	db := InitCouchbaseDb(config, logger)

	cluster := db.GetDbInstance()

	// col := bucket.Scope("_default").Collection("customers")

	// Get the document back
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

	getResult, _ := db.GetDb().Collection("customers").Get("C1", &gocb.GetOptions{})

	user := &userType{}

	err = getResult.Content(&user)
	if err != nil {
		t.Errorf("Issue in unmarsheling: %q", err)

	}

	fmt.Println("output: ", user)

	user1 := &userType{}

	getResult, err = db.FetchById("customers", "C1")

	if err != nil {
		t.Errorf("Error in FetchById: %s", err)
	}

	if err := ExtractContent(getResult, user1, logger); err != nil {
		t.Errorf("Error in extracting content: %s", err)

	}
	t.Errorf("output: %s", user1)

	// user := &UserType{}
	// err = getResult.Content(&user)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("User: %v\n", user)

}
