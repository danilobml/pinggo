package errs

import "errors"

var ErrInvalidInputFile = errors.New("invalid input file type. Must be a .txt")

var ErrPingFailed = errors.New("ping to this URL failed")
