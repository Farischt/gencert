package cert

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ParseCsvFile(path string) ([]*Cert, error) {

	var certs []*Cert

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := 1
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), ",")
		cert, err := New(args[0], args[1], args[2])

		if err != nil {
			return nil, fmt.Errorf("%s - in line %d of file '%s'", err.Error(), line, path)
		}

		certs = append(certs, cert)
		line++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return certs, nil

}
