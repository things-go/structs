# mapstruct

Go library for encoding native Go structures into generic map values.

[![GoDoc](https://godoc.org/github.com/things-go/mapstruct?status.svg)](https://godoc.org/github.com/things-go/mapstruct)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/things-go/mapstruct?tab=doc)
![Action Status](https://github.com/things-go/mapstruct/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/things-go/mapstruct)](https://goreportcard.com/report/github.com/things-go/mapstruct)
[![Licence](https://img.shields.io/github/license/things-go/mapstruct)](https://raw.githubusercontent.com/things-go/mapstruct/main/LICENSE)
[![Tag](https://img.shields.io/github/v/tag/things-go/mapstruct)](https://github.com/things-go/mapstruct/tags)

### Installation

Use go get.

```bash
    go get -u github.com/things-go/mapstruct
```

Then import the mapstruct package into your own code.

```go
    import "github.com/things-go/mapstruct"
```

### Usage && Example

#### API

```go
// Encode takes an input structure and uses reflection to translate it to
// the output map[string]interface{} with default tag "map"
func Encode(input interface{}) map[string]interface{}

// Encode takes an input structure and uses reflection to translate it to
// the output map[string]interface{} with the custom tag name
func EncodeWithTag(input interface{}, tagName string) map[string]interface{}
```

This function converts a struct to map. To use this function, import it first:

```go
import "github.com/things-go/mapstruct"
```

Here is an example:

```go
m := mapstruct.Encode(struct {
    Id      int64
    Name    string
    }{
        Id:        1001,
        Name:      "mapstruct",
    })
/*
    map[string]interface{}{
        "Id": 1001,  
        "Name": "mapstruct",
    }
*/
```

Encode will create and return a new map from the given struct. The keys of the map will be the
name of fields. The values will be the value of fields.

Only [public fields](https://golang.org/doc/effective_go.html#names) will be processed. So **fields
starting with lowercase will be ignored**.

#### Name Tags

```go
type AA struct {
    Id        int64    `map:"id"`
    Name      string   `map:"name"`
}
```
We can give the field a tag to specify another name to be used as the key.

#### Ignore Field

```go
type AA struct {
    Ignore string `map:"-"`
}
```
If we give the special tag "-" to a field, it will be ignored.

#### Omit Empty

```go
type AA struct {
    Desc        string    `map:"desc,omitempty"`
}
```
If tag option is "omitempty", this field will not appear in the map if the value is empty.
Empty values are 0, false, "", nil, empty array and empty map.

#### To String

```go
type AA struct {
    Id        int64    `map:"id,string"`
    Price     float32  `map:"price,string"`
}
```
If tag option is "string", this field will be converted to string type. Encode will put the
original value to the map if the conversion is failed.

## References

- [mapstructure](https://github.com/mitchellh/mapstructure)

## License

This project is under MIT License. See the [LICENSE](LICENSE) file for the full license text.