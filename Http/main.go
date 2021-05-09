package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Salad struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Ingredients []string `json:"ingredients"`
	Photo       string   `json:"photo"`
}

func main() {

	/* GET */

	resp, err := http.Get("http://localhost:3000/")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	/* First implementation */

	bs := make([]byte, 99999) // create byte slice with capacity 99999
	a, e := resp.Body.Read(bs)
	fmt.Println(string(bs))
	fmt.Println(a, e)

	/* Advanced implementation */

	/* 1) read from resp.Body and save into buffer
	 * 2) then from buffer write into stdout of type File
	 */
	io.Copy(os.Stdout, resp.Body)

	/* POST */

	values := &Salad{
		Name:        "Salad",
		Description: "garden salad",
		Ingredients: []string{"salad", "tomato", "oignon"},
		Photo:       "photo",
	}
	json_data, _ := json.Marshal(values)
	fmt.Println(string(json_data))

	postresp, _ := http.Post("http://localhost:3000/addSalad", "application/json",
		bytes.NewBuffer(json_data))

	fmt.Println(postresp)
}
