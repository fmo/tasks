package project_errors

import "errors"

var ErrPriorityRange = errors.New("priority has to be between 1-5")
