package profilePrimitive

import "task-tracker/infrastructure/errors"

var (
	ErrFirstNameIsRequired = errors.NewError("PRP-001", "Имя должно быть заполнено")
	ErrLastNameIsRequired  = errors.NewError("PRP-002", "Фамилия должна быть заполнена")
)
