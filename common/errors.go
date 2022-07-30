package common

type MallError struct {
	Message string
}

func (m MallError) Error() string {
	return m.Message
}

func Fail(message string) error {
	return &MallError{Message: message}
}
