// Package fail - geojson.go Gets data from servers
package main

import "encoding/json"
import "fmt"


// Geojson feature results
type Result struct {
	Type string `json: "type"`
	Features []interface{} `json: "features"`
}

// Geojson reports object
type Reports struct {
	StatusCode float64 `json: "statusCode"`
	Result Result `json: "result"`
}

// Get json data from server
// TODO return error
// TODO return json object proper
// TODO documentation standard
// TODO code layout
func GetJSON(address string) (Reports, error) {

	// Create empty json reports object
	reports := Reports{}

  logger.Printf("Contacting server %s\n", address)

	// Send request
	resp, err := client.Get(address)
	if err!= nil {
		return reports, err
	}

	// Close output
	defer resp.Body.Close()

	// Decode json stream
	err = json.NewDecoder(resp.Body).Decode(&reports)
	if err != nil {
		fmt.Printf("Error %+v\n", err)
		return reports, err // On decode error return empty reports with error
	}
	// Return
	return reports, nil
}
