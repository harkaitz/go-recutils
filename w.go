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
	"io"
	"strings"
)

type Writer struct {
	w io.StringWriter
}

func NewWriter(w io.StringWriter) *Writer {
	return &Writer{w}
}

func (w *Writer) RecordStart() (written int, err error) {
	return w.w.WriteString("\n")
}

func (w *Writer) WriteFields(fs ...Field) (written int, err error) {
	var n int
	for _, f := range fs {
		n, err = w.w.WriteString(f.Name + ": " + strings.TrimLeft(f.Value, " ") + "\n")
		written += n
		if err != nil {
			return
		}
	}
	return
}

func (w *Writer) WriteFieldMultiline(name string, lines []string) (written int, err error) {
	var n int
	n, err = w.w.WriteString(name + ": " + strings.TrimLeft(lines[0], " ") + "\n")
	written += n
	if err != nil {
		return
	}
	for _, l := range lines[1:] {
		n, err = w.w.WriteString("+ " + l + "\n")
		written += n
		if err != nil {
			return
		}
	}
	return
}
