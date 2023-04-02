package main

import (
	"bytes"
	"certgen/cmd"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

// CertificateData holds the data you want to replace in the template.
type CertificateData struct {
	Name          string `json:"name"`
	Hours         string `json:"hours"`
	CourseTitle   string `json:"course_title"`
	DateCompleted string `json:"date_completed"`
	Signature     string
}

// NamesList holds the list of certificate data read from the JSON file.
type NamesList struct {
	Certificates []CertificateData `json:"certificates"`
	Signature    string            `json:"signature"`
}

func main() {

	filename, outputDir, templateFile := cmd.Execute()

	// Read the HTML template
	htmlBytes, err := os.ReadFile(templateFile)
	if err != nil {
		fmt.Println("Error reading HTML template:", err)
		return
	}

	// Create a new template and parse the HTML
	tmpl, err := template.New("certificate").Parse(string(htmlBytes))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	// Read the JSON file containing the certificate data
	jsonBytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Unmarshal the JSON data into a NamesList struct
	var namesList NamesList
	err = json.Unmarshal(jsonBytes, &namesList)
	if err != nil {
		fmt.Println("Error unmarshaling JSON data:", err)
		return
	}

	// Create the output directory if it does not exist
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		fmt.Println("Error creating output directory:", err)
		return
	}

	// Iterate over the certificate data and generate a certificate for each entry
	for _, data := range namesList.Certificates {

		data.Signature = namesList.Signature

		// Execute the template with the data
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, data)
		if err != nil {
			fmt.Println("Error executing template:", err)
			return
		}

		// Replace spaces in the name with underscores and convert to lowercase
		modifiedName := strings.ReplaceAll(data.Name, " ", "_")
		modifiedName = strings.ToLower(modifiedName)

		// Write the generated HTML to a file in the specified output directory
		outputFilename := filepath.Join(outputDir, fmt.Sprintf("certificate_%s.html", modifiedName))
		err = os.WriteFile(outputFilename, buf.Bytes(), 0644)
		if err != nil {
			fmt.Println("Error writing generated HTML to file:", err)
			return
		} else {
			fmt.Printf("Certificate generated successfully for %s\n", data.Name)
		}
	}
}
