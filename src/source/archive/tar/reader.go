package tar

import "errors"

var (
	ErrHeader = errors.New("archive/tar: invalid tar header")
)
