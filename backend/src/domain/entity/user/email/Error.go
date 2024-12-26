package emailPrimitive

import "task-tracker/infrastructure/errors"

var (
	ErrEmailIsEmpty   = errors.NewError("EMP-001", "Email пустой")
	ErrEmailIsInvalid = errors.NewError("EMP-002", "Некорректный Email")
)