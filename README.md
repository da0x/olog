![Go](https://github.com/da0x/olog/workflows/Go/badge.svg) ![CodeQL](https://github.com/da0x/olog/workflows/CodeQL/badge.svg)
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
	olog.Print(data)
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
olog.Print(olog.Soft, data)
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
olog.Print(olog.Bold, data)
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
olog.Print(olog.Strong, data)
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
olog.Print(olog.VStrong, data)
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
olog.Print(olog.HStrong, data)
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
olog.Print(olog.Clear, data)
```
```                       
 Name         Age   Score 
                          
 John Smith   30    99.223
 Jane Smith   30    99.223
```                       
#### Markdown
```go
olog.Print(olog.Markdown, data)
```
```
| Name       | Age | Score  |
|:----------:|:---:|:------:|
| John Smith | 30  | 99.223 |
| Jane Smith | 30  | 99.223 |
```
#### Block
```go
olog.Print(olog.Block, data)
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
