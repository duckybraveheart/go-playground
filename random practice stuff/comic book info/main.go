package main

import "fmt"

func main() {
	var publisher string
	var writer string
	var artist string
	var title string
	var year uint
	var pageNumber uint
	var grade float32

	title = "Mr. GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "DizzyBooks Publishing Inc."
	year = 1997
	pageNumber = 14
	grade = 6.5

	fmt.Println(title, "Written by:", writer, "Drawn by:", artist, "Published by:", publisher, "Published in:", year, "Number of pages:", pageNumber, "Condition:", grade)
}
