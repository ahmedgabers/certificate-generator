# Certificate Generator

The Certificate Generator is a command-line tool written in Go that generates personalised HTML certificates based on a provided JSON file and an HTML template.

## Features

- Generates personalised HTML certificates.
- Reads certificate data from a JSON file.
- Customizable HTML template for certificates.
- Supports command-line flags for input/output configuration.

## Installation

1. Go to the [releases page](https://github.com/ahmedgabers/certificate-generator/releases) of this repository.

1. Download the appropriate binary for your operating system.

1. Make the binary executable, for example:

    ```bash
      chmod +x certgen-linux-amd64
    ```

1. Move the binary to a location in your PATH:

   ```bash
   sudo mv certgen-linux-amd64 /usr/local/bin/certgen
   ```

1. Verify that the binary is installed correctly:

    ```bash
    certgen --help
    ```

## Customizing the HTML Template

You can customise the appearance of the generated certificates by modifying the example `certificate.html` file.

>Note: Replace the `<svg></svg>` block in the `certificate.html` to change the logo.

Use the following placeholders to insert certificate data into the template:

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

2. Run the Certificate Generator:

    ```bash
    certgen --filename certificates.json --template certificate.html --output certificates
    ```

This command will generate one HTML certificate file in the certificates directory.

## FAQ

**Q: Can I generate certificates in other formats, such as PDF?**

A: Currently, the Certificate Generator only supports generating certificates in HTML format. However, you can consider using third-party libraries or APIs to convert the generated HTML certificates into other formats, such as PDF.

**Q: How can I customise the appearance of the generated certificates?**

A: You can customise the appearance by modifying the `certificate.html` file. Use the provided placeholders (e.g., `{{.Name}}`, `{{.Hours}}`, etc.) to insert certificate data into the template. You can also modify the CSS styles and HTML structure to achieve the desired look and feel.

**Q: I encountered an issue or need help. Where can I get support?**

A: Please open an issue on the GitHub repository, and I will try to help you as soon as I can.
