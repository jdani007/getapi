package geohazard

import "fmt"

type geometry struct {
	Type        string
	Coordinates []float64
	ID          string
}

type Earthquake struct {
	Metadata map[string]interface{}
	Features []struct {
		Properties map[string]interface{}
		Geometry   geometry
	}
}

func (e Earthquake) GetMetadata(m string) interface{}{
	return e.Metadata[m]
}

func (e Earthquake) GetPlaces() {
	for _, v := range e.Features {
		fmt.Println(v.Properties["place"].(string))
	}
}

func (e Earthquake) GetCoordinates() {
	for _, v := range e.Features {
		fmt.Println(v.Geometry.Coordinates)
	}
}

func (e Earthquake) GetMagnitude(mag float64) {
	for _, v := range e.Features {
		if v.Properties["mag"].(float64) >= mag {
			fmt.Println(v.Properties["place"].(string))
		}
	}
}

func (e Earthquake) GetFelt() {
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
