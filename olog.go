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
	"fmt"
	"reflect"
)

func PrintWithStyle(rows interface{}, s Style) {
	h := headers(rows)
	v := values(rows)
	lengths := lengths(h, v)

	var o string
	o += fmt.Sprintln(s.topRow(lengths))
	o += fmt.Sprintln(s.middleRow(h, lengths))
	o += fmt.Sprintln(s.seperatorRow(lengths))
	for _, row := range v {
		o += fmt.Sprintln(s.middleRow(row, lengths))
	}
	o += fmt.Sprintln(s.bottomRow(lengths))
	print(o)
}

func PrintSliceOfString(strings []string, s Style) {
	lengths := vlen(strings)
	var o string
	o += fmt.Sprintln(s.topRow(lengths))
	o += fmt.Sprintln(s.middleRow(strings, lengths))
	o += fmt.Sprintln(s.bottomRow(lengths))
	print(o)
}

func PrintStructs(i interface{}, s Style) {
	PrintWithStyle(i, s)
}

type StructField struct {
	Key   string
	Value string
}

func PrintStruct(object interface{}, s Style) {
	h := headersOfStruct(object)
	v := valuesOfItem(reflect.ValueOf(object))
	var structFields []StructField
	for i, header := range h {
		structFields = append(structFields, StructField{header, v[i]})
	}
	PrintAnyType(structFields, s)
}

func PrintSliceOfSliceOfString(csv [][]string, s Style) {
	if len(csv) == 1 {
		PrintSliceOfString(csv[0], s)
		return
	}
	lengths := lengthsOfCSV(csv)
	var o string
	o += fmt.Sprintln(s.topRow(lengths))
	for i, row := range csv {
		o += fmt.Sprintln(s.middleRow(row, lengths))
		if i == 0 {
			o += fmt.Sprintln(s.seperatorRow(lengths))
		}
	}
	o += fmt.Sprintln(s.bottomRow(lengths))
	print(o)
}

func PrintAnyType(object interface{}, s Style) {
	o := reflect.TypeOf(object)
	v := reflect.ValueOf(object)
	switch o.Kind() {
	case reflect.Struct:
		PrintStruct(object, s)
	case reflect.String:
		PrintSliceOfString([]string{v.String()}, s)
	case reflect.Slice:
		switch o.Elem().Kind() {
		case reflect.String:
			PrintSliceOfString(object.([]string), s)
		case reflect.Struct:
			PrintStructs(object, s)
		case reflect.Slice:
			switch o.Elem().Elem().Kind() {
			case reflect.String:
				PrintSliceOfSliceOfString(object.([][]string), s)
			default:
				panic("only [][]string is supported")
			}
		}
	default:
		panic("unsupported type")
	}
}

func Print(rows interface{})         { PrintAnyType(rows, Normal) }
func PrintSoft(rows interface{})     { PrintAnyType(rows, Soft) }
func PrintBold(rows interface{})     { PrintAnyType(rows, Bold) }
func PrintStrong(rows interface{})   { PrintAnyType(rows, Strong) }
func PrintVStrong(rows interface{})  { PrintAnyType(rows, VStrong) }
func PrintHStrong(rows interface{})  { PrintAnyType(rows, HStrong) }
func PrintClear(rows interface{})    { PrintAnyType(rows, Clear) }
func PrintMarkdown(rows interface{}) { PrintAnyType(rows, Markdown) }
func PrintBlock(rows interface{})    { PrintAnyType(rows, Block) }
