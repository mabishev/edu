package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type StudentGroup struct {
	Students []Student
}

type Student struct {
	Rating []int `json:"rating"`
}

type StudentStats struct {
	Average float64
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var group StudentGroup
	if err := json.Unmarshal(data, &group); err != nil {
		log.Fatal(err)
	}

	var ratings int

	for _, stud := range group.Students {
		ratings += len(stud.Rating)
	}

	avgRatings := float64(ratings) / float64(len(group.Students))

	stats := StudentStats{Average: avgRatings}

	resp, err := json.MarshalIndent(stats, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(resp)
}
