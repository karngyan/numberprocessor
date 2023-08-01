# numberprocessor

This repository is a dummy Go project aimed at teaching debugging techniques in Go, specifically with Delve. It demonstrates how to debug common runtime errors such as "nil pointer dereference" and how to set conditions on breakpoints.

The code consists of a simple Go package called numberprocessor, which includes a function that processes a slice of numbers.

### Usage

You can use the package like this:

```go
package main

import (
	"fmt"
	"github.com/karngyan/numberprocessor"
)

func main() {
	data := make([]*numberprocessor.Number, 5)
    data[0] = &numberprocessor.Number{Value: 1}
    data[1] = &numberprocessor.Number{Value: 2}
    // This code deliberately introduces a 'nil pointer dereference' error

    result := numberprocessor.ProcessData(data)
    fmt.Println(result) 
}
```

### Debugging

We use Delve for debugging. To debug tests, use:

```bash
dlv test
```

To set a conditional breakpoint in the ProcessData function:

```bash
(dlv) break numberprocessor/numberprocessor.go:9 if item == nil
```

This helps to identify and fix runtime errors in the code.