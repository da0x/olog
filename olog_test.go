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

func TestTable(t *testing.T) {
	var data = []Data{
		Data{Name: "Daher", Age: 39, Score: 99.223},
		Data{Name: "Justin", Age: 25, Score: 99.223},
	}
	Print(data)
}
