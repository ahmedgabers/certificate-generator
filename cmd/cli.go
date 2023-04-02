package cmd

import (
	"flag"
	"fmt"
	"os"
)

func printHelp() {
	helpText := `Usage: certgen [options]

Options:
  -f, --filename FILE    JSON file containing the certificate data (default: "names.json")
  -o, --output DIR       The directory where the generated HTML files will be created (default: ".")
  -t, --template FILE    HTML template file to use for generating certificates (default: "certificate.html")

  -h, --help             Show this help message
`
	fmt.Println(helpText)
}

func parseFlags() (string, string, string) {
	filename := flag.String("filename", "names.json", "JSON file containing the certificate data")
	outputDir := flag.String("output", ".", "The directory where the generated HTML files will be created")
	templateFile := flag.String("template", "certificate.html", "HTML template file to use for generating certificates")
	help := flag.Bool("help", false, "Show help message")

	flag.StringVar(filename, "f", "names.json", "JSON file containing the certificate data (shorthand)")
	flag.StringVar(outputDir, "o", ".", "The directory where the generated HTML files will be created (shorthand)")
	flag.StringVar(templateFile, "t", "certificate.html", "HTML template file to use for generating certificates (shorthand)")
	flag.BoolVar(help, "h", false, "Show help message (shorthand)")

	flag.Parse()

	if *help {
		printHelp()
		os.Exit(0)
	}

	return *filename, *outputDir, *templateFile
}

func Execute() (string, string, string) {
	filename, outputDir, templateFile := parseFlags()

	// Return the filename, outputDir, and templateFile values
	return filename, outputDir, templateFile
}
