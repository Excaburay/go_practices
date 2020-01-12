package main

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/common/log"
	"go_project/config"
	"go_project/learntest"
	"go_project/stringutil"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	fmt.Println("stringUtil.Revers: ", stringutil.Reverse("!oG ,olleH"))
	fmt.Println("config test: This is name ", config.GetName())

	strings := []string{"one", "two", "", "three"}
	fmt.Println("no empty in string ", learntest.NoEmpty(strings))

	slice := []int{5, 4, 3, 2, 1}
	fmt.Println("reverse element ", learntest.ReverseSlice(slice))

	fmt.Println("remove element ", learntest.RemoveElement(slice, 2))
	fmt.Println(slice)

	ages := make(map[string]int)
	ages["alice"] = 32
	ages["charlie"] = 45
	fmt.Println("is map exist bob", learntest.IsExistElement(ages))

	ages2 := map[string]int{
		"alice":   32,
		"charlie": 45,
	}
	fmt.Println("is ages and ages2 are same ", learntest.MapEqual(ages, ages2))

	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Kacsel", Year: 1980, Color: true,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Govel", Year: 2020, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	}
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatal("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	data, err = json.MarshalIndent(movies, "", "	 ")
	if err != nil {
		log.Fatal("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)

}
