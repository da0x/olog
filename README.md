![Go](https://github.com/da0x/olog/workflows/Go/badge.svg)
# olog
This is a box printing library for Go. It takes any array of structs and assembles an ascii box around the objects. This logger is not meant for large amount of data and instead should be used only on smaller arrays.
### Installation
To download the library, simply run:
```
$ go get github.com/da0x/golang/olog
```
To import olog in your code, use:
```
import "github.com/da0x/golang/olog"
```
### Example
Here is an example of how to use olog:
```go
type Data struct {
	Name  string
	Age   int
	Score float32
}

func main() {
	var data = []Data{
		Data{Name: "John Smith", Age: 30, Score: 100.0},
		Data{Name: "Jane Smith", Age: 30, Score: 100.0},
	}
	olog.Print(data)        // same as PrintWithStyle(data, olog.Normal)
}
```
The above example prints the following output:
```
┌────────────┬─────┬────────┐
│ Name       │ Age │ Score  │
├────────────┼─────┼────────┤
│ John Smith │ 30  │ 99.223 │
│ Jane Smith │ 30  │ 99.223 │
└────────────┴─────┴────────┘
```
## Styles
Here are the available styles
#### Soft
```go
olog.PrintSoft(data)  // same as PrintWithStyle(data, olog.Soft)
```
```
╭────────────┬─────┬────────╮
│ Name       │ Age │ Score  │
├────────────┼─────┼────────┤
│ John Smith │ 30  │ 99.223 │
│ Jane Smith │ 30  │ 99.223 │
╰────────────┴─────┴────────╯
```
#### Bold
```go
olog.PrintBold(data)  // same as PrintWithStyle(data, olog.Bold)
```
```
┏━━━━━━━━━━━━┳━━━━━┳━━━━━━━━┓
┃ Name       ┃ Age ┃ Score  ┃
┣━━━━━━━━━━━━╋━━━━━╋━━━━━━━━┫
┃ John Smith ┃ 30  ┃ 99.223 ┃
┃ Jane Smith ┃ 30  ┃ 99.223 ┃
┗━━━━━━━━━━━━┻━━━━━┻━━━━━━━━┛
```
#### Strong
```go
olog.PrintStrong(data)  // same as PrintWithStyle(data, olog.Strong)
```
```
╔════════════╦═════╦════════╗
║ Name       ║ Age ║ Score  ║
╠════════════╬═════╬════════╣
║ John Smith ║ 30  ║ 99.223 ║
║ Jane Smith ║ 30  ║ 99.223 ║
╚════════════╩═════╩════════╝
```
#### Vertical Strong 
```go
olog.PrintVStrong(data) // same as PrintWithStyle(data, olog.VStrong)
```
```
╓────────────╥─────╥────────╖
║ Name       ║ Age ║ Score  ║
╟────────────╫─────╫────────╢
║ John Smith ║ 30  ║ 99.223 ║
║ Jane Smith ║ 30  ║ 99.223 ║
╙────────────╨─────╨────────╜
```
#### Horizontal Strong
```go
olog.PrintHStrong(data) // same as PrintWithStyle(data, olog.HStrong)
```
```
╒════════════╤═════╤════════╕
│ Name       │ Age │ Score  │
╞════════════╪═════╪════════╡
│ John Smith │ 30  │ 99.223 │
│ Jane Smith │ 30  │ 99.223 │
╘════════════╧═════╧════════╛
```
#### Clear
```go
olog.PrintClear(data)    // same as PrintWithStyle(data, olog.Clear)
```
```                       
 Name         Age   Score 
                          
 John Smith   30    99.223
 Jane Smith   30    99.223
```                       
#### Markdown
```go
olog.PrintMarkdown(data) // same as PrintWithStyle(data, olog.Markdown)
```
```
| Name       | Age | Score  |
|:----------:|:---:|:------:|
| John Smith | 30  | 99.223 |
| Jane Smith | 30  | 99.223 |
```
#### Block
```go
olog.PrintBlock(data) // same as PrintWithStyle(data, olog.Block)
```
```
▛▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▜
▌ Name       ┃ Age ┃ Score  ▐
▌━━━━━━━━━━━━╋━━━━━╋━━━━━━━━▐
▌ John Smith ┃ 30  ┃ 99.223 ▐
▌ Jane Smith ┃ 30  ┃ 99.223 ▐
▙▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▟
```
### Maintainer
Daher Alfawares
### Todo
☐ Single Object |Name|Value|
☐ Identical Cell Merging
☐ Header & Footer
☐ CSV
☐ Text Alignment
