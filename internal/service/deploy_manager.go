package service

import (
	"context"
	"errors"
	"net/http"
	"ysxs_ops_agent/internal/schema"
	"ysxs_ops_agent/utils/local_command"
)

type DMService struct{}

func NewDMService() *DMService {
	return &DMService{}
}

func (dms *DMService) DeployNormalWithStream(_ context.Context, req *schema.DeployReq, w http.ResponseWriter) error {
	flusher, ok := w.(http.Flusher)
	if !ok {
		return errors.New("streaming not supported")
	}
	if err := local_command.ExecCmdWithOutput(
		w, flusher, "/bin/bash", "-c", req.Script); err != nil {
		return err
	}
	return nil
}
