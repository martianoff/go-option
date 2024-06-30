[![CircleCI](https://dl.circleci.com/status-badge/img/gh/martianoff/go-option/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/martianoff/go-option/tree/main)
[![codecov](https://codecov.io/gh/martianoff/go-option/graph/badge.svg?token=NQICPHBEUQ)](https://codecov.io/gh/martianoff/go-option)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/martianoff/go-option)](https://pkg.go.dev/github.com/martianoff/go-option)
[![Go Report Card](https://goreportcard.com/badge/github.com/martianoff/go-option)](https://goreportcard.com/report/github.com/martianoff/go-option)

# Go Option

The idea is based on Scala Option and Java Optional. The package allows to create optional values in Golang. This unlocks composition of option values that makes code much cleaner.

## Example

Let's check an example without using option,

```
type userModel struct {
	carModel *carModel
}

type carModel struct {
	name        string
	plateNumber *string
}

type User struct {
	car *Car
}

type Car struct {
	name           string
	hasPlateNumber bool
}

func getUser(user userModel) User {
	if user.carModel != nil {
		hasPlateNumber := false
		if user.carModel.plateNumber != nil {
			hasPlateNumber = true
		}
		return User{car: &Car{
			name:           user.carModel.name,
			hasPlateNumber: hasPlateNumber,
		}}
	}
	return User{car: nil}
}
```

Now the same code with option

``` 
type userModel struct {
	carModel Option[carModel]
}

type carModel struct {
	name        string
	plateNumber Option[string]
}

type User struct {
	car Option[Car]
}

type Car struct {
	name           string
	hasPlateNumber bool
}

func getUser(user userModel) User {
	return User{
		car: Map[carModel, Car](user.carModel, func(model carModel) Car {
		    // if model is Some[carModel], transform it to Some[Car]
			return Car{
				name:           model.name,
				hasPlateNumber: model.plateNumber.NonEmpty(),
			}
		}),
	}
}
```

## Functions

Set a non-empty value:
```
option.Some[int](10)
option.Some(10) // short style
```

Set an empty value:
```
v := option.None[int]() 
```

Set an empty value if v is nil, otherwise set non-empty value
```
v := option.NewOption[int](10)
v := option.NewOption(10) // short style
```

Convert pointer to an object to an option
```
v := option.NewOptionFromPointer[Car](nil) // None[Car]
v := option.NewOptionFromPointer[Car](&Car{}) // Some[Car]
v := option.NewOptionFromPointer(&Car{}) // Some[Car] // short style
```

Transform underlying value of option to non option value
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

Option composition. Transform underlying value of option to another option value
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
