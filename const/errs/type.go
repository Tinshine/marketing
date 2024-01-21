package errs

import "fmt"

type Error struct {
	code int
	msg  string
}

func (e *Error) Error() string {
	return fmt.Sprintf("code = %d, msg = %s", e.code, e.msg)
}

func newErr(code int, msg string) *Error {
	return &Error{code: code, msg: msg}
}

func (e *Error) GetCode() int   { return e.code }
func (e *Error) GetMsg() string { return e.msg }
