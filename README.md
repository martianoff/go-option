[![CircleCI](https://dl.circleci.com/status-badge/img/gh/martianoff/go-option/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/martianoff/go-option/tree/main)
[![codecov](https://codecov.io/gh/martianoff/go-option/graph/badge.svg?token=NQICPHBEUQ)](https://codecov.io/gh/martianoff/go-option)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/martianoff/go-option)](https://pkg.go.dev/github.com/martianoff/go-option)
[![Go Report Card](https://goreportcard.com/badge/github.com/martianoff/go-option)](https://goreportcard.com/report/github.com/martianoff/go-option)

# Go Option

The idea is based on Scala Option and Java Optional. The package allows to create optional values in Golang

## Functions

Set an non-empty value:
```
option.Some[int](10) 
```

Set an empty value:
```
v := option.None[int]() 
```

Set an empty value if v is nil, otherwise set non-empty value
```
v := option.NewOption[int](10)
```

Remap one option to another option
```
import "github.com/martianoff/go-option"

type Car struct {
    name        string
    plateNumber option.Option[string]
}

carOpt := option.Some[Car](
    Car {
        name:        "bmw"
    },
)

// get car name as option
carNameOpt := option.Map[Car, string](carOpt, func(c Car) string {
   return c.name
})
```

Option composition
```
import "github.com/martianoff/go-option"

type Car struct {
    name        string
    plateNumber option.Option[string]
}

type User struct {
    name string
    car  option.Option[Car]
}

u := User{
    name: "jake",
    car: option.Some[Car](
        Car{
            name:        "bmw",
            plateNumber: option.Some[string]("X723"),
        },
    ),
}

// get car plate as option
carPlateOpt := option.FlatMap[Car, string](u.car, func(c Car) option.Option[string] {
    return c.plateNumber
})
```

## Methods of Option

| Method              |                   Description                   |
|---------------------|:-----------------------------------------------:|
| Get()               |         gets underlying value (unsafe*)         |
| GetOrElse(fallback) | gets underlying value or returns fallback value |
| OrElse(fallbackOpt) |   returns fallback option if option is empty    |
| Empty()             |            checks if value is empty             |
| NonEmpty()          |             checks if value is set              |
| String()            |              string representation              |
| Map()               |           transform to another option           |
| FlatMap()           |               compose two options               |
`* - empty value will panic`

## Json serialization and deserialization

Build-in serialization and deserialization

## Custom equality

Custom equality powered by `github.com/google/go-cmp/cmp`

```
cmp.Equal(Some(11), Some(11)) // returns true
```

---
## Testing

To run all tests in this module:

```
go test ./...
```
