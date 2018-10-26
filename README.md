nulls
-----

Структуры разных полезных типов

## Zero типы данных

* date
* time

## Nullable типы

* bool
* date
* int64
* string
* time

## Хелпер-функии и переменные

### NewXXX(v interface{}) *XXX

Инициирует новую переменную соответствующего типа, реализуя `Scan` метод на переданное значение.

### NewNullXXX(v interface{}) *NullXXX

Инициирует новую nullable переменную из переданного значения, реализуя `Scan` метод на переданное значение.

### NewNullXXXf(v interface{}) *NullXXX

Инициирует `NullString` переменную из переданного значения. Доступно только для строк.

### TimeFormat

Для `Time`, `NullTime` имеется возможность передать формат времени, в котором и из которого будет делаться (Un)MarshallJSON.

Чтобы переопределить default формат (`time.RFC3339`), нужно переопределить переменную `stucts.TimeFormat` в приложении на уровне конфигурации приложения:


```go
package main

import (
	"gitlab.teamc.io/teamc.io/golang/structs"
	"time"
)

func main() {
	structs.TimeFormat = func() string {
		return time.RFC1123
	}
}
```

### DateFormat

Для `Date`, `NullDate` имеется возможность передать формат даты, в котором и из которого будет делаться (Un)MarshallJSON.

Чтобы переопределить default формат (`YYYY-MM-DD`), нужно переопределить переменную `stucts.DateFormat` в приложении на уровне конфигурации приложения:


```go
package main

import (
	"gitlab.teamc.io/teamc.io/golang/structs"
)

func init() {
	structs.DateFormat = func() string {
		return "02.01.2006"
	}
}
```