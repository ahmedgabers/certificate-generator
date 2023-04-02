# HTML to PDF Converter

This script converts HTML files in a specified directory into PDF files using wkhtmltopdf.

## Requirements

- wkhtmltopdf (<https://wkhtmltopdf.org/>)
- Bash (Unix shell)

## Usage

1. Make sure you have a directory containing the HTML files you want to convert.

2. Run the script:

    ```bash
    ./html_to_pdf.sh -d HTML_DIR [OPTIONS]
    ```

    This will convert all HTML files in the specified input directory (HTML_DIR) into PDF files and store them in the `pdf_output` directory by default. The input directory can be specified using absolute or relative paths.

    Available options:

    `-d`, `--html-dir DIR` Set the input directory containing HTML files (required).
    `-o`, `--output-dir DIR` Set the output directory for the converted PDF files (default: pdf_output).
    `-p`, `--parallel-processes N`Set the number of parallel processes to run simultaneously (default: 4).
    `-h`, `--help` Display this help message and exit.
