//
//   style.go
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

// Style represents the table structure in 4 strings.
// 't' for top.
// 'm' for middle.
// 's' for seperator.
// 'b' for bottom.
// ┌─┬┐  ╔═╦╗  ╓─╥╖  ╒═╤╕
// │ ││  ║ ║║  ║ ║║  │ ││
// ├─┼┤  ╠═╬╣  ╟─╫╢  ╞═╪╡
// └─┴┘  ╚═╩╝  ╙─╨╜  ╘═╧╛
// ┌───────────────────┐
// │  ╔═══╗ Some Text  │▒
// │  ╚═╦═╝ in the box │▒
// ╞═╤══╩══╤═══════════╡▒
// │ ├──┬──┤           │▒
// │ └──┴──┘           │▒
// └───────────────────┘▒
//  ▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒
//
// https://en.wikipedia.org/wiki/Box-drawing_character
type Style struct{ t, m, s, b StyleRow }

// Style row represents one row of a style.
// b for begining.
// m for middle.
// i for intersection.
// e for end.
type StyleRow struct{ b, m, i, e string }

// Normal is a thin bordered table.
var Normal = Style{
	t: StyleRow{b: "┌─", m: "─", i: "─┬─", e: "─┐"},
	m: StyleRow{b: "│ ", m: " ", i: " │ ", e: " │"},
	s: StyleRow{b: "├─", m: "─", i: "─┼─", e: "─┤"},
	b: StyleRow{b: "└─", m: "─", i: "─┴─", e: "─┘"},
}

// Soft is a thin bordered table.
var Soft = Style{
	t: StyleRow{b: "╭─", m: "─", i: "─┬─", e: "─╮"},
	m: StyleRow{b: "│ ", m: " ", i: " │ ", e: " │"},
	s: StyleRow{b: "├─", m: "─", i: "─┼─", e: "─┤"},
	b: StyleRow{b: "╰─", m: "─", i: "─┴─", e: "─╯"},
}

// Bold is a thin bordered table.
var Bold = Style{
	t: StyleRow{b: "┏━", m: "━", i: "━┳━", e: "━┓"},
	m: StyleRow{b: "┃ ", m: " ", i: " ┃ ", e: " ┃"},
	s: StyleRow{b: "┣━", m: "━", i: "━╋━", e: "━┫"},
	b: StyleRow{b: "┗━", m: "━", i: "━┻━", e: "━┛"},
}

// Strong is a double lined table.
var Strong = Style{
	t: StyleRow{b: "╔═", m: "═", i: "═╦═", e: "═╗"},
	m: StyleRow{b: "║ ", m: " ", i: " ║ ", e: " ║"},
	s: StyleRow{b: "╠═", m: "═", i: "═╬═", e: "═╣"},
	b: StyleRow{b: "╚═", m: "═", i: "═╩═", e: "═╝"},
}

// StrongVertical is a slim border with a double lined vertical borders.
var VStrong = Style{
	t: StyleRow{b: "╓─", m: "─", i: "─╥─", e: "─╖"},
	m: StyleRow{b: "║ ", m: " ", i: " ║ ", e: " ║"},
	s: StyleRow{b: "╟─", m: "─", i: "─╫─", e: "─╢"},
	b: StyleRow{b: "╙─", m: "─", i: "─╨─", e: "─╜"},
}

// StrongHorizantal is a slim border with a double lined horizantal borders.
var HStrong = Style{
	t: StyleRow{b: "╒═", m: "═", i: "═╤═", e: "═╕"},
	m: StyleRow{b: "│ ", m: " ", i: " │ ", e: " │"},
	s: StyleRow{b: "╞═", m: "═", i: "═╪═", e: "═╡"},
	b: StyleRow{b: "╘═", m: "═", i: "═╧═", e: "═╛"},
}

// Clear is a thin bordered table.
var Clear = Style{
	t: StyleRow{b: " ", m: " ", i: "   ", e: ""},
	m: StyleRow{b: " ", m: " ", i: "   ", e: ""},
	s: StyleRow{b: " ", m: " ", i: "   ", e: ""},
	b: StyleRow{b: " ", m: " ", i: "   ", e: ""},
}

// Markdown is a thin bordered table.
var Markdown = Style{
	t: StyleRow{b: "  ", m: " ", i: "   ", e: "  "},
	m: StyleRow{b: "| ", m: " ", i: " | ", e: " |"},
	s: StyleRow{b: "|:", m: "-", i: ":|:", e: ":|"},
	b: StyleRow{b: "  ", m: " ", i: "   ", e: "  "},
}

// Block is a slim border with a double lined horizantal borders.
var Block = Style{
	t: StyleRow{b: "▛▀", m: "▀", i: "▀▀▀", e: "▀▜"},
	m: StyleRow{b: "▌ ", m: " ", i: " ┃ ", e: " ▐"},
	s: StyleRow{b: "▌━", m: "━", i: "━╋━", e: "━▐"},
	b: StyleRow{b: "▙▄", m: "▄", i: "▄▄▄", e: "▄▟"},
}

func row(s StyleRow, lengths []int) string {
	o := s.b
	for section, length := range lengths {
		for i := 0; i < length; i++ {
			o += s.m
		}
		if section != len(lengths)-1 {
			o += s.i
		}
	}
	o += s.e
	return o
}

func (s *Style) topRow(lengths []int) string {
	return row(s.t, lengths)
}

func (s *Style) seperatorRow(lengths []int) string {
	return row(s.s, lengths)
}

func (s *Style) bottomRow(lengths []int) string {
	return row(s.b, lengths)
}

func (s *Style) middleRow(words []string, lengths []int) string {
	o := s.m.b
	for section, length := range lengths {
		o += words[section]
		i := len(words[section])
		for i < length {
			o += s.m.m
			i++
		}
		if section != len(lengths)-1 {
			o += s.m.i
		}
	}
	o += s.m.e
	return o
}
