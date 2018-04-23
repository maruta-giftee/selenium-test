package parser

import (
	"os"

	"github.com/gocarina/gocsv"
)

type Input struct {
	URL   string `csv:"url"`
	SSNO  string `csv:"ssno"`
	Name  string `csv:"name"`
	Item  string `csv:"item"`
	Limit string `csv:"limit"`
}

func ParseCSV(file *os.File) ([]*Input, error) {
	inputs := []*Input{}
	if err := gocsv.UnmarshalFile(file, &inputs); err != nil {
		return nil, err
	}
	return inputs, nil
}
