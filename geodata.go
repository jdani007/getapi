package main

import "fmt"

type geometry struct {
	Type        string
	Coordinates []float64
	ID          string
}

type earthquake struct {
	Metadata map[string]interface{}
	Features []struct {
		Properties map[string]interface{}
		Geometry   geometry
	}
}

func (e earthquake) getMetadata(m string) interface{}{
	return e.Metadata[m]
}

func (e earthquake) getPlaces() {
	for _, v := range e.Features {
		fmt.Println(v.Properties["place"].(string))
	}
}

func (e earthquake) getCoordinates() {
	for _, v := range e.Features {
		fmt.Println(v.Geometry.Coordinates)
	}
}

func (e earthquake) getMagnitude(mag float64) {
	for _, v := range e.Features {
		if v.Properties["mag"].(float64) >= mag {
			fmt.Println(v.Properties["place"].(string))
		}
	}
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
}
