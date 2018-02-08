# sf_time


#### Download and install

    go get github.com/codermarcel/sf_time


#### Create file `main.go`

```go
package main

import (
	"fmt"

	"github.com/codermarcel/sf_time"
)

func main() {
	t, err := sf_time.NewFromSalesforceFormat("2015-10-05T17:17:02.001Z")

	if err != nil {
		panic(err)
	}

	fmt.Println(t.ToUnix())        //Convert to Unix time
	fmt.Println(t.ToTime().Unix()) //Or use time.Time

	t2, err := sf_time.NewFromSalesforceFormat("2012-02-08T19:17:32+01:20") //different format

	if err != nil {
		panic(err)
	}

	fmt.Println(t2.ToUnix())        //Convert to Unix time
    fmt.Println(t2.ToTime().Unix()) //Or use time.Time

    //For more examples check out the test file
}
```
