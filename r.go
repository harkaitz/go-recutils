/*
recfile -- GNU recutils'es recfiles parser on pure Go
Copyright (C) 2020-2022 Sergey Matveev <stargrave@stargrave.org>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, version 3 of the License.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package recfile

import (
	"bufio"
	"errors"
	"io"
	"regexp"
	"strings"
)

var KeyValRe = regexp.MustCompile(`([a-zA-Z%][a-zA-Z0-9_]*):\s*(.*)$`)

type Reader struct {
	scanner *bufio.Scanner
}

// Create Reader for iterating through the records. It uses
// bufio.Scanner, so can read more than currently parsed records take.
func NewReader(r io.Reader) *Reader {
	return &Reader{bufio.NewScanner(r)}
}

// Get next record. Each record is just a collection of fields. io.EOF
// is returned if there is nothing to read more.
func (r *Reader) Next() ([]Field, error) {
	fields := make([]Field, 0, 1)
	var text string
	var name string
	var line string
	lines := make([]string, 0)
	continuation := false
	var err error
	for {
		if !r.scanner.Scan() {
			if err := r.scanner.Err(); err != nil {
				return fields, err
			}
			err = io.EOF
			break
		}
		text = r.scanner.Text()

		if continuation {
			if text[len(text)-1] == '\\' {
				lines = append(lines, text[:len(text)-1])
			} else {
				lines = append(lines, text)
				fields = append(fields, Field{name, strings.Join(lines, "")})
				lines = make([]string, 0)
				continuation = false
			}
			continue
		}

		if len(text) > 0 && text[0] == '#' {
			continue
		}

		if len(text) > 0 && text[0] == '+' {
			lines = append(lines, "\n")
			if len(text) > 1 {
				if text[1] == ' ' {
					lines = append(lines, text[2:])
				} else {
					lines = append(lines, text[1:])
				}
			}
			continue
		}

		if len(lines) > 0 {
			fields = append(fields, Field{name, strings.Join(lines, "")})
			lines = make([]string, 0)
		}

		if text == "" {
			break
		}

		matches := KeyValRe.FindStringSubmatch(text)
		if len(matches) == 0 {
			return fields, errors.New("invalid field format")
		}
		name = matches[1]
		line = matches[2]

		if len(line) > 0 && line[len(line)-1] == '\\' {
			continuation = true
			lines = append(lines, line[:len(line)-1])
		} else {
			lines = append(lines, line)
		}
	}
	if continuation {
		return fields, errors.New("left continuation")
	}
	if len(lines) > 0 {
		fields = append(fields, Field{name, strings.Join(lines, "")})
	}
	if len(fields) == 0 {
		if err == nil {
			return r.Next()
		}
		return fields, err
	}
	return fields, nil
}

// Same as Next(), but with unique keys and last value.
func (r *Reader) NextMap() (map[string]string, error) {
	fields, err := r.Next()
	if err != nil {
		return nil, err
	}
	m := make(map[string]string, len(fields))
	for _, f := range fields {
		m[f.Name] = f.Value
	}
	return m, nil
}

// Same as Next(), but with unique keys and slices of values.
func (r *Reader) NextMapWithSlice() (map[string][]string, error) {
	fields, err := r.Next()
	if err != nil {
		return nil, err
	}
	m := make(map[string][]string)
	for _, f := range fields {
		m[f.Name] = append(m[f.Name], f.Value)
	}
	return m, nil
}
