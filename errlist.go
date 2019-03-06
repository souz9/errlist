package errlist

import (
	"strings"
)

// Append appends a error(s) to the list of errors.
// If the appending error(s) is nil then it will be ignored.
func Append(errs []error, err ...error) []error {
	for _, err := range err {
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

// Error returns one new error built from the list of errors.
// If nil or an empty list is given then nil (no error) will be returned.
func Error(errs []error) error {
	if len(errs) > 0 {
		return list(errs)
	}
	return nil
}

type list []error

func (l list) Error() string {
	line := &strings.Builder{}
	if len(l) > 0 {
		line.WriteString(l[0].Error())
	}
	for i := 1; i < len(l); i++ {
		line.WriteString("; ")
		line.WriteString(l[i].Error())
	}
	return line.String()
}
