// Implement datastores helpers

// NOTE 1: There are many redis libraries out there, I've just picked the
//         one with more stars on github since I have no time to compare
//         them at the moment.

// NOTE 2: All Connection information are hardcoded, it will eventually go
//         into a configuration file

package metrics

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gopkg.in/mgo.v2"
	"gopkg.in/redis.v3"
)

// As stated here: https://github.com/go-redis/redis/blob/v3.2.13/redis.go#L179
//
// Client is a Redis client representing a pool of zero or more
// underlying connections. It's safe for concurrent use by multiple
// goroutines.

var redisClient *redis.Client

// Initialize or retrieve a redis connection
func Redis() *redis.Client {

	/* Eventually switch to Failover client */
	if redisClient == nil {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
	}

	return redisClient
}

// Mongo DB Session
var mongoSession *mgo.Session

// Initialize or retrieve a Mongo DB Session
func MongoSession() *mgo.Session {
	if mongoSession == nil {
		var err error
		mongoSession, err = mgo.Dial("localhost")
		failOnError(err, "Cannot connect to MongoDB.")
		ensureMongoIndex()
	}

	return mongoSession
}

var mongoHourlyLog *mgo.Collection

// Return MongoDB HourlyLog collection in a persistent connection
func MongoCollection() *mgo.Collection {
	if mongoHourlyLog == nil {
		mongoHourlyLog = MongoSession().DB("koding").C("hourlyLog")
	}

	return mongoHourlyLog
}

// Ensure the collection has a proper index with expire time
func ensureMongoIndex() {

	err := MongoCollection().EnsureIndex(mgo.Index{
		Key:         []string{"createdat"},
		Unique:      false,
		DropDups:    false,
		Background:  true,
		Sparse:      true,
		ExpireAfter: 60, // in seconds
	})

	if err != nil {
		log.Println("ERROR: unable to create mongo index")
	}
}

// Postgres Database handler
var dbh *sql.DB

// Return a persistent database handler
func Postgres() *sql.DB {
	if dbh == nil {
		var err error
		dbh, err = sql.Open("postgres", "dbname=metrics sslmode=disable")
		failOnError(err, "Cannot connect to postgres database.")
	}
	return dbh
}

// log error and die
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
