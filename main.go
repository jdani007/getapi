package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type geometry struct {
	Type string
	Coordinates []float64
	ID string
}

type earthquake struct {
	Metadata map[string]interface{}
	Bbox []float64
	Features []struct {
		Properties map[string]interface{}
		Geometry geometry
	}

}

func main() {
	resp, err := http.Get("https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/2.5_day.geojson")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var record earthquake

	err = json.NewDecoder(resp.Body).Decode(&record)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(record.Bbox)
	fmt.Printf("-----------------------\n\n")
	

	fmt.Println(record.Metadata["title"])
	fmt.Printf("-----------------------\n\n")


	fmt.Println(record.Metadata["count"])
	fmt.Printf("-----------------------\n\n")


	for _, v := range record.Features {
		fmt.Println(v.Properties["place"])
	}
	fmt.Printf("-----------------------\n\n")


	for _, v := range record.Features {
		if v.Properties["mag"].(float64) >= 4.0 {
			fmt.Println(v.Properties["place"])
		}
	}
	fmt.Printf("-----------------------\n\n")


	for _, v := range record.Features {
		felt := v.Properties["felt"]
		if felt != nil {
			if felt.(float64) > 0 {
				fmt.Println(v.Properties["place"], felt, "times")
			}
		}
	}
	fmt.Printf("-----------------------\n\n")

		for _, v := range record.Features {
		fmt.Println(v.Geometry.Coordinates)
	}

}
