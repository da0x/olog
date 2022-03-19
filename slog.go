//
//   slog.go
//   olog
//
//   Copyright 2022 Daher Alfawares
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License
//   You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2
//
//   Unless required by applicable law or agreed to in writing,
//   distributed under the License is distributed on an "AS IS" BASIS
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied
//   See the License for the specific language governing permissions
//   limitations under the License

package olog

import (
	"reflect"
)

func (i info) sprintTop() string {
	return (i.style.topRow(i.lengths)) + "\n"
}

func (i info) sprint() string {
	var o string
	for _, s := range i.rows {
		if len(s) == 0 {
			o += i.style.seperatorRow(i.lengths) + "\n"
		} else {
			o += i.style.middleRow(s, i.lengths) + "\n"
		}
	}
	return o
}

func (i info) sprintEdge(width int, upperLengths, lowerLengths []int) string {
	return (i.style.edge(width, upperLengths, lowerLengths)) + "\n"
}

func (i info) sprintBottom() string {
	return (i.style.bottomRow(i.lengths)) + "\n"
}

func Sprint(input ...interface{}) string {
	var info []info
	var s Style = Normal
	for _, o := range input {
		if reflect.TypeOf(o).String() == "olog.Style" {
			s = reflect.ValueOf(o).Interface().(Style)
			continue
		}
		n := infoOfAnyType(o)
		n.style = s
		info = append(info, n)
	}
	var max int = 0
	for c := range info {
		info[c].analyze()
		max = iMax(max, info[c].width())
	}
	var out string
	for i, o := range info {
		o.adjust(max - o.width())
		if i == 0 {
			out += o.sprintTop()
		}
		if i > 0 && i < len(info) {
			out += o.sprintEdge(max, info[i-1].lengths, info[i].lengths)
		}
		out += o.sprint()
		if i == len(info)-1 {
			out += o.sprintBottom()
		}
	}
	return out
}
