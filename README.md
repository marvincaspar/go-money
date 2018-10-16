# Go Money

[![Build Status](https://travis-ci.org/mc388/go-money.svg?branch=master)](https://travis-ci.org/mc388/go-money)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/mc388/go-money/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/mc388/go-money/?branch=master)
[![Code Coverage](https://scrutinizer-ci.com/g/mc388/go-money/badges/coverage.png?b=master)](https://scrutinizer-ci.com/g/mc388/go-money/?branch=master)
[![GitHub](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/mc388/go-money)


Go libraray to make working with money safer and easier!

```go
package main

import "github.com/mc388/go-money"

func main() {
    oneEuro := money.New(100, money.EUR())
    twoEuro, err := oneEuro.Add(oneEuro)

    if err != nil {
        log.Fatal(err)
    }

    parties, err := twoEuro.Allocate(1, 1, 1)

    if err != nil {
        log.Fatal(err)
    }

    parties[0].Display() // €0.67
    parties[1].Display() // €0.67
    parties[2].Display() // €0.66
}
```

## Install

Get the package:
```sh
go get github.com/mc388/go-money
```

## Usage

### Initialzation

Initialze Money by using the smalles unit (e.g. 100 represents 1€) and one of the given currency factories.

```go
euro := money.New(100, money.EUR())
```

### Calculations

The following calculations are available:
* Add
* Subtract
* Multiply
* Allocate

#### Add

Use `Add(money)` to perform additions.


```go
oneEuro := money.New(100, money.EUR())
twoEuro := money.New(200, money.EUR())
threeEuro, err := oneEuro.Add(twoEuro) // €3, nil
}
```

#### Subtract

Use `Subtract(money)` to perform substractions.


```go
threeEuro := money.New(300, money.EUR())
twoEuro := money.New(200, money.EUR())
oneEuro, err := threeEuro.Subtract(twoEuro) // €1, nil
}
```

#### Multiply

Use `Multiply(multiplicator)` to perform multiplications.


```go
twoEuro := money.New(200, money.EUR())
eightEuro, err := twoEuro.Multiply(4) // €8, nil
}
```

#### Allocate

Use `Allocate(...ratioX)` to allocate money by a given rations without loosing any cent.


```go
oneEuro := money.New(100, money.EUR())
parties, err := oneEuro.Allocate(50, 50)

if err != nil {
    log.Fatal(err)
}

parties[0].Display() // €0.50
parties[1].Display() // €0.50
}
```


```go
oneEuro := money.New(100, money.EUR())
parties, err := oneEuro.Allocate(1, 1, 1)

if err != nil {
    log.Fatal(err)
}

parties[0].Display() // €0.34
parties[1].Display() // €0.33
parties[2].Display() // €0.33
}
```

### Comparison

The following comparion methods are available:
* Equals
* GreaterThan
* GreaterThanOrEqual
* LessThan
* LessThanOrEqual

Comparison can only performe between the same currency.

```go
oneEuro := money.New(100, money.EUR())
twoEuro := money.New(200, money.EUR())
oneDollar := money.New(100, money.USD())

result1, err := oneEuro.GreaterThan(twoEuro) // false, nil
result2, err := oneEuro.LessThan(twoEuro) // true, nil
result3, err := oneEuro.Equals(oneDollar) // nil, error: Currency don't match
}
```

## Tests

Run the unit tests with `go test -v -covermode=count -coverprofile=coverage.out`.

To view the coverage report use `go tool cover -func=coverage.out`