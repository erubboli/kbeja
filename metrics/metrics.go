//Data sample:
// {
//  "username": "kodingbot",  // string
//  "count": 12412414,    // int64
//  "metric": "kite_call" // string
// }

package metrics

import (
	"fmt"
	"time"
)

type Metric struct {
	CreatedAt time.Time
	Username  string
	Count     int64
	Metric    string
}

func (m Metric) String() string {
	return fmt.Sprintf("Username:%s;Count:%d;Metric:%s", m.Username, m.Count, m.Metric)
}
