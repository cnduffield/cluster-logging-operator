package cmd

import (
	"context"
	"io"
	"os/exec"

	"github.com/ViaQ/logerr/log"
	"github.com/openshift/cluster-logging-operator/test/runtime"

	corev1 "k8s.io/api/core/v1"
)

// Writer writes to a running exec.Cmd.
type Writer struct {
	cmd    *exec.Cmd
	w      io.WriteCloser
	stderr stderrBuffer
}

// NewWriter starts an exec.Cmd and returns a CmdWriter for its stdin.
func NewWriter(cmd *exec.Cmd) (*Writer, error) {
	log.V(3).Info("Creating cmd.Writer", "cmd", cmd.String())
	p, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	w := &Writer{cmd: cmd, w: p}
	if cmd.Stderr == nil {
		// Capture stderr because exec.Start() doesn't fill in exec.ExitError.Stderr.
		cmd.Stderr = &w.stderr
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return w, nil
}

// Write implements io.Writer, no timeout.
func (w *Writer) Write(b []byte) (int, error) {
	n, err := w.w.Write(b)
	return n, closeOnErr(w, err)
}

/// FIXME HERE

// WriteContext returns on success or when the context is cancelled or times out.
func (w *Writer) WriteContext(ctx context.Context, data []byte) (int, error) {
	type result struct {
		n   int
		err error
	}
	done := make(chan result)

	// Read in a goroutine, abandon if context ends.
	go func() { n, err := w.w.Write(data); done <- result{n, err} }()
	select {
	case r := <-done:
		return r.n, closeOnErr(w, r.err)
	case <-ctx.Done():
		return 0, closeOnErr(w, ctx.Err())
	}
}

// Close kills the underlying process, if running.
// This closes the stdout pipe, so Read or ReadLine will return an error.
// Returns the error returned by the sub-process.
func (w *Writer) Close() error {
	w.w.Close()              // Try to close politely
	_ = w.cmd.Process.Kill() // Kill to make sure.
	return closeErr(w.cmd.Wait(), w.stderr)
}

// AppendWriter returns a CmdWriter that appends to file at path on pod.
func AppendWriter(pod *corev1.Pod, path string) (*Writer, error) {
	return NewWriter(runtime.Exec(pod, "tee", "-a", path))
}

// AppendWriterForContainer returns a CmdWriter that tails file at path on container in pod.
func AppendWriterForContainer(pod *corev1.Pod, containerName, path string) (*Writer, error) {
	return NewWriter(runtime.ExecContainer(pod, containerName, "tee", "-a", path))
}
