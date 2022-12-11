# gencert

This project was made for educationnal purposes by @Farischt

Simple certificates generation using a [gofpdf](https://github.com/jung-kurt/gofpdf) module to create PDFs.

## Using the CLI

First build the project:

`go build -o build`

Once you've built the project you can start using the CLI as following:

`./build -type your_extension[pdf || html] -filename your_file_name.csv`

## Commands

Every command is prefixed by the prefix `-type`.
Here a the differents types available:

- pdf: `-type pdf`
- html: `-type html`

Every command also needs a csv file containing the students data just as [this example](./students.csv). **_Please make sure to use the exact same synthax !!!_**

`-filename students.csv`

An example of a full command:

`./build -type pdf -filename students.csv`
