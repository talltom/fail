// Package fail - state.go Determines state of servers from responses
package main

import "net/http"
import "time"

// http client
var client = &http.Client{Timeout: 10 * time.Second}

// states
var States = map[int]string{
	0: "RED",
	1: "GREEN",
	2: "BLUE",
}

// Poll a server to get flood state
func getState(server string) int {

	// Get reports
	reports, err := GetJSON(server)

	// Check state
	if err == nil && reports.StatusCode == 200 && len(reports.Result.Features) == 0 {
		return 1
	} else if err == nil && reports.StatusCode == 200 && len(reports.Result.Features) > 0 {
		return 2
	}
	// Default
	return 0
}

func PollState(servers []string) string {

  state := 0
  for _, name := range servers {
    val := getState(name)
    if val == 0 {
      state = 0 // always set state 0
    } else if val > state {
      state = val // raise state
    }
  }
  return States[state]
}
