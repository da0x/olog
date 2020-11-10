# olog
Welcome to `olog`! This is a table printing library for Go. It takes any array of structs and assembles an ascii box around the objects. This logger is not meant for large amount of data and instead should be used only on smaller arrays.
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
```
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
	olog.PrintStrong(data)  // same as PrintWithStyle(data, olog.Strong)
	olog.PrintVStrong(data) // same as PrintWithStyle(data, olog.VStrong)
	olog.PrintHStrong(data) // same as PrintWithStyle(data, olog.HStrong)
}
```
The above example prints the following output:
```
┌──────────┬───┬─────┐
│Name      │Age│Score│
├──────────┼───┼─────┤
│John Smith│30 │100  │
│Jane Smith│30 │100  │
└──────────┴───┴─────┘
╔══════════╦═══╦═════╗
║Name      ║Age║Score║
╠══════════╬═══╬═════╣
║John Smith║30 ║100  ║
║Jane Smith║30 ║100  ║
╚══════════╩═══╩═════╝
╓──────────╥───╥─────╖
║Name      ║Age║Score║
╟──────────╫───╫─────╢
║John Smith║30 ║100  ║
║Jane Smith║30 ║100  ║
╙──────────╨───╨─────╜
╒══════════╤═══╤═════╕
│Name      │Age│Score│
╞══════════╪═══╪═════╡
│John Smith│30 │100  │
│Jane Smith│30 │100  │
╘══════════╧═══╧═════╛
```
### Maintainer
Daher Alfawares
### Todo
- Identical Cell Merging
- Header & Footer
- CSV
- Text Alignment
