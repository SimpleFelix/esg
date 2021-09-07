# ESG
ESG is short for Error Struct Generator

### Build
`go build -o esg`

### Usage
`./esg output_dir pkg_name error_code formatted_message [name_of_arguments..]`

name_of_argument must not be ErrorCode.
### Example
`./esg /src/myproj errors InvalidPhone "%v is not valid phone number." phone`

The generated file is at `./src/myproj/errors/InvalidPhone.go`

The source code is like below.
```go
// Package errors
// Generated by ESG. https://github.com/SimpleFelix/esg
package errors

import "fmt"

type InvalidPhone struct {
	phone interface{}
}

func (e InvalidPhone)Code() string {
	return "InvalidPhone"
}

func (e InvalidPhone)Error() string {
	return fmt.Sprintf("%v is not valid phone number.", e.phone)
}

func NewInvalidPhone(phone interface{}) InvalidPhone {
	return InvalidPhone{
		phone: phone,
	}
}
```