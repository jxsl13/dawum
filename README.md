# dawum - a dawum.de Go client

This is a small utility for fetching German election poll data from the https://dawum.de api.

Reference: https://dawum.de/API/

# import dependency

```shell
go get github.com/jxsl13/dawum@latest
```

# example usage

```go
package main

import (
    "fmt"
    "github.com/jxsl13/dawum"
)

func main() {
    ExampleGetLastUpdate()
    ExampleGetData()
}

func ExampleGetData() {
	data, err := dawum.GetData(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
	// {"Database":{"License":{"Name":"ODC Open Database License","Shortcut":"ODC-ODbL","Link":"https........
}

func ExampleGetLastUpdate() {
	datetime, err := dawum.GetLastUpdate(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(datetime)
	// 2023-09-29T09:41:59+02:00
}
```
