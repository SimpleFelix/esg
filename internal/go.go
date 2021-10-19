package internal

import (
	"fmt"
	"os"
	"time"
)

func GenerateGoCode(args []string) (dir, file, source string) {
	numberOfArgs := len(args)
	if numberOfArgs < 4 {
		showHelpAndExit()
	}

	statusCode := "500"
	if args[0] == "-sc" {
		if numberOfArgs < 6 {
			showHelpAndExit()
		}
		statusCode = args[1]
		args = args[2:]
	}

	dir = args[0]
	pkg := args[1]
	errCode := args[2]
	msg := args[3]
	var formatArgs []string
	haveArgs := numberOfArgs > 4
	if haveArgs {
		formatArgs = args[4:]
	}

	// write struct
	lastIdx := len(formatArgs) - 1
	source = fmt.Sprintf(`// Generated by ESG at %s. github.com/simplefelix/esg

package %s

import "fmt"

type %s struct {
	_extra_ interface{}
`, time.Now().Format("2006-01-02 15:04:05"),
		pkg, errCode)
	if haveArgs {
		for _, arg := range formatArgs {
			source += fmt.Sprintf("	%s interface{}", arg)
			source += "\n"
		}
	}
	source += "}\n"

	// write Code() function
	source += fmt.Sprintf(`
// ErrorCode change it as you prefer.
func (e *%s) ErrorCode() interface{} {
	return "%s"
}

// StatusCode refers to http response status code.
// Developer may want to set response status code based on error.
// For example, if the error is caused by bad request, then change the return value to 400.
// Ignore this function if no need for your project.
func (e *%s) StatusCode() int {
	return %s
}

// Extra returns _extra_ which can be set by user. Usage of _extra_ is determined by user.
func (e *%s) Extra() interface{} {
	return e._extra_
}

// SetExtra sets _extra_ with a value by user. Usage of _extra_ is determined by user.
func (e *%s) SetExtra(extra interface{}) {
	e._extra_ = extra
}

`, errCode, errCode, errCode, statusCode, errCode, errCode)

	// write Error() function
	source += `// Error implementation to error interface.`
	source += fmt.Sprintf("\nfunc (e *%s) Error() string {\n	return fmt.Sprintf(`%s`", errCode, msg)
	if haveArgs {
		for _, arg := range formatArgs {
			source += fmt.Sprintf(", e.%s", arg)
		}
	}
	source += ")\n}\n"

	// write New() function
	source += fmt.Sprintf(`
// Err%s is convenient constructor.
func Err%s(`, errCode, errCode)
	if haveArgs {
		for idx, arg := range formatArgs {
			source += fmt.Sprintf("%s", arg)
			if idx != lastIdx {
				source += ", "
			} else {
				source += " interface{}"
			}
		}
	}
	source += fmt.Sprintf(`) *%s {
	return &%s{
`, errCode, errCode)
	if haveArgs {
		for _, arg := range formatArgs {
			source += fmt.Sprintf("		%s: %s,\n", arg, arg)
		}
	}
	source += "	}\n}\n"

	return dir, errCode + ".go", source
}

func showHelpAndExit() {
	fmt.Println(`
Usage: esg go [-sc statu_code] output_dir pkg_name error_code formatted_message [name_of_arguments..]
Example: esg go . errors InvalidPhone "%s is not valid phone number." Phone
`)
	os.Exit(0)
}
