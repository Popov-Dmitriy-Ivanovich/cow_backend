package load

import (
	"encoding/csv"
	"errors"
	"os"
)

var separator = []rune{',', ';', '|'}

type autoCsv struct {
	Reader *csv.Reader
	header []string
}

func GetCsvReader(file *os.File) (reader *csv.Reader, header []string, err error) {
	for _, rn := range separator {

		reader := csv.NewReader(file)
		reader.Comma = rn
		header, err = reader.Read()
		if err != nil {
			return nil, nil, err
		}

		if len(header) > 1 {
			return reader, header, err
		}

	}
	return nil, nil, errors.New("не удалось найти подходящий разделитель")
}
