package main

import "encoding/json"
//import "errors"
import "fmt"
import "net/http"
import "time"

// TODO add logging

//var servers = [3]string{"https://data-dev.petabencana.id", "https://data-dev.riskmap.in", "https://data-dev.riskmap.us"}

// list of servers to poll
var servers = [2]string{"https://data-dev.petabencana.id/reports/?timeperiod=1&geoformat=geojson", "https://data-dev.riskmap.in/reports/?timeperiod=1&geoformat=geojson"}

// http client
var client = &http.Client{Timeout: 10 * time.Second}

// states
var mapState = map[int]string{
	0: "RED",
	1: "GREEN",
	2: "BLUE",
}

// GeoJSON feature results
type Result struct {
	Type string `json: "type"`
	Features []interface{} `json: "features"`
}

// GeoJSON reports object
type Reports struct {
	StatusCode float64 `json: "statusCode"`
	Result Result `json: "result"`
}

// Get JSON data from server
// TODO return error
// TODO return JSON object proper
// TODO documentation standard
// TODO code layout
func getJson(url string) (Reports, error) {

	reports := Reports{}

	// Send request
	resp, err := client.Get(url)
	if err!= nil {
		return reports, err
	}

	// Close output
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&reports)
	// Return
	return reports, nil
}

// Poll a server to get flood state
func serverStatus(server string) int {
	// List server
	fmt.Printf("Polling " + server + "\n")

	// Get reports
	reports, err := getJson(server)

	// Check state
	if err == nil && reports.StatusCode == 200 && len(reports.Result.Features) == 0 {
		return 1
	} else if err == nil && reports.StatusCode == 200 && len(reports.Result.Features) > 0 {
		return 2
	}
	// Default
	return 0
}

func main(){

	state := 0 // Default state
	for _, name := range servers {
		val := serverStatus(name)
		//fmt.Printf("%s Code %s....%s\n", name, mapState[state], time.Now())
		if val == 0 {
			state = val // always set state 0
		} else if val > state {
			state = val // raise state
		}
	}
	// Print out final state
	fmt.Printf("Code %s....%s\n", mapState[state], time.Now())
}
