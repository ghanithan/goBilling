package database

import (
	"log"
	"sync"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/ghanithan/goBilling/config"
)

// Module for couchbase db wrapper

type couchbaseDatabase struct {
	Db *gocb.Cluster
}

var (
	once       sync.Once
	dbInstance couchbaseDatabase
)

func InitCouchbaseDb(config *config.Config) couchbaseDatabase {
	once.Do(func() {
		options := gocb.ClusterOptions{
			Authenticator: gocb.PasswordAuthenticator{
				Username: config.Db.Username,
				Password: config.Db.Password,
			},
		}

		// Sets a pre-configured profile called "wan-development" to help avoid latency issues
		// when accessing Capella from a different Wide Area Network
		// or Availability Zone (e.g. your laptop).
		if err := options.ApplyProfile(gocb.ClusterConfigProfileWanDevelopment); err != nil {
			log.Fatal(err, "Error in profile")
		}

		// Initialize the Connection
		cluster, err := gocb.Connect("couchbase://"+config.Db.Host, options)
		if err != nil {
			log.Fatal(err, " connecting cluster")
		}

		// bucket := cluster.Bucket(config.Db.Reponame)

		err = cluster.WaitUntilReady(5*time.Second, nil)
		if err != nil {
			log.Fatal(err)
		}

		dbInstance = couchbaseDatabase{
			Db: cluster,
		}
	})
	return dbInstance
}

func (cb *couchbaseDatabase) GetDb() *gocb.Cluster {
	return cb.Db
}

func (cb *couchbaseDatabase) Close() {
	closeOpt := gocb.ClusterCloseOptions{}
	cb.Db.Close(&closeOpt)
}
