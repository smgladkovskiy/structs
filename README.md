structs
-------

Various helpful golang structs package

## Zero types

* date
* time

## Nullable types

* bool
* date
* int64
* string
* time

## Helper functions and variables

### NewXXX(v interface{}) *XXX

Initiates new value of XXX type, using `Scan` method with passed value.

### NewNullXXX(v interface{}) *NullXXX

Initiates new nullable value of XXX type, using `Scan` method with passed value.

### NewNullXXXf(format string, a ...interface{}) *NullXXX

Initiates new `NullString` type value for passed format string and format variables. Available for strings.

### TimeFormat

There is an opportunity for `Time` and `NullTime` types to set time format witch will be used with (Un)MarshallJSON methods.

To override default package format for time (`time.RFC3339`), there must be an `stucts.TimeFormat` function 
overriding in your app at the configuration or init level:


```go
package main

import (
	"github.com/smgladkovskiy/structs"
	"time"
)

func main() {
	structs.TimeFormat = func() string {
		return time.RFC1123
	}
}
```

### DateFormat

For `Date` and `NullDate` types there is the same thing with format for (Un)MarshallJSON.

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