package schema

type DeployReq struct {
	Script string `json:"script" validate:"required"`
}
