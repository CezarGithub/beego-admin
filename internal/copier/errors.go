package copier

import "errors"

var (
	ErrInvalidCopyDestination        = errors.New("{{copier.invalid_destination}}")
	ErrInvalidCopyFrom               = errors.New("{{copier.invalid_source}}")
	ErrMapKeyNotMatch                = errors.New("{{copier.key_not_match}}")
	ErrNotSupported                  = errors.New("{{copier.not_supported}}")
	ErrFieldNameTagStartNotUpperCase = errors.New("{{copier.tag_name_upper}}")
)
