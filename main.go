

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
	Features []struct {
		Properties map[string]interface{}
		Geometry geometry
	}

}

func (e earthquake)getMetadata(m string){
	fmt.Println(e.Metadata[m])
	fmt.Println()
}

func (e earthquake)getPlaces() {
	for _, v := range e.Features {
		fmt.Println(v.Properties["place"].(string)) 
	}
	fmt.Println()
}

func (e earthquake)getMagnitude(mag float64){
	for _, v := range e.Features {
		if v.Properties["mag"].(float64) >= mag {
			fmt.Println(v.Properties["place"].(string))
		}
	}
	fmt.Println()

}

func (e earthquake) getFelt() {
	for _, v := range e.Features {
		felt := v.Properties["felt"]
		if felt != nil {
			if felt.(float64) == 1 {
				fmt.Println(v.Properties["place"], felt, "time")
			} else if felt.(float64) > 1 {
				fmt.Println(v.Properties["place"], felt, "times")
			}
		}
	}
	fmt.Println()
}

func(e earthquake)getCoordinates(){
	for _, v := range e.Features {
		fmt.Println(v.Geometry.Coordinates)
	}
}


func main() {
	resp, err := http.Get("https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/2.5_day.geojson")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var record earthquake
	if resp.StatusCode == 200 {
		err = json.NewDecoder(resp.Body).Decode(&record)

		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(resp.StatusCode, "error received from the server.", )
		resp.Body.Close()
		return
	}

	
	record.getMetadata("title")

	record.getMetadata("count")

	record.getPlaces()

	record.getMagnitude(4)

	record.getFelt()

	record.getCoordinates()

}
