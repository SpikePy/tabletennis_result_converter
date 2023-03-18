package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	filePattern := "Vereinsspielplan_*.csv"
	include_unplayed := false

	filepaths, err := filepath.Glob(filePattern)
	if err != nil {
		fmt.Println("Error finding files:", err)
		return
	}
	if len(filepaths) == 0 {
		fmt.Printf("No file matching pattern '%s' found\n", filePattern)
		return
	}
	firstFile := filepaths[0]

	// Open the file
	inputFile, err := os.Open(firstFile)
	if err != nil {
		fmt.Printf("Error opening file '%s': %s\n", firstFile, err)
		return
	}
	defer inputFile.Close()

	// Create the output HTML file
	outputFile, err := os.Create("Vereinsspielplan.html")
	if err != nil {
		log.Fatal("Error creating output file:", err)
	}
	defer outputFile.Close()

	// Create a CSV reader for the input file
	reader := csv.NewReader(inputFile)
	reader.Comma = ';'

	// Read and discard the header line
	_, err = reader.Read()
	if err != nil {
		log.Fatal("Error reading CSV file:", err)
	}

	// Write the HTML table header to the output file
	fmt.Fprintf(outputFile, "<table>\n    <tr><th>Termin</th><th>Staffel</th><th>Heim-Verein</th><th>Gast-Verein</th><th>Ergebnis</th></tr>\n")

	// Loop over the remaining lines in the input file
	for {
		// Read the next line of the CSV file
		record, err := reader.Read()
		if err != nil {
			// If we've reached the end of the file, break out of the loop
			if err == io.EOF {
				break
			}
			log.Fatal("Error reading CSV file:", err)
		}

		// Write the extracted columns to the output file as an HTML table row
		if record[27] != "0" && record[28] != "0" {
			fmt.Fprintf(outputFile, "    <tr><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s:%s</td></tr>\n", record[0], record[7], record[20], record[26], record[27], record[28])
		} else if include_unplayed == true {
			fmt.Fprintf(outputFile, "    <tr><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>-</td></tr>\n", record[0], record[7], record[20], record[26])
		}
	}

	// Write the HTML table footer to the output file
	fmt.Fprintf(outputFile, "</table>\n")
}
