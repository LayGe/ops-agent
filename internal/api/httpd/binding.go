package httpd

import (
	"encoding/json"
	"net/http"
	"ysxs_ops_agent/internal/api/httpd/res"
	"ysxs_ops_agent/internal/api/httpd/validator"
)

func bindAndCheck(w http.ResponseWriter, r *http.Request, req interface{}) bool {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		res.ErrorResponse(w, r, err.Error())
		return true
	}

	if eStr := validator.DefaultProbeValidator.Check(req); eStr != "" {
		res.ErrorResponse(w, r, eStr)
		return true
	}
	return false
}
