package userstory

import (
	"fmt"
	"io"
)

type printRead struct {
	r io.Reader
}

func (s *printRead) Read(p []byte) (n int, err error) {
	n, err = s.r.Read(p)
	if err == nil {
		fmt.Print(string(p))
	}
	return
}
