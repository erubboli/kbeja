package main

import (
	"flag"
	"fmt"

	"github.com/erubboli/kbeja/workers"
	"github.com/erubboli/kbeja/metrics"
)

var startDistinctNameWorker = flag.Bool("distinctName", false, "Start Distinct Name Worker")
var startHourlyLogWorker = flag.Bool("hourlyLog", false, "Start Hourly Log Worker")
var startAccountNameWorker = flag.Bool("accountName", false, "Start Account Name Worker")
var sendMessageToQueue = flag.Bool("sendMessage", false, "Send Message [username] [metric]")

func main() {
	flag.Parse()

	if *startDistinctNameWorker {
    workers.Execute(metrics.DistinctName, "DistinctName")

	} else if *startHourlyLogWorker {
    workers.Execute(metrics.HourlyLog, "HourlyLog")

	} else if *startAccountNameWorker {
    workers.Execute(metrics.AccountName, "AccountName")

	} else if *sendMessageToQueue {
    SendMessage()

	} else {
		fmt.Println("Usage:")
		flag.PrintDefaults()
	}
}
