package closer

import (
	"io"
	"sync"
)

// Closer interface.
type Closer = io.Closer

// CloseFunc interface.
type CloseFunc func() error

var (
	defaultCloser = NewLifoCloser()
	once          sync.Once
)

// Add appends closers to default closers list.
func Add(closers ...Closer) {
	defaultCloser.Add(closers...)
}

// Fn return Closer from close func.
func Fn(closeFunc CloseFunc) Closer {
	return &fnCloser{closer: closeFunc}
}

// Close calls close func from default closer.
func Close() (err error) {
	once.Do(func() {
		err = defaultCloser.Close()
	})

	return
}

type fnCloser struct {
	closer CloseFunc
}

// Close calls closer.
func (c *fnCloser) Close() error {
	return c.closer()
}
