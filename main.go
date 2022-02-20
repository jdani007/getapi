package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/james-daniels/getapi/geohazard"
)

func main() {
	resp, err := http.Get("https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/2.5_day.geojson")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var record geohazard.Earthquake
	if resp.StatusCode == 200 {
		err = json.NewDecoder(resp.Body).Decode(&record)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println(resp.StatusCode, "error received from the server.")
		resp.Body.Close()
		return
	}

	fmt.Println(record.GetMetadata("title"))

	fmt.Println(record.GetMetadata("count"))

	record.GetPlaces()

	record.GetMagnitude(4)

	record.GetFelt()

	record.GetCoordinates()

}
