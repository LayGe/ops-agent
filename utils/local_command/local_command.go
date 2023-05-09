package local_command

import (
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"os"
	"os/exec"
	"ysxs_ops_agent/pkg/log"
)

func ExecCmdWithOutput(w http.ResponseWriter, flush http.Flusher, bashCmd string, arg ...string) error {
	cmd := exec.Command(bashCmd, arg...)
	cmd.Env = os.Environ()

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}

	defer func() {
		if err := cmd.Wait(); err != nil {
			log.MainLog.Errorf("command finished with error: %s", err.Error())
		}
	}()

	eg := new(errgroup.Group)

	eg.Go(func() error {
		_, err := io.Copy(&chunkWriter{w, flush}, stdout)
		return err
	})

	eg.Go(func() error {
		_, err = io.Copy(&chunkWriter{w, flush}, stderr)
		return err
	})

	if err = eg.Wait(); err != nil {
		return err
	}
	return nil
}

type chunkWriter struct {
	http.ResponseWriter
	http.Flusher
}

func (cw *chunkWriter) Write(p []byte) (int, error) {
	n, err := cw.ResponseWriter.Write(p)
	if err != nil {
		return n, err
	}
	cw.Flush()
	return n, nil
}
