# Certificate Generator

The Certificate Generator is a command-line tool written in Go that generates personalized HTML certificates based on a provided JSON file and an HTML template.

## Features

- Generates personalized HTML certificates.
- Reads certificate data from a JSON file.
- Customizable HTML template for certificates.
- Supports command-line flags for input/output configuration.

## Prerequisites

- Go 1.17 or later

## Installation

1. Go to the [releases page](https://github.com/ahmedgabers/certificate-generator/releases) of this repository.

1. Download the appropriate binary for your operating system.

1. Make the binary executable, for example:

    ```bash
      chmod +x certgen-linux-amd64
    ```

## Usage

```bash
certgen --filename <JSON_FILE> --template <HTML_TEMPLATE> --output <OUTPUT_DIR>
```

- `<JSON_FILE>`: The path to the JSON file containing the certificate data (default: certificates.json).
- `<HTML_TEMPLATE>`: The path to the HTML template file used for generating certificates (default: certificate.html).
- `<OUTPUT_DIR>`: The path to the directory where the generated HTML files will be created (default: ./certificates).

## JSON File Format

The JSON file should have the following format:

```json
{
  "certificates": [
    {
      "name": "John Doe",
      "hours": "60 Hours",
      "course_title": "Cloud Native Bootcamp",
      "date_completed": "27/03/2023"
    },
  ],
  "signature": "Jane Smith"
}
```

- `certificates`: An array of certificate data objects.
- `signature`: The signature text that will be included in all certificates.

Each certificate data object should include:

- `name`: The name of the person receiving the certificate.
- `hours`: The number of hours completed.
- `course_title`: The title of the course.
- `date_completed`: The date the course was completed (formatted as "DD/MM/YYYY").

## Customizing the HTML Template

You can customize the appearance of the generated certificates by modifying the certificate.html file. Use the following placeholders to insert certificate data into the template:

- `{{.Name}}`: The name of the person receiving the certificate.
- `{{.Hours}}`: The number of hours completed.
- `{{.CourseTitle}}`: The title of the course.
- `{{.DateCompleted}}`: The date the course was completed (formatted as "DD/MM/YYYY").
- `{{.Signature}}`: The signature text.

## Example

1. Create a `certificates.json` file with the following content:

    ```json
    {
      "certificates": [
        {
          "name": "John Doe",
          "hours": "60 Hours",
          "course_title": "Cloud Native Bootcamp",
          "date_completed": "27/03/2023"
        },
        // ...
      ],
      "signature": "Jane Smith"
    }
    ```

1. Run the Certificate Generator:

    ```bash
    certgen --filename certificates.json --template certificate.html --output certificates
    ```

This command will generate one HTML certificate file in the certificates directory.

## FAQ

**Q: Can I generate certificates in other formats, such as PDF?**

A: Currently, the Certificate Generator only supports generating certificates in HTML format. However, you can consider using third-party libraries or APIs to convert the generated HTML certificates into other formats, such as PDF.

**Q: How can I customize the appearance of the generated certificates?**

A: You can customize the appearance by modifying the `certificate.html` file. Use the provided placeholders (e.g., `{{.Name}}`, `{{.Hours}}`, etc.) to insert certificate data into the template. You can also modify the CSS styles and HTML structure to achieve the desired look and feel.

**Q: I encountered an issue or need help. Where can I get support?**

A: Please open an issue on the GitHub repository, and I will try to help you as soon as I can.

## Changelog

- Version 1.0.0: Initial release.
  - Generate personalized HTML certificates based on JSON input and an HTML template.
  - Command-line flags for input/output configuration.
