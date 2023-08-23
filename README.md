
# Helper Go

`helper` is a project containing implementations of commonly used functions in Golang.

## Installation

To utilize the functions from this package, install it directly into your project with the following command:

```bash
go get github.com/tgkzz/helper
```
    
## Usage/Examples

Below is a simple example demonstrating how to use the functions:

```go
package main

import (
	"fmt"

	"github.com/tgkzz/helper"
)

func main() {
	str := "254"

	num := helper.Atoi(str)

	fmt.Println(helper.IsPrime(num))
}
```


## List of functions

Here is a list of currently existing commands

* Atoi 
* Itoa
* CharFrequency
* IsNumeric
* IsPrime
* ToUpper
* ToLower
* SplitWhiteSpace
## Running Tests

The package also includes its own set of tests. To run them, first clone the repository:

```bash
git clone https://github.com/tgkzz/helper.git
cd helper
```

Then run test 
```bash
go test -v
```


