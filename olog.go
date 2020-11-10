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
	items := reflect.ValueOf(rows)
	if items.Kind() != reflect.Slice {
		fmt.Errorf("olog.Table() only supports arrays. got %v instead",
			items.Kind())
	}
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
	println(o)
}

func Print(rows interface{}) {
	PrintWithStyle(rows, Normal)
}

func PrintSoft(rows interface{}) {
	PrintWithStyle(rows, Soft)
}

func PrintBold(rows interface{}) {
	PrintWithStyle(rows, Bold)
}

func PrintStrong(rows interface{}) {
	PrintWithStyle(rows, Strong)
}

func PrintVStrong(rows interface{}) {
	PrintWithStyle(rows, VStrong)
}

func PrintHStrong(rows interface{}) {
	PrintWithStyle(rows, HStrong)
}

func PrintClear(rows interface{}) {
	PrintWithStyle(rows, Clear)
}

func PrintMarkdown(rows interface{}) {
	PrintWithStyle(rows, Markdown)
}

func PrintBlock(rows interface{}) {
	PrintWithStyle(rows, Block)
}
