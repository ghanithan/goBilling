package database

import (
	"log"
	"sync"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/ghanithan/goBilling/config"
	"github.com/ghanithan/goBilling/instrumentation"
)

// Module for couchbase db wrapper

type couchbaseDatabase struct {
	Db     *gocb.Cluster
	Bucket *gocb.Bucket
	logger instrumentation.GoLogger
}

var (
	once       sync.Once
	dbInstance couchbaseDatabase
)

func InitCouchbaseDb(config *config.Config, logger instrumentation.GoLogger) couchbaseDatabase {
	defer logger.TimeTheFunction(time.Now(), "InitCouchbaseDb")
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

		bucket := cluster.Bucket(config.Db.Reponame)
		err = bucket.WaitUntilReady(5*time.Second, nil)
		if err != nil {
			log.Fatal(err)
		}

		dbInstance = couchbaseDatabase{
			Db:     cluster,
			Bucket: bucket,
			logger: logger,
		}
	})
	return dbInstance
}

func (cb *couchbaseDatabase) GetDbInstance() *gocb.Cluster {
	return cb.Db
}

func (cb *couchbaseDatabase) GetDb() *gocb.Bucket {
	return cb.Bucket
}

func (cb *couchbaseDatabase) Close() {
	closeOpt := gocb.ClusterCloseOptions{}
	cb.Db.Close(&closeOpt)
}

func (cb *couchbaseDatabase) FetchById(collection string, id string) (*gocb.GetResult, error) {
	getResult, err := cb.Bucket.Collection(collection).Get(id, &gocb.GetOptions{})
	if err != nil {
		cb.logger.Error("Issue in fetching from collection: %q", err)
		return nil, err
	}

	return getResult, nil
}

func ExtractContent[T any](getResult *gocb.GetResult, extractedOutput *T, logger instrumentation.GoLogger) error {
	err := getResult.Content(getResult)
	if err != nil {
		logger.Error("Issue in unmarsheling: %q", err)
		return err
	}
	return nil

}
