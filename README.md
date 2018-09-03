nulls
-----

[![pipeline status](https://gitlab.teamc.io/teamc.io/golang/nulls/badges/master/pipeline.svg)](https://gitlab.teamc.io/teamc.io/golang/nulls/commits/master) [![coverage report](https://gitlab.teamc.io/teamc.io/golang/nulls/badges/master/coverage.svg)](https://gitlab.teamc.io/teamc.io/golang/nulls/commits/master)

Структуры nullable типов

## Реализованные типы

* bool
* int64
* string
* time

## Хелпер-функии и переменные

### NewNullXXX(v interface{}) *NullXXX

Инициирует новую nullable переменную из переданного значения, реализуя `Scan` метод. Доступно для всех типов

### NewNullXXXf(v interface{}) *NullXXX

Инициирует `NullString` переменную из переданного значения. Доступно только для строк.

### TimeFormat

Для `NullTime` имеется возможность передать формат времени, в котором и из которого будет делаться (Un)MarshallJSON.

Чтобы переопределить default формат (`time.RFC3339`), нужно переопределить переменную `nulls.TimeFormat` в приложении на уровне конфигурации приложения:

```go
import "gitlab.teamc.io/teamc.io/golang/nulls.git"

func init() {
    nulls.TimeFormat = func() string {
    	return time.RFC1123
    }
}
```