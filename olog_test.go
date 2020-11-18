//
//   olog_test.go
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
	"testing"
)

type Data struct {
	Name  string
	Age   int
	Score float32
}

// func TestPrint(t *testing.T) {
// 	var data = []Data{
// 		Data{Name: "John Smith", Age: 30, Score: 99.223},
// 		Data{Name: "Jane Smith", Age: 30, Score: 99.223},
// 	}
// 	Print(data)         // same as PrintWithStyle(data, olog.Normal)
// 	PrintSoft(data)     // same as PrintWithStyle(data, olog.Soft)
// 	PrintBold(data)     // same as PrintWithStyle(data, olog.Bold)
// 	PrintStrong(data)   // same as PrintWithStyle(data, olog.Strong)
// 	PrintVStrong(data)  // same as PrintWithStyle(data, olog.VStrong)
// 	PrintHStrong(data)  // same as PrintWithStyle(data, olog.HStrong)
// 	PrintClear(data)    // same as PrintWithStyle(data, olog.Clear)
// 	PrintMarkdown(data) // same as PrintWithStyle(data, olog.Markdown)
// 	PrintBlock(data)    // same as PrintWithStyle(data, olog.Block)
// }
//
// func TestPrintString(t *testing.T) {
// 	Print("Test Print string")
// 	Print([]string{"Test Print []string", "String 1", "String 2"})
// }
//
// func TestStruct(t *testing.T) {
// 	Print(Data{"Ahmed", 25, 9.32})
// }
//
// func TestPrintCSV(t *testing.T) {
// 	Print([][]string{{"Test Print [][]string", "Single", "Row"}})
// 	Print([][]string{
// 		{"Test Print [][]string", "Multiple", "Rows"},
// 		{"value 100", "value 010", "value 001"},
// 		{"value 200", "value 020", "value 002"},
// 		{"value 300", "value 030", "value 003"},
// 	})
// }
//
// func TestEmpty(t *testing.T) {
// 	Print([]Data{})
// 	Print("")
// 	Print([]string{""})
// 	Print([]string{"", ""})
// 	Print([][]string{{"", ""}, {"", ""}})
// }

func TestVariadic(t *testing.T) {
	Print("String Only")
	Print([]string{"[]string", "hello", "world"})
	Print(
		[][]string{
			{"Test Print [][]string", "Multiple", "Rows"},
			{"value 100", "value 010", "value 001"},
			{"value 200", "value 020", "value 002"},
			{"value 300", "value 030", "value 003"},
		})
	Print(Data{"Ahmed", 25, 9.32})
	Print([]Data{Data{Name: "John Smith", Age: 30, Score: 99.223}, Data{Name: "Jane Smith", Age: 30, Score: 99.223}})
	Print(
		"string",
		[]string{"[]string", "hello", "world"},
		[]string{"[]string", "hello", "world", "for real", "now"},
		[][]string{
			{"Test Print [][]string", "Multiple", "Rows"},
			{"value 100", "value 010", "value 001"},
			{"value 200", "value 020", "value 002"},
			{"value 300", "value 030", "value 003"},
		},
		Data{"Ahmed", 25, 9.32},
		[]Data{
			Data{Name: "John Smith", Age: 30, Score: 99.223},
			Data{Name: "Jane Smith", Age: 30, Score: 99.223},
		},
		[]string{"[]string", "hello", "world", "for real", "now"},
	)
}
