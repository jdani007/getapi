//Lesson I was learning in Python and wanted to recreate it in Go.

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

func (e earthquake)getMetadata(m string) interface{}{
	return e.Metadata[m]
}

func (e earthquake)getPlaces() string {
	var place string
	for _, v := range e.Features {
		place = place + "\n" + v.Properties["place"].(string)
	}
	return place
}

func (e earthquake)getMagnitude(mag float64) interface{} {
	var place string
	for _, v := range e.Features {
		if v.Properties["mag"].(float64) >= mag {
			place = place + "\n" + v.Properties["place"].(string)
		}
	}
	return place
}

func(e earthquake)getCoordinates() string {
	var coord string
	for _, v := range e.Features {
		coord = coord + "\n" + fmt.Sprint(v.Geometry.Coordinates)
	}
	return coord

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
	

	fmt.Println(record.getMetadata("title"))

	fmt.Println(record.getMetadata("count"))

	fmt.Println(record.getPlaces())

	fmt.Println(record.getCoordinates())

	fmt.Println(record.getMagnitude(4))

	// for _, v := range record.Features {
	// 	felt := v.Properties["felt"]
	// 	if felt != nil {
	// 		if felt.(float64) == 1 {
	// 			fmt.Println(v.Properties["place"], felt, "time")
	// 		} else if felt.(float64) > 1 {
	// 			fmt.Println(v.Properties["place"], felt, "times")
	// 		}
	// 	}
	// }
	// fmt.Printf("-----------------------\n\n")




}
