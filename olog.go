//
//   olog.go
//   olog
//
//   Copyright 2020 Daher Alfawares
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

type info struct {
	rows    [][]string
	lengths []int
	style   Style
}

func infoOfSliceOfStruct(s interface{}) info {
	var i info
	i.rows = append(i.rows, headers(s))
	i.rows = append(i.rows, []string{})
	i.rows = merge(i.rows, values(s))
	return i
}

func infoOfStruct(object interface{}) info {
	h := headersOfStruct(object)
	v := valuesOfItem(reflect.ValueOf(object))
	var i info
	for index, header := range h {
		i.rows = append(i.rows, []string{header, v[index]})
	}
	return i
}

func infoOfAnyType(object interface{}) info {
	o := reflect.TypeOf(object)
	v := reflect.ValueOf(object)
	switch o.Kind() {
	case reflect.String:
		return info{rows: [][]string{{v.String()}}}
	case reflect.Struct:
		return infoOfStruct(object)
	case reflect.Slice:
		switch o.Elem().Kind() {
		case reflect.String:
			return info{rows: [][]string{object.([]string)}}
		case reflect.Struct:
			return infoOfSliceOfStruct(object)
		case reflect.Slice:
			switch o.Elem().Elem().Kind() {
			case reflect.String:
				return info{rows: object.([][]string)}
			default:
				panic("only [][]string is supported")
			}
		}
	default:
		panic("unsupported type")
	}
	return info{}
}

func vlengths(rows [][]string) [][]int {
	var o [][]int
	return o
}

func (i *info) analyze() {
	if len(i.rows) == 0 {
		return
	}
	for _, row := range i.rows {
		if len(row) == 0 {
			continue
		}
		if len(i.lengths) == 0 {
			i.lengths = vlen(row)
		} else {
			i.lengths = vMax(i.lengths, vlen(row))
		}
	}
}

func (i *info) width() int {
	return sigma(i.lengths) + (len(i.lengths)-1)*i.style.seperatorWidth()
}

func (i *info) adjust(w int) {
	var c int
	var l int = len(i.lengths)
	if l == 0 {
		return
	}
	for w > 0 {
		i.lengths[c%l] += 1
		w -= 1
		c += 1
	}
}

func (i info) printTop() {
	println(i.style.topRow(i.lengths))
}

func (i info) print() {
	for _, s := range i.rows {
		if len(s) == 0 {
			println(i.style.seperatorRow(i.lengths))
		} else {
			println(i.style.middleRow(s, i.lengths))
		}
	}
}

func (i info) printEdge(width int, upperLengths, lowerLengths []int) {
	println(i.style.edge(width, upperLengths, lowerLengths))
}

func (i info) printBottom() {
	println(i.style.bottomRow(i.lengths))
}

func Print(input ...interface{}) {
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
	for i, o := range info {
		o.adjust(max - o.width())
		if i == 0 {
			o.printTop()
		}
		if i > 0 && i < len(info) {
			o.printEdge(max, info[i-1].lengths, info[i].lengths)
		}
		o.print()
		if i == len(info)-1 {
			o.printBottom()
		}
	}
}
