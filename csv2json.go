package csv2json

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
)

type csv2json struct {
	header  int
	fromCol int
	toCol   int
	fromRow int
	toRow   int
}

func index2name(index int) string {
	b := bytes.Buffer{}
	for {
		b.WriteByte(byte('A' + (index % 26)))
		index /= 26
		if index == 0 {
			break
		}
	}
	str := b.String()

	b.Reset()
	for i := 0; i < len(str); i++ {
		b.WriteByte(byte(str[len(str)-1-i]))
	}

	return b.String()
}

func (c2j *csv2json) conv(c io.Reader, j io.Writer) error {
	var header []string
	var result []map[string]string

	r := csv.NewReader(c)

	for i := 0; ; i++ {
		record, err := r.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return err
		}

		item := make(map[string]string)

		if c2j.header < 0 {
			if i < c2j.fromRow || i > c2j.toRow {
				continue
			}
			for i, v := range record {
				if i >= c2j.fromCol && i <= c2j.toCol {
					item[index2name(i)] = v
				}
			}
			result = append(result, item)
			continue
		}

		if i < c2j.header {
			continue
		}
		if i == c2j.header {
			header = record
			continue
		}
		if i < c2j.fromRow || i > c2j.toRow {
			continue
		}
		for i, v := range record {
			if i >= c2j.fromCol && i <= c2j.toCol {
				item[header[i]] = v
			}
		}
		result = append(result, item)
	}

	return json.NewEncoder(j).Encode(result)
}

type Option func(*csv2json)

func WithHeader(row int) Option {
	return func(c2j *csv2json) {
		c2j.header = row
	}
}

func WithCol(from, to int) Option {
	return func(c2j *csv2json) {
		c2j.fromCol = from
		c2j.toCol = to
	}
}

func WithRow(from, to int) Option {
	return func(c2j *csv2json) {
		c2j.fromRow = from
		c2j.toRow = to
	}
}

func Conv(c io.Reader, j io.Writer, ops ...Option) error {
	c2j := &csv2json{
		header:  0,
		fromCol: 0,
		toCol:   0x7fffffff,
		fromRow: 0,
		toRow:   0x7fffffff,
	}
	for _, f := range ops {
		f(c2j)
	}

	if c2j.header > c2j.fromCol {
		return errors.New("header must gt from col")
	}

	return c2j.conv(c, j)
}
