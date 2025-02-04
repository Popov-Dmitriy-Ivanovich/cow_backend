package load

import (
	"encoding/csv"
	"errors"
	"os"
	"time"
)

var separator = []rune{',', ';', '|'}

type autoCsv struct {
	Reader *csv.Reader
	header []string
}

func GetCsvReader(file *os.File) (reader *csv.Reader, header []string, err error) {
	for _, rn := range separator {
		if _, err := file.Seek(0, 0); err != nil {
			return nil, nil, err
		}
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

var timeFormats = []string{
	time.DateOnly,
	"02.01.2006",
	"02.01.06",
	"02/01/2006",
	"02/01/06",
	"02-01-2006",
	"02-01-06",
}

func ParseTime(timeStr string) (time.Time, error) {
	for _, format := range timeFormats {
		t, err := time.Parse(format, timeStr)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, errors.New("дата не соответсвует ни одному из доступных форматов")
}
