// Package fail - Flood Alert Indicator Light
package main

import "log"
import "bytes"
import "fmt"

// TODO add config

var buf bytes.Buffer
var	logger = log.New(&buf, "log ", log.Ldate | log.Ltime | log.LUTC)



// List of servers to poll
var servers = []string{"https://data-dev.petabencana.id/reports/?timeperiod=1&geoformat=geojson", "https://data-dev.riskmap.in/reports/?timeperiod=1&geoformat=geojson"}

func main(){

	logger.Print("Application starting")

	// Get combined state across all servers
	state := PollState(servers)

	// Print
	//fmt.Printf("Code %s....%s\n", state, time.Now())
	logger.Printf("Code %s\n", state)
	fmt.Print(&buf)

}
