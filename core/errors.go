package core

type BadGuessError struct{}

func NewBadGuessError() *BadGuessError {
	return &BadGuessError{}
}

func (badGuessError *BadGuessError) Error() string {
	return "bad guess"
}
