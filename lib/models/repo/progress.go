package repo

import (
	"strings"
)

type cloneProgress struct {
	channel chan interface{}
}

func (p cloneProgress) Write(s []byte) (int, error) {
	if strings.HasPrefix(string(s), "Total") {
		go func() {
			p.channel <- nil
		}()
	}
	return len(s), nil
}