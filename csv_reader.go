package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CSVReader struct {
	fileName string
}

func newCSVReader(fileName string) *CSVReader {
	return &CSVReader{fileName: fileName}
}

func (c *CSVReader) Headers() ([]string, error) {
	f, err := os.Open(c.fileName)
	if err != nil {
		return nil, err
	}
	return csv.NewReader(f).Read()
}

func (c *CSVReader) Rows() ([][]string, error) {
	contents, err := c.read()
	if err != nil {
		return nil, err
	}
	if len(contents) <= 1 {
		return nil, fmt.Errorf("Not enough rows in CSV file")
	}
	return contents[1:], nil
}

func (c *CSVReader) Columns() ([][]string, error) {
	rows, err := c.Rows()
	if err != nil {
		return nil, err
	}
	xl := len(rows[0])
	yl := len(rows)
	columns := make([][]string, xl)
	for i := range columns {
		columns[i] = make([]string, yl)
	}
	for r := 0; r < xl; r++ {
		for c := 0; c < yl; c++ {
			columns[r][c] = rows[c][r]
		}
	}
	return columns, nil
}

func (c *CSVReader) read() ([][]string, error) {
	f, err := os.Open(c.fileName)
	if err != nil {
		return nil, err
	}
	return csv.NewReader(f).ReadAll()
}
