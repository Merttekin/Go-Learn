package main

import "fmt"

func main() {
	languages := []string{"Javascript", "Phyton", "Ruby"}
	languages = append(languages, "Go")

	fmt.Println(languages)
	fmt.Println(languages[1:3])

	for i, language := range languages {
		fmt.Printf(
			"Language slice %s\n",
			i,
			language,
		)
	}
}
