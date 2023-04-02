#!/bin/bash

# Default values
output_dir="pdf_output"
parallel_processes=4
html_dir=""

# Usage message
usage() {
  echo "Usage: $0 [OPTIONS]"
  echo
  echo "Converts HTML files in the specified directory into PDF files using wkhtmltopdf."
  echo
  echo "Options:"
  echo "  -d, --html-dir DIR         Set the input directory containing HTML files (required)."
  echo "  -o, --output-dir DIR       Set the output directory for the converted PDF files (default: $output_dir)."
  echo "  -p, --parallel-processes N Set the number of parallel processes to run simultaneously (default: $parallel_processes)."
  echo "  -h, --help                 Display this help message and exit."
}

# Parse command-line arguments
while [ "$#" -gt 0 ]; do
  case "$1" in
    -d|--html-dir)
      html_dir=$(realpath "$2")
      shift 2
      ;;
    -o|--output-dir)
      output_dir=$(realpath "$2")
      shift 2
      ;;
    -p|--parallel-processes)
      parallel_processes="$2"
      shift 2
      ;;
    -h|--help)
      usage
      exit 0
      ;;
    *)
      echo "Unknown option: $1"
      usage
      exit 1
      ;;
  esac
done

# Check if the input HTML directory is provided
if [ -z "$html_dir" ]; then
  echo "Error: The input HTML directory must be specified using -d or --html-dir option."
  usage
  exit 1
fi

# Check if the input HTML directory is readable
if [ ! -r "$html_dir" ]; then
  echo "Error: Input HTML directory $html_dir is not readable."
  exit 1
fi

# Creating the output directory if it doesn't exist

if [ ! -d "$output_dir" ]; then
  mkdir -p "$output_dir"
fi

# Check if the output directory is writable
if [ ! -w "$output_dir" ]; then
  echo "Error: Output directory $output_dir is not writable."
  exit 1
fi

# Define a function to convert HTML to PDF
convert_to_pdf() {
  file_path="$1"
  file=$(basename "$file_path")
  filename="${file%.*}"
  wkhtmltopdf --enable-local-file-access "$file_path" "${output_dir}/$filename.pdf"
}

export -f convert_to_pdf
export output_dir

# Iterate through the HTML files and convert them to PDF using xargs to run multiple instances in parallel
find "$html_dir" -maxdepth 1 -name "*.html" -print0 | xargs -0 -I {} -P "$parallel_processes" bash -c 'convert_to_pdf "{}"'
