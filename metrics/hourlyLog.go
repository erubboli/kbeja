// Second one (hourlyLog) collects all items that occurred in the last hour into
// MongoDB (Bonus: how would you calculate the average of incoming values in a
// distributed environment?)

// NOTE: Index with expire at 1hour is set right after the database connection using
//       EnsureIndex

// To calculate the average of the values it's possible to use map-reduce as described here:
// https://docs.mongodb.org/manual/core/map-reduce/
package metrics

import (
	"log"
)

// Insert metrics into HourlyLog mongo collection
func HourlyLog(m Metric) {
	err := MongoCollection().Insert(m)
	if err != nil {
		log.Printf("ERROR %s", err)
	} else {
		log.Printf("Insert metric into HourlyLog " + m.String())
	}
}
