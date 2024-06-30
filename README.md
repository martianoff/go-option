[![CircleCI](https://dl.circleci.com/status-badge/img/gh/martianoff/go-option/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/martianoff/go-option/tree/main)
[![codecov](https://codecov.io/gh/martianoff/go-option/graph/badge.svg?token=NQICPHBEUQ)](https://codecov.io/gh/martianoff/go-option)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/martianoff/go-option)](https://pkg.go.dev/github.com/martianoff/go-option)
[![Go Report Card](https://goreportcard.com/badge/github.com/martianoff/go-option)](https://goreportcard.com/report/github.com/martianoff/go-option)

# Go Option

The idea is based on Scala Option and Java Optional. The package allows to create optional values in Golang. 

## Features

- **Optional Values:** This module introduces the Option[T] type, which represents optional values. An Option[T] either contains a value of type T (represented as Some(T)) or no value at all (represented as None).
- **Safe Handling of Missing Values:** The Option[T] type encourages safer handling of missing values, avoiding null pointer dereferences. It can be checked for its state (whether it's Some or None) before usage, ensuring that no nil value is being accessed.
- **Functional Methods:** The module provides functional methods such as Map, FlatMap etc. that make operations on Option[T] type instances easy, clear, and less error-prone.
- **Pattern Matching:** The Match function allows for clean and efficient handling of Option[T] instances. Depending on whether the Option[T] is Some or None, corresponding function passed to Match get executed, which makes the code expressive and maintains safety.
- **Equality Check:** The Equal method provides an efficient way to compare two Option[T] instances, checking if they represent the same state and hold equal values.
- **Generics-Based:** With Generics introduced in Go, the module offers a powerful, type-safe alternative to previous ways of handling optional values.
- **Json serialization and deserialization:** Build-in json support

## Example

Let's check an example without using options

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

Now the same code with options

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

**Map**: Option transformation. Transform underlying value of option to another non option value
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

**FlatMap:** Option composition. Transform underlying value of option to another option value and flat them to an option
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

## Pattern matching

Pattern matching is the ability to check and manipulate the value inside the Option[T] type safely without having to manually check the existence of the value every time. It likens to 'pattern matching'

The Match function  takes three parameters: Option[T1], and two functions (func(T1) T2).feature found in traditional functional languages. This function behaves differently depending on whether the given Option[T1] is Some or None:
- If Option[T1] contains a value (or, in other words, it's a Some), the function applies the first provided function func(T1) T2 on this value, aiming to transform it into a new T2 type value.
- If Option[T1] is None, meaning it does not contain a value, the second provided function func(T1) T2 is applied. Typically, this function should handle the case when no value is present, e.g., by providing a default value, handling an error, etc.

``` 
func getCarPlate(plateNumber Option[string]) string {
	return Match(plateNumber, func(p string) string {
		return p
	}, func() string {
		return "N/A"
	})
}
```

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
