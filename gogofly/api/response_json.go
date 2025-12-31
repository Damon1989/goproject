package api

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type ResponseJson struct {
	Status int    `json:"-"`
	Code   int    `json:"code,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   any    `json:"data,omitempty"`
}

func (r ResponseJson) IsEmpty() bool {
	return reflect.DeepEqual(r, ResponseJson{})
}

func HttpResponse(ctx *gin.Context, status int, resp ResponseJson) {
	if resp.IsEmpty() {
		ctx.AbortWithStatus(status)
		return
	}
	ctx.AbortWithStatusJSON(status, resp)
}

func buildStatus(resp ResponseJson, nDefault int) int {
	if 0 == resp.Status {
		return nDefault
	}
	return resp.Status
}

func OK(ctx *gin.Context, json ResponseJson) {
	HttpResponse(ctx, buildStatus(json, http.StatusOK), json)
}
func Fail(ctx *gin.Context, json ResponseJson) {
	HttpResponse(ctx, buildStatus(json, http.StatusBadRequest), json)
}

func ServerFail(ctx *gin.Context, json ResponseJson) {
	HttpResponse(ctx, buildStatus(json, http.StatusInternalServerError), json)
}
