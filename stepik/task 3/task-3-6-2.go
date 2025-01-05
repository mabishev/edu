package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type EconomyActivity struct {
	GlobalID       int    `json:"global_id"`
	SystemObjectID string `json:"system_object_id"`
	SignatureDate  string `json:"signature_date"`
	Razdel         string `json:"Razdel"`
	Kod            string `json:"Kod"`
	Name           string `json:"Name"`
	Idx            string `json:"Idx"`
}

func main() {
	file, err := os.Open("./data-20190514T0100.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	var activities []EconomyActivity

	if err := json.Unmarshal(data, &activities); err != nil {
		fmt.Println(err)
		return
	}

	sum := 0

	for _, a := range activities {
		sum += a.GlobalID
	}

	fmt.Printf("%v", sum)

}
