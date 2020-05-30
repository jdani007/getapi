### getapi

```Go
package main

import (
"fmt"
"encoding/json"
"net/http"
)

type person struct {
	Data []struct {
		ID int `json:"id"`
		Email string `json:"email"`
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
		Avatar string `json:"avatar"`
	}
}


func main() {
	resp, err := http.Get("https://reqres.in/api/users/")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var record person

	err = json.NewDecoder(resp.Body).Decode(&record)
	if err != nil {
		fmt.Println(err)
	}

	for i ,v := range record.Data {
		fmt.Println(i, "Name:", v.FirstName, v.LastName)
	}
}
```