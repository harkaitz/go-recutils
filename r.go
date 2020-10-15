/*
recfile -- GNU recutils'es recfiles parser on pure Go
Copyright (C) 2020 Sergey Matveev <stargrave@stargrave.org>

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

type Reader struct {
	scanner *bufio.Scanner
}

func NewReader(r io.Reader) *Reader {
	return &Reader{bufio.NewScanner(r)}
}

type Field struct {
	Name  string
	Value string
}

var KeyValRe = regexp.MustCompile(`([a-zA-Z%][a-zA-Z0-9_]*):\s*(.*)$`)

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
