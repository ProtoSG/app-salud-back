package utils

import "fmt"

type ErrInvalidCredentials struct {
	msg string
}

func NewErrInvalidCredentials(msg string) *ErrInvalidCredentials {
	return &ErrInvalidCredentials{msg}
}

func (this *ErrInvalidCredentials) Error() string {
	return fmt.Sprintf("%s", this.msg)
}
