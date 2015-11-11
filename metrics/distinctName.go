// First one (distinctName) collects daily occurrences of distinct events in Redis.
// Metrics that are older than 30 days are merged into a monthly bucket, then
// cleared.

package metrics

import (
	"fmt"
	"log"
	"strings"
	"time"
)

const (
	limit = 30 * 24 * time.Hour
)

// Increase a metric counter in redis
func DistinctName(m Metric) {
	key := metricKey(m)
	err := Redis().Incr(key).Err()
	if err != nil {
		log.Printf("ERROR %s", err)
	} else {
		log.Printf("increase counter for %s", key)
	}
	ClearOlderMetrics()
}

// Archive metrics older than 30 days into a monthly bucket
// NOTE: this can be moved into a specific command so can be
//       called as cron job if needed
func ClearOlderMetrics() {
	keys := Redis().Keys("day*")
	for _, k := range keys.Val() {
		archiveAndClearIfNeeded(k)
	}
}

func archiveAndClearIfNeeded(key string) {

	date, metric, err := parseKey(key)
	if err != nil {
		log.Printf("ERROR %s", key)
		return
	}

	// check if date is 30 days old
	if time.Since(date) > limit {
		oldValue, _ := Redis().Get(key).Int64()
		newKey := metricMonthKey(metric, date)
		log.Printf("Archive key %s into %s", key, newKey)

		Redis().IncrBy(newKey, oldValue)
		Redis().Del(key)
	}
}

func parseKey(key string) (time.Time, string, error) {
	s := strings.Split(key, "_")
	date, err := time.Parse("2006-1-2", s[1])

	if err != nil {
		return time.Now(), "", err
	}

	return date, s[2], nil
}

func metricMonthKey(metric string, date time.Time) string {
	return fmt.Sprintf("month_%d_%d_%s", date.Year(), date.Month(), metric)
}

func metricKey(m Metric) string {
	today := time.Now()
	return fmt.Sprintf("day_%d-%d-%d_%s", today.Year(), today.Month(), today.Day(), m.Metric)
}
