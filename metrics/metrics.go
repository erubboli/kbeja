//Data sample:
// {
//  "username": "kodingbot",  // string
//  "count": 12412414,    // int64
//  "metric": "kite_call" // string
// }

package metrics

import (
	"time"
  "fmt"
  "encoding/json"
)

type Metric struct {
  CreatedAt time.Time `json:"-"`
  Username  string `json:"username"`
  Count     int64 `json:"count"`
  Metric    string `json:"metric"`
}

func (m Metric) String() string {
  return fmt.Sprintf("Username:%s;Count:%d;Metric:%s", m.Username, m.Count, m.Metric)
}

func (m Metric) JSON() []byte {
  json, _ := json.Marshal(m)
  return json
}
