package rect

import (
	"audit-rectification/internal/rectserver/global"
	"audit-rectification/internal/rectserver/model"
	"audit-rectification/internal/rectserver/repo"
	"audit-rectification/internal/rectserver/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"regexp"
)

type ServiceRect struct {
	rep repo.Repository
}

func NewServiceRect(rs repo.Repository) service.Service {
	return &ServiceRect{rs}
}

func (rs *ServiceRect) Fetching(c *gin.Context, str string) (model.Response, error) {
	var qv = reflect.ValueOf(rs)
	for name, pattern := range global.Patterns {
		if pattern.MatchString(str) {
			ret := qv.MethodByName(name).Call(
				[]reflect.Value{reflect.ValueOf(c), reflect.ValueOf(pattern), reflect.ValueOf(str)})
			return ret[0].Interface().(model.Response), nil
		}
	}
	return model.Response{}, fmt.Errorf("no match found")
}

// P1 札幌营业部未按时完成的整改报告有哪些？
func (rs *ServiceRect) P1(ctx *gin.Context, reg *regexp.Regexp, str string) (model.Response, error) {
	// TODO: 查询nebula数据库
	fmt.Println(ctx.Get("id"))
	fmt.Println(reg.FindAllStringSubmatch(str, -1))
	return rs.rep.Fetch(ctx, str)
}