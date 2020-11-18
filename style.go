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

// StyleRow represents one row of a style.
// b for begining.
// m for middle.
// i for intersection.
// e for end.
type StyleRow struct {
	B string
	M string
	I string
	E string
}

type StyleSep struct {
	B string
	M string
	U string
	I string
	L string
	E string
}

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
type Style struct {
	T StyleRow
	M StyleRow
	S StyleRow
	L StyleRow
	B StyleSep
}

// Normal is a thin bordered table.
var Normal = Style{
	T: StyleRow{B: "┌─", M: "─", I: "─┬─", E: "─┐"},
	M: StyleRow{B: "│ ", M: " ", I: " │ ", E: " │"},
	S: StyleRow{B: "├─", M: "─", I: "─┼─", E: "─┤"},
	L: StyleRow{B: "└─", M: "─", I: "─┴─", E: "─┘"},
	B: StyleSep{B: "├─", M: "─", U: "┴", I: "┼", L: "┬", E: "─┤"},
}

// Soft is a thin bordered table.
var Soft = Style{
	T: StyleRow{B: "╭─", M: "─", I: "─┬─", E: "─╮"},
	M: StyleRow{B: "│ ", M: " ", I: " │ ", E: " │"},
	S: StyleRow{B: "├─", M: "─", I: "─┼─", E: "─┤"},
	L: StyleRow{B: "╰─", M: "─", I: "─┴─", E: "─╯"},
	B: StyleSep{B: "╞═", M: "═", U: "╧", I: "╪", L: "╤", E: "═╡"},
}

// Bold is a thin bordered table.
var Bold = Style{
	T: StyleRow{B: "┏━", M: "━", I: "━┳━", E: "━┓"},
	M: StyleRow{B: "┃ ", M: " ", I: " ┃ ", E: " ┃"},
	S: StyleRow{B: "┣━", M: "━", I: "━╋━", E: "━┫"},
	L: StyleRow{B: "┗━", M: "━", I: "━┻━", E: "━┛"},
	B: StyleSep{B: "┣━", M: "━", U: "┻", I: "╋", L: "┳", E: "━┫"},
}

// Strong is a double lined table.
var Strong = Style{
	T: StyleRow{B: "╔═", M: "═", I: "═╦═", E: "═╗"},
	M: StyleRow{B: "║ ", M: " ", I: " ║ ", E: " ║"},
	S: StyleRow{B: "╠═", M: "═", I: "═╬═", E: "═╣"},
	L: StyleRow{B: "╚═", M: "═", I: "═╩═", E: "═╝"},
	B: StyleSep{B: "╠═", M: "═", U: "╩", I: "╬", L: "╦", E: "═╣"},
}

// StrongVertical is a slim border with a double lined vertical borders.
var VStrong = Style{
	T: StyleRow{B: "╓─", M: "─", I: "─╥─", E: "─╖"},
	M: StyleRow{B: "║ ", M: " ", I: " ║ ", E: " ║"},
	S: StyleRow{B: "╟─", M: "─", I: "─╫─", E: "─╢"},
	L: StyleRow{B: "╙─", M: "─", I: "─╨─", E: "─╜"},
	B: StyleSep{B: "╟─", M: "─", U: "╨", I: "╫", L: "╥", E: "─╢"},
}

// StrongHorizantal is a slim border with a double lined horizantal borders.
var HStrong = Style{
	T: StyleRow{B: "╒═", M: "═", I: "═╤═", E: "═╕"},
	M: StyleRow{B: "│ ", M: " ", I: " │ ", E: " │"},
	S: StyleRow{B: "╞═", M: "═", I: "═╪═", E: "═╡"},
	L: StyleRow{B: "╘═", M: "═", I: "═╧═", E: "═╛"},
	B: StyleSep{B: "╞═", M: "═", U: "╧", I: "╪", L: "╤", E: "═╡"},
}

// Clear is a thin bordered table.
var Clear = Style{
	T: StyleRow{B: " ", M: " ", I: "   ", E: ""},
	M: StyleRow{B: " ", M: " ", I: "   ", E: ""},
	S: StyleRow{B: " ", M: " ", I: "   ", E: ""},
	L: StyleRow{B: " ", M: " ", I: "   ", E: ""},
	B: StyleSep{B: "", M: "", U: "", I: "", L: "", E: ""},
}

// Markdown is a thin bordered table.
var Markdown = Style{
	T: StyleRow{B: "  ", M: " ", I: "   ", E: "  "},
	M: StyleRow{B: "| ", M: " ", I: " | ", E: " |"},
	S: StyleRow{B: "|:", M: "-", I: ":|:", E: ":|"},
	L: StyleRow{B: "  ", M: " ", I: "   ", E: "  "},
	B: StyleSep{B: "", M: "", U: "", I: "", L: "", E: ""},
}

// Block is a slim border with a double lined horizantal borders.
var Block = Style{
	T: StyleRow{B: "▛▀", M: "▀", I: "▀▀▀", E: "▀▜"},
	M: StyleRow{B: "▌ ", M: " ", I: " ┃ ", E: " ▐"},
	S: StyleRow{B: "▌━", M: "━", I: "━╋━", E: "━▐"},
	L: StyleRow{B: "▙▄", M: "▄", I: "▄▄▄", E: "▄▟"},
	B: StyleSep{B: "▌■", M: "■", U: "■", I: "■", L: "■", E: "■▐"},
}

func row(s StyleRow, lengths []int) string {
	o := s.B
	for section, length := range lengths {
		for i := 0; i < length; i++ {
			o += s.M
		}
		if section != len(lengths)-1 {
			o += s.I
		}
	}
	o += s.E
	return o
}

func repeat(str string, n int) string {
	o := ""
	for i := 0; i < n; i++ {
		o += str
	}
	return o
}

func (s *Style) edge(width int, upperLengths, lowerLengths []int) string {
	var o []int
	for i := 0; i < width; i++ {
		o = append(o, 0)
	}
	var sum int
	for i, u := range upperLengths {
		if i == len(upperLengths)-1 {
			break
		}
		sum += u + 1
		o[sum] = 1
		sum += 2
	}
	sum = 0
	for i, v := range lowerLengths {
		if i == len(lowerLengths)-1 {
			break
		}
		sum += v + 1
		if o[sum] == 1 {
			o[sum] = 2
		} else {
			o[sum] = 3
		}
		sum += 2
	}
	out := s.B.B
	for _, c := range o {
		switch c {
		case 0:
			out += s.B.M
		case 1:
			out += s.B.U
		case 2:
			out += s.B.I
		case 3:
			out += s.B.L
		}
	}
	out += s.B.E
	return out
}

func (s *Style) seperatorWidth() int {
	return 3
	return len(s.M.I)
}

func (s *Style) topRow(lengths []int) string {
	return row(s.T, lengths)
}

func (s *Style) seperatorRow(lengths []int) string {
	return row(s.S, lengths)
}

func (s *Style) bottomRow(lengths []int) string {
	return row(s.L, lengths)
}

func (s *Style) middleRow(words []string, lengths []int) string {
	o := s.M.B
	for section, length := range lengths {
		o += words[section]
		i := len(words[section])
		for i < length {
			o += s.M.M
			i++
		}
		if section != len(lengths)-1 {
			o += s.M.I
		}
	}
	o += s.M.E
	return o
}
