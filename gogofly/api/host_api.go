package api

import (
	"github.com/damon/gogofly/service"
	"github.com/damon/gogofly/service/dto"
	"github.com/gin-gonic/gin"
)

type HostApi struct {
	BaseApi
	Service *service.HostService
}

func NewHostApi() HostApi {
	return HostApi{
		Service: service.NewHostService(),
	}
}

func (m HostApi) Shutdown(c *gin.Context) {
	var iShutdownHostDTO dto.ShutdownHostDTO
	if m.BuildRequest(BuildRequestOptions{
		Ctx: c,
		DTO: &iShutdownHostDTO,
	}).GetError() != nil {
		return
	}
	err := m.Service.Shutdown(iShutdownHostDTO)
	if err != nil {
		m.Fail(ResponseJson{
			Code: 10001,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(ResponseJson{
		Msg: "success",
	})
}
