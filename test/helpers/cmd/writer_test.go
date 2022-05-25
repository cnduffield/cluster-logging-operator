package cmd

import (
	"context"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/openshift/cluster-logging-operator/test/matchers"
)

var _ = Describe("Writer", func() {
	It("writes to command", func() {
		tmp, err := os.CreateTemp("", "test")
		ExpectOK(err)
		defer func() { _ = os.Remove(tmp.Name()) }()
		cmd := exec.Command("tee", "-a", tmp.Name())
		w, err := NewWriter(cmd)
		ExpectOK(err)
		defer w.Close()
		hello := "hello world\n"
		goodbye := "goodbye cruel world\n"
		_, err = w.Write([]byte(hello))
		ExpectOK(err)
		_, err = w.Write([]byte(goodbye))
		ExpectOK(err)
		Eventually(func() string {
			got, err := ioutil.ReadFile(tmp.Name())
			ExpectOK(err)
			return string(got)
		}).Should(Equal(hello + goodbye))
	})

	It("times out", func() {
		// Set up a cat command to write to a pipe that is never read - this will block the writer.
		cmd := exec.Command("cat")
		_, pw, err := os.Pipe()
		ExpectOK(err)
		cmd.Stdout = pw
		ExpectOK(err)
		w, err := NewWriter(cmd)
		ExpectOK(err)
		defer w.Close()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second/100)
		defer cancel()
		// Keep writing so we fill process buffers and eventually block and time out.
		for err == nil {
			_, err = w.WriteContext(ctx, []byte("testing"))
		}
		Expect(err.Error()).To(Equal("context deadline exceeded: signal: killed"))
	})
})
