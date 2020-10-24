//
//   info.go
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

import "reflect"
import "fmt"

func headers(datasets interface{}) []string {
	var o = []string{}
	items := reflect.ValueOf(datasets)
	if items.Kind() == reflect.Slice {
		if items.Len() <= 0 {
			return []string{}
		}
		i := 0 // we just need one of them.
		item := items.Index(i)
		if item.Kind() == reflect.Struct {
			v := reflect.Indirect(item)
			for j := 0; j < v.NumField(); j++ {
				o = append(o, fmt.Sprintf("%v", v.Type().Field(j).Name))

			}
		}
	}
	return o
}

func values(datasets interface{}) [][]string {
	var o = [][]string{}
	items := reflect.ValueOf(datasets)
	if items.Kind() != reflect.Slice {
		return [][]string{}
	}
	for i := 0; i < items.Len(); i++ {
		var r = []string{}
		item := items.Index(i)
		if item.Kind() == reflect.Struct {
			v := reflect.Indirect(item)
			for j := 0; j < v.NumField(); j++ {
				r = append(r, fmt.Sprintf("%v", v.Field(j).Interface()))
			}
		}
		o = append(o, r)
	}
	return o
}

func iMax(v1 int, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

func vMax(v1 []int, v2 []int) []int {
	var o = []int{}
	for i, vv1 := range v1 {
		vv2 := v2[i]
		o = append(o, iMax(vv1, vv2))
	}
	return o
}

func vlen(in []string) []int {
	var o = []int{}
	for _, v := range in {
		o = append(o, len(v))
	}
	return o
}

func lengths(headers []string, rows [][]string) []int {
	o := vlen(headers)
	for _, row := range rows {
		o = vMax(o, vlen(row))
	}
	return o
}
