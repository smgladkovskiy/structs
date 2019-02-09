structs
-------

Various helpful golang structs package

## Zeroable types

Default value is zero value.

* date
* time

## Nullable types

Default value is null value.

* string
* bool
* int64
* float64
* time
* date

## Usage

Get package via `go get github.com/smgladkovskiy/structs` or via `dep ensure add github.com/smgladkovskiy/structs` 
if you use dep.

Import package in your go file:

```go
package main

import (
	"github.com/smgladkovskiy/structs"      // if need to set date or time format
	"github.com/smgladkovskiy/structs/null" // if nullables are used
	"github.com/smgladkovskiy/structs/zero" // if zeroables are used
)
```  

Use in code:

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/smgladkovskiy/structs"      // if need to set date or time format
	"github.com/smgladkovskiy/structs/null" // if nullables are used
	"github.com/smgladkovskiy/structs/zero" // if zeroables are used
)

func main() {
	structs.TimeFormat = func() string {
	    return time.RFC1123
	  }
	
	customTime, _ := null.NewTime(time.Now())
	
	if !customTime.Valid {
		fmt.Println("time is null")
		return
	}
	
	fmt.Printf("custom time is %s", customTime.Time)
}
```

## Helper functions and variables

### zero.NewXXX(v interface{}) *zero.XXX

Initiates new zero value of XXX type, using `Scan` method with passed value.

### null.NewXXX(v interface{}) *null.XXX

Initiates new nullable value of XXX type, using `Scan` method with passed value.

### null.NewXXXf(format string, a ...interface{}) *null.XXX

Initiates new `null.String` type value for passed format string and format variables. Available for strings.

### TimeFormat

There is an opportunity for `zero.Time` and `null.Time` types to set time format witch will be used with (Un)MarshallJSON methods.

To override default package format for time (`time.RFC3339`), there must be an `stucts.TimeFormat` function 
overriding in your app at the configuration or init level:


```go
package main

import (
	"time"
	
	"github.com/smgladkovskiy/structs"
)

func main() {
	structs.TimeFormat = func() string {
		return time.RFC1123
	}
}
```

### DateFormat

For `zero.Date` and `null.Date` types there is the same thing with format for (Un)MarshallJSON.

Default package date format (`YYYY-MM-DD`) must be overridden with `stucts.DateFormat` function:


```go
package main

import (
	"github.com/smgladkovskiy/structs"
)

func init() {
	structs.DateFormat = func() string {
		return "02.01.2006"
	}
}
```