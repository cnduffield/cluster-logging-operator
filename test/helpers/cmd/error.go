package cmd

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

const stderrLimit = 1024

// stderrBuffer is a bounded buffer for capturing (some of) stderr output
type stderrBuffer struct{ b bytes.Buffer }

func (b *stderrBuffer) Write(data []byte) (int, error) {
	max := stderrLimit - b.b.Len()
	if max < 0 {
		max = 0
	}
	if len(data) > max {
		data = data[:max]
	}
	return b.b.Write(data)
}

// closeOnErr if err != nil close c, return err.
func closeOnErr(c io.Closer, err error) error {
	if err != nil {
		if err2 := c.Close(); err2 != nil {
			err = fmt.Errorf("%v: %w", err, err2)
		}
	}
	return err
}

// closeErr returns an error including the stderr text if there is any.
func closeErr(err error, stderr stderrBuffer) error {
	errtxt := strings.TrimSpace(stderr.b.String())
	if err != nil && errtxt != "" {
		return fmt.Errorf("%w: %s", err, errtxt)
	}
	return err
}
