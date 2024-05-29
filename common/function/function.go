package function

import (
	"context"
	"framework-gin/common"
	"framework-gin/pojo/request"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/qiafan666/gotato/commons"
	"github.com/qiafan666/gotato/commons/log"
	"github.com/qiafan666/gotato/commons/utils"
	"reflect"
)

// BindAndValid binds and validates data
func BindAndValid(entity interface{}, ctx *gin.Context) (commons.ResponseCode, error) {

	//set base request parameter
	object := reflect.ValueOf(entity)

	baseRequest := ctx.Keys[(common.BaseRequest)].(request.BaseRequest)
	elem := object.Elem()
	base := elem.FieldByName("BaseRequest")
	if base.Kind() != reflect.Invalid {
		base.Set(reflect.ValueOf(baseRequest))
	}

	baseTokenRequest, _ := ctx.Keys[(common.BaseTokenRequest)].(request.BaseTokenRequest)
	baseToken := elem.FieldByName("BaseTokenRequest")
	if baseToken.Kind() != reflect.Invalid {
		baseToken.Set(reflect.ValueOf(baseTokenRequest))
	}

	err := ctx.MustBindWith(entity, binding.JSON)
	if err != nil {
		log.Slog.ErrorF(ctx.Value("ctx").(context.Context), "BindAndValid error: %v", err)
		return commons.ParameterError, err
	}

	if err = utils.Validate(entity); err != nil {
		log.Slog.ErrorF(ctx.Value("ctx").(context.Context), "Validate error: %v", err)
		return commons.ValidateError, err
	}

	return commons.OK, nil
}

func GetTraceId(ctx *gin.Context) string {
	if traceId, ok := ctx.Value("trace_id").(string); ok {
		return traceId
	} else {
		return ""
	}
}
func GetCtx(ctx *gin.Context) context.Context {
	if ctx, ok := ctx.Value("ctx").(context.Context); ok {
		return ctx
	} else {
		return context.Background()
	}
}
