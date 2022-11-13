package golang_testing

import "errors"

var (
	ErrNotFound = errors.New("not found")
)

type Hashtable interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
}
