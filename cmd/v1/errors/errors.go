package errors

import (
	"errors"
)

var (
	NothingTodo       = errors.New("nothing todo")
	InvalidKey        = errors.New("invalid key")
	NoModelCanExecute = errors.New("no model executed, make sure to run cmd using flag --models")
	NoModelFound      = errors.New("no model found")
	InvalidModelFlag  = errors.New("invalid model flag, example: Company ~model~")
	NoEndModelFound   = errors.New("invalid model flag, no end model found")
)
