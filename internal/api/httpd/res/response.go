package res

import (
	"bytes"
	"encoding/json"
	"net/http"
	"ysxs_ops_agent/internal/dto"
	"ysxs_ops_agent/pkg/log"
)

type H map[string]any

func JSONResponse(w http.ResponseWriter, r *http.Request, result interface{}, codes ...int) {

	var code = http.StatusOK
	if codes != nil && len(codes) > 0 {
		code = codes[0]
	}
	var data = &dto.Response{
		Code:    code,
		Data:    result,
		Success: true,
	}

	body, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.MainLog.Errorf("JSON marshal failed, err: [%s]", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	w.Write(prettyJSON(body))
}

func ErrorResponse(w http.ResponseWriter, r *http.Request, error interface{}, codes ...int) {
	var code = http.StatusOK
	if codes != nil && len(codes) > 0 {
		code = codes[0]
	}
	data := &dto.Response{
		Error:   error,
		Code:    code,
		Success: false,
	}

	body, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.MainLog.Errorf("JSON marshal failed, err:[%s]", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(prettyJSON(body))
}

func prettyJSON(b []byte) []byte {
	var out bytes.Buffer
	json.Indent(&out, b, "", "  ")
	return out.Bytes()
}
