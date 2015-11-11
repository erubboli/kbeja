// Third one (accountName) collects all the account names that sent metrics, with
// their first occurrence datetime (UTC) into PostgreSQL.

// NOTE: Database data structure can be created with the following commands:
// go get github.com/mattes/migrate
// `migrate -url pg://localhost/metrics\?sslmode=disable -path ./db up`

package metrics

import "log"

//Save account first occurrence's datatime
func AccountName(m Metric) {
	var userId int
	Postgres().QueryRow("SELECT id FROM users WHERE username=$1 LIMIT 1", m.Username).Scan(&userId)

	if userId != 0 {
		log.Printf("Username %s already present", m.Username)
	} else {
		log.Println("Saving %s into users list", m.Username)
		_, err := Postgres().Exec("INSERT INTO users (username,created_at) VALUES($1, $2)", m.Username, m.CreatedAt.UTC())
		if err != nil {
			log.Printf("ERROR %s", err)
		}
	}
}
