package tar

import "errors"

var (
	ErrWriteTooLong = errors.New("archive/tar:write data too long")
)
