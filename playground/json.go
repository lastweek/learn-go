package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Movie struct {
	Title string
	Year  int

	// variable with lower case initials are not exported
	// hence will not be included in the json output
	actors []string
}

var movies = []Movie{
	{Title: "movie1", Year: 2000, actors: []string{"aa", "bb"}},
	{Title: "movie2", Year: 2001, actors: []string{"cc", "dd"}},
}

func main() {
	fmt.Println(movies)

	// data, err := json.MarshalIndent(movies, "", "	")
	data, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", data)
}
