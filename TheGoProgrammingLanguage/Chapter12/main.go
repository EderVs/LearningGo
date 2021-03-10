package main

import "fmt"

// Movie contains information of a movie.
type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
	Example         map[struct{ E string }]string
	ExampleArray    map[[2]string]string
	OtherMovie      *Movie
}

func main() {
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "Geroge C. Scoot",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
		Example:      map[struct{ E string }]string{{E: "hello"}: "world"},
		ExampleArray: map[[2]string]string{{"hello", "hello2"}: "world"},
	}
	//strangelove.OtherMovie = &strangelove
	Display("strangeLove", strangelove)
	b, err := Marshal(strangelove)
	if err != nil {
		fmt.Printf("there was an error in Marshal(%v): %v\n", strangelove, err)
	}
	fmt.Printf("%v\n", string(b))
}
