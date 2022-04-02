package errors

import "errors"

var (
	NoRecordFound = errors.New("repository: record not found")
)
