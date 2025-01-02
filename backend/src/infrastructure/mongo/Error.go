package mongoRepo

import "task-tracker/infrastructure/errors"

var (
	ErrMongoDbIsRequired = errors.NewError("SYS", "MongoDB is required")
	ErrLoggerIsRequired  = errors.NewError("SYS", "Logger is required")
)